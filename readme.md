# terraform-provider-xiiot

[Terraform](https://www.terraform.io) Provider for [Nutanix Xi IoT](https://www.nutanix.fr/products/iot/)

## Description

This project is a terraform custom provider for Nutanix Xi IoT. It uses the REST api manage resources.

The REST api client is automatically generated using [go-swagger](https://github.com/go-swagger/go-swagger).

## Requirement

* [hashicorp/terraform](https://github.com/hashicorp/terraform)

## Usage

### Install terraform

Install terraform as explained [here](https://learn.hashicorp.com/terraform/getting-started/install)

### Install terraform-provider-xiiot (on windows)

* Create folder `%APPDATA%/terraform.d/plugins/` if it doesn't exist yet.
* Copy `terraform-provider-xiiot.exe` in the `%APPDATA%/terraform.d/plugins/` folder.

Alternatively, you can simply run `plugin.bat` from this repository.

### Configure credentials

The safer way to configure credentials is to use environment variables. This limits the risks of accidentaly pushing your credentials in your scm repository.

On windows:

```
set XI_USER=email@domain.com
set XI_PASSWORD=password

REM you can customize Xi IoT REST api endpoint if needed (defaults to iot.nutanix.com if not specified)
set XI_ENDPOINT=xi.endpoint.com
```

On linux & co:

```
export XI_USER=email@domain.com
export XI_PASSWORD=password

# you can customize Xi IoT REST api endpoint if needed (defaults to iot.nutanix.com if not specified)
export XI_ENDPOINT=xi.endpoint.com
```

### Use terraform as usual

Run `terraform init` once the use `terraform plan`, `terraform apply`, etc... as usual.

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

* `username` - (Required) This is the user name to access the Xi IoT REST api. Can be specified in environment variable `XI_USER`.
* `password` - (Required) This is the password to access the Xi IoT REST api. Can be specified in environment variable `XI_PASSWORD`.
* `endpoint` - (Optionnal) Xi IoT REST api endpoint. Can be specified in environment variable `XI_ENDPOINT`, if not specified, default to `iot.nutanix.com`.

## Build from sources

TODO

