package outscale

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-outscale/osc/fcu"
)

func resourceOutscaleOAPIReservedVmsOfferPurchase() *schema.Resource {
	return &schema.Resource{
		Create: resourceOutscaleOAPIReservedVmsOfferPurchaseCreate,
		Read:   resourceOutscaleOAPIReservedVmsOfferPurchaseRead,
		Delete: resourceOutscaleOAPIReservedVmsOfferPurchaseDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vm_count": {
				Type:     schema.TypeInt,
				ForceNew: true,
				Required: true,
			},
			"reserved_vms_offer_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			// Attributes
			"reserved_vms_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reserved_instances_offerings_set": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sub_region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"currency_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"duration": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"fixed_price": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"vm_tenancy": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vm_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"marketplace": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"offering_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"product_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reserved_vms_offer_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"usage_price": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"pricing_details_set": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"recurring_charges": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"frequency": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceOutscaleOAPIReservedVmsOfferPurchaseCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).FCU

	i, ok1 := d.GetOk("vm_count")
	r, ok2 := d.GetOk("reserved_vms_offer_id")

	if ok1 && ok2 {
		return fmt.Errorf("vm_count and reserved_vms_offer_id are required")
	}

	req := &fcu.PurchaseReservedInstancesOfferingInput{
		InstanceCount:               aws.Int64(int64(i.(int))),
		ReservedInstancesOfferingId: aws.String(r.(string)),
	}

	var resp *fcu.PurchaseReservedInstancesOfferingOutput
	var err error
	err = resource.Retry(120*time.Second, func() *resource.RetryError {
		resp, err = conn.VM.PurchaseReservedInstancesOffering(req)

		if err != nil {
			if strings.Contains(err.Error(), "RequestLimitExceeded:") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return resource.RetryableError(err)
	})
	if err != nil {
		log.Printf("[DEBUG] Error creating lin (%s)", err)
		return err
	}

	if resp == nil {
		return fmt.Errorf("Cannot create the oAPI vpc, empty response")
	}

	d.SetId(*resp.ReservedInstancesId)
	d.Set("reserved_vms_id", *resp.ReservedInstancesId)

	return resourceOutscaleLinRead(d, meta)
}

func resourceOutscaleOAPIReservedVmsOfferPurchaseRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).FCU

	req := &fcu.DescribeReservedInstancesOfferingsInput{
		ReservedInstancesOfferingIds: []*string{aws.String(d.Id())},
	}

	var resp *fcu.DescribeReservedInstancesOfferingsOutput
	var err error
	err = resource.Retry(120*time.Second, func() *resource.RetryError {
		resp, err = conn.VM.DescribeReservedInstancesOfferings(req)

		if err != nil {
			if strings.Contains(err.Error(), "RequestLimitExceeded:") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return resource.RetryableError(err)
	})
	if err != nil {
		log.Printf("[DEBUG] Error reading lin (%s)", err)
	}

	if resp == nil {
		d.SetId("")
		return fmt.Errorf("outscale_reserved_vms_offer_purchase not found")
	}

	if len(resp.ReservedInstancesOfferings) == 0 {
		d.SetId("")
		return fmt.Errorf("outscale_reserved_vms_offer_purchase not found")
	}

	rs := make([]map[string]interface{}, len(resp.ReservedInstancesOfferings))

	for k, v := range resp.ReservedInstancesOfferings {
		r := make(map[string]interface{})

		r["sub_region"] = *v.AvailabilityZone
		r["currency_code"] = *v.CurrencyCode
		r["duration"] = *v.Duration
		r["fixed_price"] = *v.FixedPrice
		r["vm_tenancy"] = *v.InstanceTenancy
		r["vm_type"] = *v.InstanceType
		r["marketplace"] = *v.Marketplace
		r["offering_type"] = *v.OfferingType
		r["product_description"] = *v.ProductDescription
		r["reserved_vms_offer_id"] = *v.ReservedInstancesOfferingId
		r["usage_price"] = *v.UsagePrice
		var a []map[string]interface{}
		for _, j := range v.PricingDetails {
			a = append(a, map[string]interface{}{
				"count": *j.Count,
			})
		}
		r["pricing_details"] = a

		var b []map[string]interface{}
		for _, l := range v.RecurringCharges {
			b = append(b, map[string]interface{}{
				"frequency": *l.Frequency,
			})
		}
		r["recurring_charges"] = b

		rs[k] = r
	}

	d.Set("reserved_instances_offerings_set", rs)
	d.Set("request_id", resp.RequestId)

	return nil
}

func resourceOutscaleOAPIReservedVmsOfferPurchaseDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")

	return nil
}
