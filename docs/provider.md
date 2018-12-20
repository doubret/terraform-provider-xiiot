# terraform-provider-xiiot

## Provider reference

Terraform resource name : `xiiot`

## Example

```
provider "xiiot" {
  # The recommended way to configure username and password is to use the XI_USER and XI_PASSWORD environment variables.
  # You can also set the configuration directly in the provider :
  #   username = "${var.xi_user}"
  #   password = "${var.xi_password}"

  # The endpoint can be configured using the XI_ENDPOINT environment variable or set directly in the configuration.
  endpoint = "iot.nutanix.com"
}
```

## Arguments

* __username__ (_Required_)
  * Type: string
  * Username used for authenification against the Xi IoT REST api.
  * If not specified, terraform will attempt to retrieve the value from the `XI_USER` environment variable.
* __password__ (_Required_)
  * Type: string
  * Password used for authenification against the Xi IoT REST api.
  * If not specified, terraform will attempt to retrieve the value from the `XI_PASSWORD` environment variable.
* __endpoint__ (_Optional_)
  * Type: string
  * Endpoint used to access the Xi IoT REST api.
  * If not specified, terraform will attempt to retrieve the value from the `XI_ENDPOINT` environment variable.
  * If not specified and the `XI_ENDPOINT` environment variable is not defined, the default value _iot.nutanix.com_ will be used.
