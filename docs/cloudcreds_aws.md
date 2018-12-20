# terraform-provider-xiiot

## Cloudcreds AWS reference

Terraform resource name : `xiiot_cloudcredsaws`

## Example

```
resource "xiiot_cloudcredsaws" "cloudcredsaws" {
  name                          = "name"
  description                   = "description"
  access_key                    = "access_key"
  secret_key                    = "secret_key"
}
```

## Arguments

* __id__ (_Computed_)
  * Type: string
  * Id of the resource, this is the id for both terraform and Xi IoT.
* __name__ (_Required_)
  * Type: string
  * Cloudcreds name.
* __description__ (_Optional_)
  * Type: string
  * Cloudcreds description.
* __access_key__ (_Required_)
  * Type: string
  * AWS credentials access key.
* __secret_key__ (_Required_)
  * Type: string
  * AWS credentials secret key.
