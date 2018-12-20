resource "xiiot_user" "admin" {
  name                          = "admin"
  email                         = "admin@email.com"
  password                      = "123"
  role                          = "INFRA_ADMIN"
}

resource "xiiot_user" "user" {
  name                          = "user"
  email                         = "user@email.com"
  password                      = "123"
  role                          = "USER"
}

resource "xiiot_user" "user_without_role" {
  name                          = "user_without_role"
  email                         = "user_without_role@email.com"
  password                      = "123"
}