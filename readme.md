# terraform-provider-xiiot

[Terraform](https://www.terraform.io) Provider for Nutanix [Xi IoT](https://www.nutanix.fr/products/iot/)

## Description

This project is a terraform custom provider for Nutanix Xi IoT.
It uses the Xi IoT REST api to manage resources.

The REST api client is automatically generated using [go-swagger](https://github.com/go-swagger/go-swagger).

## Build

For now, building from sources is the only way to get the binaries for this provider.

### Assumption

* You have (some) experience with the Go language and its code organization.

### Build steps

1. Install Go tools from https://golang.org/dl/
2. Check out this repository: `git clone https://github.com/doubret/terraform-provider-xiiot.git`
3. Get dependencies
```
go get github.com/hashicorp/terraform
go get github.com/go-swagger/go-swagger
go get github.com/go-openapi/errors
go get github.com/go-openapi/runtime
go get github.com/go-openapi/runtime/client
go get github.com/go-openapi/strfmt
```
4. Build go-swagger generator
```
go install github.com/go-swagger/go-swagger/cmd/swagger
```
5. Generate api client
```
swagger generate client -f https://iot.nutanix.com/swagger.json -A xiClient
```
6. Build with `go build` or `go install`

## Usage

1. Install terraform as explained [here](https://learn.hashicorp.com/terraform/getting-started/install)
2. Build the provider from sources and copy it in the terraform plugins folder.

2.1. On windows :
* Create folder `%APPDATA%/terraform.d/plugins/` if it doesn't exist yet.
* Copy `terraform-provider-xiiot.exe` in the `%APPDATA%/terraform.d/plugins/` folder.

2.2. On linux :
* TODO

3. Configure credentials as explained in the [configuration example](sample/provider.tf)
4. Run `terraform init` once, then use terraform as usual (`terraform plan`, `terraform apply`, etc...)

## Examples

Resource configuration examples are available in the [sample](sample) folder.

* [xiiot_application](sample/application.tf)
* [xiiot_category](sample/category.tf)
* [xiiot_cloudcreds](sample/cloudcreds.tf)
* [xiiot_containerregistry](sample/containerregistry.tf)
* [xiiot_dockerprofile](sample/dockerprofile.tf)
* [xiiot_edge](sample/edge.tf)
* [xiiot_project](sample/project.tf)
* [xiiot_sensor](sample/sensor.tf)

A provider configuration example is also [available](sample/provider.tf).

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
* `endpoint` - (Optional) Xi IoT REST api endpoint. Can be specified in environment variable `XI_ENDPOINT`, if not specified, default to `iot.nutanix.com`.

## Resources reference

Not supported (yet) :
* scripts
* scriptruntimes
* datasources
* datastreams

Supported but not documented (yet) :
* projects
* applications
* edges

### Categories

```
resource "xiiot_category" "category_name" {
  name            = "name"
  purpose         = "purpose"
  values          = [ "1", "2", "3" ]
}
```

### Users

```
resource "xiiot_user" "user" {
  email           = "email"
  name            = "name"
  password        = "sha-256"
  role            = "USER"
}
```

### Docker profiles

```
resource "xiiot_dockerprofile" "dockerprofile" {
  name            = "name"
  description     = "description"
  type            = "type"
  server          = "server"
  user_name       = "user_name"
  email           = "email"
  pwd             = "pwd"
  cloud_creds_id  = "cloud_creds_id"
  credentials     = "credentials"
}
```

### Container registries

```
resource "xiiot_containerregistry" "containerregistry" {
  name            = "name"
  description     = "description"
  type            = "type"
  server          = "server"
  user_name       = "user_name"
  email           = "email"
  pwd             = "pwd"
  cloud_creds_id  = "cloud_creds_id"
}
```

### Sensors

```
resource "xiiot_sensor" "sensor" {
  edge_id         = "edge_id"
  topic_name      = "topic_name"
}
```

### Cloud credentials

```
resource "xiiot_cloudcredaws" "credential" {
  name            = "name"
  description     = "description"
  access_key      = "access_key"
  secret_key      = "secret_key"
}
```

```
resource "xiiot_cloudcredgcp" "credential" {
  name                        = "name"
  description                 = "description"
	auth_provider_x509_cert_url = "auth_provider_x509_cert_url"
	auth_uri                    = "auth_uri"
	client_email                = "client_email"
	client_id                   = "client_id"
	client_x509_cert_url        = "client_x509_cert_url"
	private_key                 = "private_key"
	private_key_id              = "private_key_id"
	project_id                  = "project_id"
	token_uri                   = "token_uri"
	type                        = "type"
}
```
