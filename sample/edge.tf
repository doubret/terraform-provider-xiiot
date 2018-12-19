resource "xiiot_edge" "edge" {
  name                          = "name"
  description                   = "description"
  edge_devices                  = 1
  gateway                       = "1.2.3.4"
  ip_address                    = "5.6.7.8"
  serial_number                 = "serial_number"
  storage_capacity              = 128
  storage_usage                 = 64
  subnet                        = "255.255.255.0"
  connected                     = false

  label {
    id                          = "${xiiot_category.category.id}"
    value                       = "Value1"
  }
}
