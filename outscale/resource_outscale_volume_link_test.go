package outscale

import (
	"fmt"
	oscgo "github.com/marinsalinas/osc-sdk-go"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccOutscaleOAPIVolumeAttachment_basic(t *testing.T) {
	omi := getOMIByRegion("eu-west-2", "centos").OMI
	region := os.Getenv("OUTSCALE_REGION")

	//var i oscgo.Vm
	//var v oscgo.Volume

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			skipIfNoOAPI(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOAPIVolumeAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOAPIVolumeAttachmentConfig(omi, "c4.large", region),
				Check:  resource.ComposeTestCheckFunc( /*
					resource.TestCheckResourceAttr(
						"outscale_volumes_link.ebs_att", "device_name", "/dev/sdh"),
					testAccCheckOSCAPIVMExists(
						"outscale_vm.web", &i),
					testAccCheckOAPIVolumeAttachmentExists(
						"outscale_volumes_link.ebs_att", &i, &v),*/
				),
			},
		},
	})
}

func testAccCheckOAPIVolumeAttachmentDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		log.Printf("\n\n----- This is never called")
		if rs.Type != "outscale_volume_link" {
			continue
		}
	}
	return nil
}

func testAccCheckOAPIVolumeAttachmentExists(n string, i *oscgo.Vm, v *oscgo.Volume) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		for _, b := range i.GetBlockDeviceMappings() {
			if rs.Primary.Attributes["device_name"] == b.GetDeviceName() {
				if rs.Primary.Attributes["volume_id"] == b.Bsu.GetVolumeId() {
					// pass
					return nil
				}
			}
		}

		return fmt.Errorf("Error finding instance/volume")
	}
}

func testAccOAPIVolumeAttachmentConfig(omi, vmType, region string) string {
	return fmt.Sprintf(`
		resource "outscale_vm" "web" {
			image_id                 = "%s"
			vm_type                  = "%s"
			keypair_name             = "terraform-basic"
			security_group_ids       = ["sg-f4b1c2f8"]
			placement_subregion_name = "%[3]sb"
		}

		resource "outscale_volume" "volume" {
			subregion_name = "%[3]sb"
			volume_type    = "gp2"
			size           = 1
		}

		resource "outscale_volumes_link" "ebs_att" {
			device_name = "/dev/sdh"
			volume_id   = "${outscale_volume.volume.id}"
			vm_id       = "${outscale_vm.web.id}"
		}
	`, omi, vmType, region)
}
