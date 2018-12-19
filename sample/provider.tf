provider "xiiot" {
  # the recommended way to configure username and password is to use the XI_USER and XI_PASSWORD environment variables
  # the endpoint can be configured using the XI_ENDPOINT environment variable
  /* you can also set the configuration directly in the provider like this :
  username = "${var.xi_user}"
  password = "${var.xi_password}"
  endpoint = "${var.xi_endpoint}"
  */
  endpoint = "iot.nutanix.com"
}
