# terraform-provider-xiiot

## Edge reference

Terraform resource name : `xiiot_edge`

## Example

```
resource "xiiot_edge" "edge" {
  name                          = "name"
  description                   = "description"
  edge_devices                  = 1
  gateway                       = "1.2.3.4"
  ip_address                    = "5.6.7.8"
  serial_number                 = "serial_number"
  storage_capacity              = 128
  storage_usage                 = 64
  subnet                        = "255.255.255.0"
  connected                     = false

  label {
    id                          = "${xiiot_category.category.id}"
    value                       = "Value1"
  }
}
```

## Arguments

* __id__ (_Computed_)
  * Type: string
  * Id of the resource, this is the id for both terraform and Xi IoT.
* __name__ (_Required_)
  * Type: string
  * Edge name.
* __description__ (_Optional_)
  * Type: string
  * Edge description.

* __label__ (_Optional_)
  * Type: set of [catgory_info](category_info.md)
  * CategoryInfo list.
