# terraform-provider-xiiot

## Container registry reference

Terraform resource name : `xiiot_containerregistry`

## Example

```
resource "xiiot_containerregistry" "containerregistry" {
  name                          = "name"
  description                   = "description"
  type                          = "type"
  server                        = "server"
  user_name                     = "user_name"
  email                         = "email"
  pwd                           = "pwd"
  cloud_creds_id                = "${xiiot_cloudcredsgcp.cloudcredsgcp.id}"
}
```

## Arguments

* __id__ (_Computed_)
  * Type: string
  * Id of the resource, this is the id for both terraform and Xi IoT
* __name__ (_Required_)
  * Type: string
  * Container registry name
* __description__ (_Optional_)
  * Type: string
  * Container registry description
* __type__ (_Required_)
  * Type: string
  * Allowed value : AWS, GCP, Azure, ContainerRegistry
  * Container registry type
* __server__ (_Required_)
  * Type: string
  * Container registry server
* __user_name__ (_Optional_)
  * Type: string
  * Container registry user name (required if type == ContainerRegistry)
* __email__ (_Optional_)
  * Type: string
  * Container registry email (required if type == ContainerRegistry)
* __pwd__ (_Optional_)
  * Type: string
  * Container registry password (required if type == ContainerRegistry)
* __cloud_creds_id__ (_Optional_)
  * Type: string
  * The cloud creds to import container registry profile from (required if type == AWS || Type == GCP)