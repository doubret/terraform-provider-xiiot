resource "xiiot_user" "admin" {
  email                         = "admin@email.com"
  name                          = "admin"
  password                      = "123"
  role                          = "INFRA_ADMIN"
}

resource "xiiot_user" "user" {
  email                         = "user@email.com"
  name                          = "user"
  password                      = "123"
  role                          = "USER"
}

resource "xiiot_user" "user_without_role" {
  email                         = "user_without_role@email.com"
  name                          = "user_without_role"
  password                      = "123"
}