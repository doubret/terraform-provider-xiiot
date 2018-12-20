# terraform-provider-xiiot

## Edge reference

Terraform resource name : `xiiot_edge`

## Example

```
resource "xiiot_edge" "edge" {
  name                          = "name"
  description                   = "description"
  gateway                       = "1.2.3.4"
  ip_address                    = "5.6.7.8"
  serial_number                 = "serial_number"
  subnet                        = "255.255.255.0"

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
* __gateway__ (_Required_)
  * Type: string
  * Gateway ip for the edge.
* __ip_address__ (_Required_)
  * Type: string
  * Ip Address for the edge.
* __serial_number__ (_Required_)
  * Type: string
  * Serial number for the edge.
* __subnet__ (_Required_)
  * Type: string
  * Subnet mask for the edge.
* __label__ (_Optional_)
  * Type: set of [catgory_info](category_info.md)
  * CategoryInfo list.
