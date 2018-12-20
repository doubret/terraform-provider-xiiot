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
