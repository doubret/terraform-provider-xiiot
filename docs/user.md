# terraform-provider-xiiot

## User reference

Terraform resource name : `xiiot_user`

## Example

```
resource "xiiot_user" "admin" {
  name                          = "admin"
  email                         = "admin@email.com"
  password                      = "123"
  role                          = "INFRA_ADMIN"
}
```

## Arguments

* __id__ (_Computed_)
  * Type: string
  * Id of the resource, this is the id for both terraform and Xi IoT.
* __name__ (_Required_)
  * Type: string
  * User name.
* __email__ (_Required_)
  * Type: string
  * User email.
* __password__ (_Required_)
  * Type: string
  * User password.
  * TODO : say something about the clear text vs sha-256
* __role__ (_Optional_)
  * Type: string
  * Allowed values: INFRA_ADMIN, USER
  * User role.
