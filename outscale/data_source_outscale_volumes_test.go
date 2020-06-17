package outscale

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccOutscaleOAPIVolumesDataSource_multipleFilters(t *testing.T) {
	region := os.Getenv("OUTSCALE_REGION")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOutscaleOAPIVolumeDataSourceConfigWithMultipleFilters(region),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPIVolumeDataSourceID("data.outscale_volumes.ebs_volume"),
					resource.TestCheckResourceAttr("data.outscale_volumes.ebs_volume", "volumes.0.size", "1"),
					resource.TestCheckResourceAttr("data.outscale_volumes.ebs_volume", "volumes.0.volume_type", "gp2"),
				),
			},
		},
	})
}

func TestAccOutscaleOAPIVolumeDataSource_multipleVIdsFilters(t *testing.T) {
	region := os.Getenv("OUTSCALE_REGION")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOutscaleOAPIVolumesDataSourceConfigWithMultipleVolumeIDsFilter(region),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPIVolumeDataSourceID("data.outscale_volumes.outscale_volumes"),
					resource.TestCheckResourceAttr("data.outscale_volumes.outscale_volumes", "volumes.0.size", "40"),
				),
			},
		},
	})
}

func TestAccOutscaleOAPIVolumesDataSource_withVM(t *testing.T) {
	region := os.Getenv("OUTSCALE_REGION")
	omi := os.Getenv("OUTSCALE_IMAGEID")

	sgName := acctest.RandomWithPrefix("testacc-sg")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOutscaleOAPIVolumesDataSourceConfigWithVM(region, omi, sgName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPIVolumeDataSourceID("data.outscale_volumes.outscale_volumes"),
					// resource.TestCheckResourceAttr("data.outscale_volumes.outscale_volumes", "volumes.0.size", "1"),
					// resource.TestCheckResourceAttr("data.outscale_volumes.outscale_volumes", "volumes.0.volume_type", "gp2"),
				),
			},
		},
	})
}

func testAccCheckOutscaleOAPIVolumeDataSourceConfigWithMultipleFilters(region string) string {
	return fmt.Sprintf(`
		resource "outscale_volume" "external" {
			subregion_name = "%sa"
			volume_type    = "gp2"
			size           = 1

			tags {
				key   = "Name"
				value = "tf-acc-test-ebs-volume-test"
			}
		}

		data "outscale_volumes" "ebs_volume" {
			filter {
				name   = "volume_sizes"
				values = ["${outscale_volume.external.size}"]
			}

			filter {
				name   = "volume_types"
				values = ["${outscale_volume.external.volume_type}"]
			}
		}
	`, region)
}

func testAccCheckOutscaleOAPIVolumesDataSourceConfigWithMultipleVolumeIDsFilter(region string) string {
	return fmt.Sprintf(`
		resource "outscale_volume" "outscale_volume" {
			subregion_name = "%[1]sa"
			size           = 40
		}

		resource "outscale_volume" "outscale_volume2" {
			subregion_name = "%[1]sa"
			size           = 40
		}

		data "outscale_volumes" "outscale_volumes" {
			filter {
				name   = "volume_ids"
				values = ["${outscale_volume.outscale_volume.volume_id}", "${outscale_volume.outscale_volume2.volume_id}"]
			}
		}
	`, region)
}

func testAccCheckOutscaleOAPIVolumesDataSourceConfigWithVM(region, imageID, sgName string) string {
	return fmt.Sprintf(`
		resource "outscale_volume" "outscale_volume" {
			subregion_name = "%[1]sa"
			volume_type    = "gp2"
			size           = 25
			tags {
				key   = "Name"
				value = "volume-gp2-1"
			}
		}

		resource "outscale_volume" "outscale_volume2" {
			subregion_name = "%[1]sa"
			volume_type    = "standard"
			size           = 13
			tags {
				key   = "Name"
				value = "volume-standard-2"
			}
		}

		resource "outscale_volume" "outscale_volume3" {
			subregion_name = "%[1]sa"
			size           = 40
			iops           = 100
			volume_type    = "io1"
			tags {
				key   = "type"
				value = "io1"
			}
		}

		resource "outscale_net" "net" {
			ip_range = "10.0.0.0/16"

			tags {
				key = "Name"
				value = "testacc-security-group-rs"
			}
		}

		resource "outscale_security_group" "sg" {
			security_group_name = "%[3]s"
			description         = "Used in the terraform acceptance tests"

			tags {
				key   = "Name"
				value = "tf-acc-test"
			}

			net_id = "${outscale_net.net.id}"
		}

		resource "outscale_vm" "outscale_vm" {
			image_id           = "%[2]s"
			vm_type            = "t2.micro"
			keypair_name       = "terraform-basic"
			security_group_ids = ["${outscale_security_group.sg.id}"]
		}

		resource "outscale_volumes_link" "outscale_volumes_link" {
			device_name = "/dev/xvdc"
			volume_id   = "${outscale_volume.outscale_volume.id}"
			vm_id       = "${outscale_vm.outscale_vm.id}"
		}

		resource "outscale_volumes_link" "outscale_volumes_link_2" {
			device_name = "/dev/xvdd"
			volume_id   = "${outscale_volume.outscale_volume2.id}"
			vm_id       = "${outscale_vm.outscale_vm.id}"
		}

		resource "outscale_volumes_link" "outscale_volumes_link_3" {
			device_name = "/dev/xvde"
			volume_id   = "${outscale_volume.outscale_volume3.id}"
			vm_id       = "${outscale_vm.outscale_vm.id}"
		}

		data "outscale_volumes" "outscale_volumes" {
			filter {
				name   = "link_volume_vm_ids"
				values = ["${outscale_vm.outscale_vm.vm_id}"]
			}
		}
	`, region, imageID, sgName)
}
