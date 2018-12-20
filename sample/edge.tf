resource "xiiot_edge" "edge" {
  name                          = "name"
  description                   = "description"
  gateway                       = "1.2.3.4"
  ip_address                    = "5.6.7.8"
  serial_number                 = "serial_number"
  subnet                        = "255.255.255.0"

  label {
    id                          = "${xiiot_category.category.id}"
    value                       = "Value1"
  }
}
