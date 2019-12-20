package outscale

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/outscale/osc-go/oapi"
	"github.com/terraform-providers/terraform-provider-outscale/utils"
)

func resourceOutscaleOAPINatService() *schema.Resource {
	return &schema.Resource{
		Create: resourceOAPINatServiceCreate,
		Read:   resourceOAPINatServiceRead,
		Delete: resourceOAPINatServiceDelete,
		Update: resourceOutscaleOAPINatServiceUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: getOAPINatServiceSchema(),
	}
}

func resourceOAPINatServiceCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OAPI

	// Create the NAT Service
	createOpts := &oapi.CreateNatServiceRequest{
		PublicIpId: d.Get("public_ip_id").(string),
		SubnetId:   d.Get("subnet_id").(string),
	}

	var resp *oapi.POST_CreateNatServiceResponses

	err := resource.Retry(5*time.Minute, func() *resource.RetryError {
		var err error
		resp, err = conn.POST_CreateNatService(*createOpts)

		if err != nil {
			if strings.Contains(err.Error(), "RequestLimitExceeded:") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})

	var errString string

	if err != nil || resp.OK == nil {
		if err != nil {
			errString = err.Error()
		} else if resp.Code401 != nil {
			errString = fmt.Sprintf("ErrorCode: 401, %s", utils.ToJSONString(resp.Code401))
		} else if resp.Code400 != nil {
			errString = fmt.Sprintf("ErrorCode: 400, %s", utils.ToJSONString(resp.Code400))
		} else if resp.Code500 != nil {
			errString = fmt.Sprintf("ErrorCode: 500, %s", utils.ToJSONString(resp.Code500))
		}

		return fmt.Errorf("[DEBUG] Error creating NAT Service (%s)", errString)
	}

	response := resp.OK

	// Get the ID and store it
	ng := response.NatService
	d.SetId(ng.NatServiceId)
	log.Printf("\n\n[INFO] NAT Service ID: %s", d.Id())

	// Wait for the NAT Service to become available
	log.Printf("\n\n[DEBUG] Waiting for NAT Service (%s) to become available", d.Id())
	stateConf := &resource.StateChangeConf{
		Pending: []string{"pending"},
		Target:  []string{"available"},
		Refresh: NGOAPIStateRefreshFunc(conn, d.Id()),
		Timeout: 10 * time.Minute,
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error waiting for NAT Service (%s) to become available: %s", d.Id(), err)
	}
	//SetTags
	if tags, ok := d.GetOk("tags"); ok {
		err := assignOapiTags(tags.([]interface{}), response.NatService.NatServiceId, conn)
		if err != nil {
			return err
		}
	}

	d.Set("request_id", resp.OK.ResponseContext.RequestId)
	// Update our attributes and return
	return resourceOAPINatServiceRead(d, meta)
}

func resourceOAPINatServiceRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OAPI

	// Refresh the NAT Service state
	ngRaw, state, err := NGOAPIStateRefreshFunc(conn, d.Id())()
	if err != nil {
		return err
	}

	status := map[string]bool{
		"deleted":  true,
		"deleting": true,
		"failed":   true,
	}

	if _, ok := status[strings.ToLower(state)]; ngRaw == nil || ok {
		fmt.Printf("\n\n[INFO] Removing %s from Terraform state as it is not found or in the deleted state.", d.Id())
		d.SetId("")
		return nil
	}

	// Set NAT Service attributes
	ng := ngRaw.(oapi.NatService)

	if ng.NatServiceId != "" {
		d.Set("nat_service_id", ng.NatServiceId)
	}
	if ng.State != "" {
		d.Set("state", ng.State)
	}
	if ng.SubnetId != "" {
		d.Set("subnet_id", ng.SubnetId)
	}
	if ng.NetId != "" {
		d.Set("net_id", ng.NetId)
	}

	if ng.PublicIps != nil {
		addresses := make([]map[string]interface{}, len(ng.PublicIps))

		for k, v := range ng.PublicIps {
			address := make(map[string]interface{})
			if v.PublicIpId != "" {
				address["public_ip_id"] = v.PublicIpId
			}
			if v.PublicIp != "" {
				address["public_ip"] = v.PublicIp
			}
			addresses[k] = address
		}
		if err := d.Set("public_ips", addresses); err != nil {
			return err
		}
	}
	if err := d.Set("tags", tagsOAPIToMap(ng.Tags)); err != nil {
		fmt.Printf("[WARN] ERROR TAGS PROBLEME (%s)", err)
	}

	return nil
}

func resourceOutscaleOAPINatServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OAPI

	d.Partial(true)

	if err := setOAPITags(conn, d); err != nil {
		return err
	}

	d.SetPartial("tags")

	d.Partial(false)
	return resourceOAPINatServiceRead(d, meta)
}

func resourceOAPINatServiceDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OAPI

	log.Printf("[INFO] Deleting NAT Service: %s\n", d.Id())
	req := &oapi.DeleteNatServiceRequest{
		NatServiceId: d.Id(),
	}

	var resp *oapi.POST_DeleteNatServiceResponses

	err := resource.Retry(5*time.Minute, func() *resource.RetryError {
		var err error
		resp, err = conn.POST_DeleteNatService(*req)
		if err != nil {
			if strings.Contains(err.Error(), "RequestLimitExceeded:") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})

	if err != nil || resp.OK == nil {
		var errString string

		if err != nil {
			if strings.Contains(err.Error(), "NatGatewayNotFound:") {
				return nil
			}
			errString = err.Error()
		} else if resp.Code401 != nil {
			errString = fmt.Sprintf("ErrorCode: 401, %s", utils.ToJSONString(resp.Code401))
		} else if resp.Code400 != nil {
			errString = fmt.Sprintf("ErrorCode: 400, %s", utils.ToJSONString(resp.Code400))
		} else if resp.Code500 != nil {
			errString = fmt.Sprintf("ErrorCode: 500, %s", utils.ToJSONString(resp.Code500))
		}
		return fmt.Errorf("error deleting Nat Service (%s)", errString)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting"},
		Target:     []string{"deleted"},
		Refresh:    NGOAPIStateRefreshFunc(conn, d.Id()),
		Timeout:    30 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	_, stateErr := stateConf.WaitForState()
	if stateErr != nil {
		return fmt.Errorf("Error waiting for NAT Service (%s) to delete: %s", d.Id(), stateErr)
	}

	return nil
}

// NGOAPIStateRefreshFunc returns a resource.StateRefreshFunc that is used to watch
// a NAT Service.
func NGOAPIStateRefreshFunc(conn *oapi.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		opts := &oapi.ReadNatServicesRequest{
			Filters: oapi.FiltersNatService{NatServiceIds: []string{id}},
		}

		var resp *oapi.POST_ReadNatServicesResponses
		var err error

		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			resp, err = conn.POST_ReadNatServices(*opts)
			return resource.NonRetryableError(err)
		})

		if err != nil || resp.OK == nil {
			var errString string
			if err != nil {

				if strings.Contains(fmt.Sprint(err), "NatGatewayNotFound") {
					return nil, "", nil
				}
				errString = err.Error()
			} else if resp.Code401 != nil {
				errString = fmt.Sprintf("ErrorCode: 401, %s", utils.ToJSONString(resp.Code401))
			} else if resp.Code400 != nil {
				errString = fmt.Sprintf("ErrorCode: 400, %s", utils.ToJSONString(resp.Code400))
			} else if resp.Code500 != nil {
				errString = fmt.Sprintf("ErrorCode: 500, %s", utils.ToJSONString(resp.Code500))
			}
			return nil, "", fmt.Errorf("[DEBUG] Error reading Subnet (%s)", errString)
		}

		response := resp.OK

		if response == nil || len(response.NatServices) == 0 {
			return oapi.NatService{NatServiceId: id, State: "deleted"}, "deleted", nil
		}

		ng := response.NatServices[0]
		return ng, ng.State, nil
	}
}

func getOAPINatServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Arguments
		"public_ip_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"token": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
			Computed: true,
		},
		"request_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subnet_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		// Attributes
		"public_ips": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"public_ip_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"public_ip": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
		"nat_service_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"state": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"net_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"tags": tagsListOAPISchema(),
	}
}
