# terraform-provider-xiiot

## Cloudcreds GCP reference

Terraform resource name : `xiiot_cloudcredsgcp`

## Example

```
resource "xiiot_cloudcredsgcp" "cloudcredsgcp" {
  name                          = "name"
  description                   = "description"
	auth_provider_x509_cert_url   = "auth_provider_x509_cert_url"
	auth_uri                      = "auth_uri"
	client_email                  = "client_email"
	client_id                     = "client_id"
	client_x509_cert_url          = "client_x509_cert_url"
	private_key                   = "private_key"
	private_key_id                = "private_key_id"
	project_id                    = "project_id"
	token_uri                     = "token_uri"
	type                          = "type"
}
```

## Arguments

* __id__ (_Computed_)
  * Type: string
  * Id of the resource, this is the id for both terraform and Xi IoT
* __name__ (_Required_)
  * Type: string
  * Cloudcreds name
* __description__ (_Optional_)
  * Type: string
  * Cloudcreds description
* __auth_provider_x509_cert_url__ (_Required_)
  * Type: string
* __auth_uri__ (_Required_)
  * Type: string
* __client_email__ (_Required_)
  * Type: string
* __client_id__ (_Required_)
  * Type: string
* __client_x509_cert_url__ (_Required_)
  * Type: string
* __private_key__ (_Required_)
  * Type: string
* __private_key_id__ (_Required_)
  * Type: string
* __project_id__ (_Required_)
  * Type: string
* __token_uri__ (_Required_)
  * Type: string
* __type__ (_Required_)
  * Type: string
