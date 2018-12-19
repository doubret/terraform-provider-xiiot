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
