# terraform-provider-xiiot

[Terraform](https://www.terraform.io) Provider for [Nutanix Xi IoT](https://www.nutanix.fr/products/iot/)

## Description

This project is a terraform custom provider for Nutanix Xi IoT. It uses the REST api manage resources.

The REST api client is automatically generated using [go-swagger](https://github.com/go-swagger/go-swagger).

## Requirement

* [hashicorp/terraform](https://github.com/hashicorp/terraform)

## Usage

TODO

## Provider configuration

```
provider "xiiot" {
    username = "${var.xi_user}"
    password = "${var.xi_password}"
    endpoint = "${var.xi_endpoint}"
}
```

### Arguments reference

The following arguments are supported.

* `username` - (Required) This is the user name to access the Xi IoT REST api. Can be specified in environment variable `XI_LOGIN`.
* `password` - (Required) This is the password to access the Xi IoT REST api. Can be specified in environment variable `XI_PASSWORD`.
* `endpoint` - (Required) Xi IoT REST api endpoint. Can be specified in environment variable `XI_ENDPOINT`.

## Build from sources

TODO

