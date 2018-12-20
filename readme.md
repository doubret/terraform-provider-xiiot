# terraform-provider-xiiot

[Terraform](https://www.terraform.io) Provider for Nutanix [Xi IoT](https://www.nutanix.fr/products/iot/)

## Description

This project is a terraform custom provider for Nutanix Xi IoT.

It uses the [Xi IoT REST api](https://iot.nutanix.com/docs/) to manage resources.

The REST api client is automatically generated using [go-swagger](https://github.com/go-swagger/go-swagger).

## Build

For now, building from sources is the only way to get the binaries for this provider.

### Assumption

* You have (some) experience with the Go language and its code organization.

### Build steps

1. Install Go tools from https://golang.org/dl/
2. Check out this repository `git clone https://github.com/doubret/terraform-provider-xiiot.git`
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

1. Install terraform. ([link](https://learn.hashicorp.com/terraform/getting-started/install))
2. Build the provider from sources and copy it in the terraform plugins folder. ([here](https://www.terraform.io/docs/plugins/basics.html))
3. Configure credentials. ([link](docs/provider.md))
4. Run `terraform init` once, then use terraform as usual (`terraform plan`, `terraform apply`, etc...).

## Reference documentation

* [provider](docs/provider.md)
* [application](docs/application.md)
* [category](docs/category.md)
* [cloudcreds aws](docs/cloudcreds_aws.md)
* [cloudcreds gcp](docs/cloudcreds_gcp.md)
* [container registry](docs/container_registry.md)
* [docker profile](docs/docker_profile.md)
* [edge](docs/edge.md)
* [project](docs/project.md)
* [sensor](docs/sensor.md)

## Examples

Resource configuration examples are available in the [sample](sample) folder.

* [xiiot_application](sample/application.tf)
* [xiiot_category](sample/category.tf)
* [xiiot_cloudcreds_aws](sample/cloudcreds_aws.tf)
* [xiiot_cloudcreds_gcp](sample/cloudcreds_gcp.tf)
* [xiiot_containerregistry](sample/containerregistry.tf)
* [xiiot_dockerprofile](sample/dockerprofile.tf)
* [xiiot_edge](sample/edge.tf)
* [xiiot_project](sample/project.tf)
* [xiiot_sensor](sample/sensor.tf)

A provider configuration example is also [available](sample/provider.tf).
