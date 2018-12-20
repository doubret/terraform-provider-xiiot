# terraform-provider-xiiot

## Docker profile reference

Terraform resource name : `xiiot_dockerprofile`

## Example

```
resource "xiiot_dockerprofile" "dockerprofile" {
  name                          = "name"
  description                   = "description"
  type                          = "type"
  server                        = "server"
  user_name                     = "user_name"
  email                         = "email"
  pwd                           = "pwd"
  cloud_creds_id                = "${xiiot_cloudcredsgcp.cloudcredsgcp.id}"
  credentials                   = "credentials"
}
```

## Arguments

* __id__ (_Computed_)
  * Type: string
  * Id of the resource, this is the id for both terraform and Xi IoT.
* __name__ (_Required_)
  * Type: string
  * Docker profile name.
* __description__ (_Optional_)
  * Type: string
  * Docker profile description.
* __type__ (_Required_)
  * Type: string
  * Allowed value : AWS, GCP, Azure, ContainerRegistry
  * Docker profile type.
* __server__ (_Required_)
  * Type: string
  * Docker profile server.
* __user_name__ (_Optional_)
  * Type: string
  * Docker profile user name (required if type == ContainerRegistry).
* __email__ (_Optional_)
  * Type: string
  * Docker profile email (required if type == ContainerRegistry).
* __pwd__ (_Optional_)
  * Type: string
  * Docker profile password (required if type == ContainerRegistry).
* __cloud_creds_id__ (_Optional_)
  * Type: string
  * The cloud creds to import docker profile profile from (required if type == AWS || Type == GCP).
* __credentials__ (_Optional_)
  * Type: string
  * Credentials of the docker profile.