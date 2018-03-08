package outscale

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-outscale/osc/fcu"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceOutscaleInboundRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceOutscaleInboundRuleCreate,
		Read:   resourceOutscaleInboundRuleRead,
		Delete: resourceOutscaleInboundRuleDelete,
		// Importer: &schema.ResourceImporter{
		// State: schema.ImportStatePassthrough,
		// State: resourceOutscaleInboundImportState,
		// },

		Schema: map[string]*schema.Schema{
			"cidr_ip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"from_port": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_security_group_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_security_group_owner_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"to_port": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"ip_permissions": getIpPermissionsSchema(),
		},
	}
}

func resourceOutscaleInboundRuleCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).FCU

	sg_id := d.Get("group_id").(string)

	awsMutexKV.Lock(sg_id)
	defer awsMutexKV.Unlock(sg_id)

	sg, err := findResourceSecurityGroup(conn, sg_id)
	if err != nil {
		return err
	}

	perm, err := expandIPPerm(d, sg)
	if err != nil {
		return err
	}

	if err := validateOutscaleSecurityGroupRule(d); err != nil {
		return err
	}

	ruleType := "ingress"
	isVPC := sg.VpcId != nil && *sg.VpcId != ""

	var autherr error
	fmt.Printf("[DEBUG] Authorizing security group %s %s rule: %#v",
		sg_id, "Ingress", perm)

	req := &fcu.AuthorizeSecurityGroupIngressInput{
		GroupId:       sg.GroupId,
		IpPermissions: []*fcu.IpPermission{perm},
	}

	resource.Retry(5*time.Minute, func() *resource.RetryError {
		_, autherr = conn.VM.AuthorizeSecurityGroupIngress(req)

		if err != nil {
			if strings.Contains(err.Error(), "RequestLimitExceeded") || strings.Contains(err.Error(), "DependencyViolation") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}

		return nil
	})

	if autherr != nil {
		if awsErr, ok := autherr.(awserr.Error); ok {
			if awsErr.Code() == "InvalidPermission.Duplicate" {
				return fmt.Errorf(`[WARN] A duplicate Security Group rule was found on (%s). This may be
a side effect of a now-fixed Terraform issue causing two security groups with
identical attributes but different source_security_group_ids to overwrite each
other in the state. See https://github.com/hashicorp/terraform/pull/2376 for more
information and instructions for recovery. Error message: %s`, sg_id, awsErr.Message())
			}
		}

		return fmt.Errorf(
			"Error authorizing security group rule type %s: %s",
			ruleType, autherr)
	}

	id := ipPermissionIDHash(sg_id, ruleType, perm)
	log.Printf("[DEBUG] Computed group rule ID %s", id)

	retErr := resource.Retry(5*time.Minute, func() *resource.RetryError {
		sg, err := findResourceSecurityGroup(conn, sg_id)

		if err != nil {
			log.Printf("[DEBUG] Error finding Security Group (%s) for Rule (%s): %s", sg_id, id, err)
			return resource.NonRetryableError(err)
		}

		var rules []*fcu.IpPermission
		rules = sg.IpPermissions

		rule := findRuleMatch(perm, rules, isVPC)

		if rule == nil {
			log.Printf("[DEBUG] Unable to find matching %s Security Group Rule (%s) for Group %s",
				ruleType, id, sg_id)
			return resource.RetryableError(fmt.Errorf("No match found"))
		}

		log.Printf("[DEBUG] Found rule for Security Group Rule (%s): %s", id, rule)
		return nil
	})

	if retErr != nil {
		return fmt.Errorf("Error finding matching %s Security Group Rule (%s) for Group %s",
			ruleType, id, sg_id)
	}

	d.SetId(id)
	return nil
}

func resourceOutscaleInboundRuleRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).FCU
	sg_id := d.Get("group_id").(string)
	sg, err := findResourceSecurityGroup(conn, sg_id)
	if _, notFound := err.(securityGroupNotFound); notFound {
		// The security group containing this rule no longer exists.
		d.SetId("")
		return nil
	}
	if err != nil {
		return fmt.Errorf("Error finding security group (%s) for rule (%s): %s", sg_id, d.Id(), err)
	}

	isVPC := sg.VpcId != nil && *sg.VpcId != ""

	var rule *fcu.IpPermission
	var rules []*fcu.IpPermission
	ruleType := "ingress"
	rules = sg.IpPermissions

	p, err := expandIPPerm(d, sg)
	if err != nil {
		return err
	}

	if len(rules) == 0 {
		log.Printf("[WARN] No %s rules were found for Security Group (%s) looking for Security Group Rule (%s)",
			ruleType, *sg.GroupName, d.Id())
		d.SetId("")
		return nil
	}

	rule = findRuleMatch(p, rules, isVPC)

	if rule == nil {
		log.Printf("[DEBUG] Unable to find matching %s Security Group Rule (%s) for Group %s",
			ruleType, d.Id(), sg_id)
		d.SetId("")
		return nil
	}

	if err := setFromIPPerm(d, sg, p); err != nil {
		return errwrap.Wrapf("Error setting IP Permission for Security Group Rule: {{err}}", err)
	}

	log.Printf("[DEBUG] Found rule for Security Group Rule (%s): %s", d.Id(), rule)

	return nil
}

func resourceOutscaleInboundRuleDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).FCU
	sg_id := d.Get("group_id").(string)

	awsMutexKV.Lock(sg_id)
	defer awsMutexKV.Unlock(sg_id)

	sg, err := findResourceSecurityGroup(conn, sg_id)
	if err != nil {
		return err
	}

	perm, err := expandIPPerm(d, sg)
	if err != nil {
		return err
	}
	fmt.Printf("\n\n[DEBUG] Revoking security group %#v %s rule: %#v",
		sg_id, "ingress", perm)
	req := &fcu.RevokeSecurityGroupIngressInput{
		GroupId:       sg.GroupId,
		IpPermissions: []*fcu.IpPermission{perm},
	}

	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		_, err = conn.VM.RevokeSecurityGroupIngress(req)

		if err != nil {
			if strings.Contains(err.Error(), "RequestLimitExceeded") {
				fmt.Printf("\n\n[INFO] Request limit exceeded")
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf(
			"Error revoking security group %s rules: %s",
			sg_id, err)
	}

	d.SetId("")

	return nil
}
