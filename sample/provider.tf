provider "xiiot" {
  # The recommended way to configure username and password is to use the XI_USER and XI_PASSWORD environment variables
  # You can also set the configuration directly in the provider :
  #   username = "${var.xi_user}"
  #   password = "${var.xi_password}"

  # The endpoint can be configured using the XI_ENDPOINT environment variable or set directly in the configuration
  endpoint = "iot.nutanix.com"
}
