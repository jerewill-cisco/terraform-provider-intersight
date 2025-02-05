### Resource Creation

```hcl
resource "intersight_software_ucsd_distributable" "software_ucsd_distributable1" {
  name      = "software_ucsd_distributable1"
  sha512sum = "<sha_512_checksum>"
  size      = 7471044747
  component_meta = [{
    additional_properties = ""
    class_id              = "firmware.ComponentMeta"
    component_type        = ""
    description           = ""
    disruption            = ""
    image_path            = ""
    object_type           = "firmware.ComponentMeta"
    component_label       = "BIOS"
    is_oob_supported      = false
    model                 = ""
    oob_manageability     = null
    packed_version        = ""
    redfish_url           = ""
    vendor                = ""
  }]
  model         = ""
  mdfid         = ""
  platform_type = ""

   release_date {
     object_type = "softwarerepository.Release"
     moid        = var.release
   }
   catalog {
     object_type = "softwarerepository.Catalog"
     moid        = var.catalog
   }
  vendor = "Cisco"
  distributable_metas = [{
    object_type           = "firmware.DistributableMeta"
    additional_properties = ""
    class_id              = "firmware.DistributableMeta"
    moid                  = ""
    selector              = ""
  }]
  supported_models = ["HyperFlex Data Platform"]
}

 variable "release" {
   type = string
   description = " value for release"
 }

 variable "catalog" {
   type =  string
   description = "value for catalog"
 }
```