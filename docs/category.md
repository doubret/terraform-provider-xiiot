# terraform-provider-xiiot

## Category reference

Terraform resource name : `xiiot_category`

## Example

```
resource "xiiot_category" "category" {
  name                          = "Category"
  purpose                       = "Purpose"
  values                        = [ "Value1", "Value2" ]
}
```

## Arguments

* __id__ (_Computed_)
  * Type: string
  * Id of the resource, this is the id for both terraform and Xi IoT
* __name__ (_Required_)
  * Type: string
  * Name of the category
* __purpose__ (_Required_)
  * Type: string
  * Purpose of the category
* __values__ (_Required_)
  * Type: set of strings
  * Values supported by the categoy
