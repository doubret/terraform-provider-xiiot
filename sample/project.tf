resource "xiiot_project" "project" {
  name                          = "name"
  description                   = "description"
  cloud_credential_ids          = [ "${xiiot_cloudcreds_gcp.cloudcredsgcp.id}" ]
  docker_profile_ids            = [ "${xiiot_dockerprofile.dockerprofile.id}" ]
  edge_ids                      = [ "${xiiot_edge.edge.id}" ]
  edge_selector_type            = "Explicit"

  edge_selector {
    id                          = "${xiiot_category.category.id}"
    value                       = "Value1"
  }

  user {
    user_id                     = "${xiiot_user.admin.id}"
    role                        = "PROJECT_ADMIN"
  }

  user {
    user_id                     = "${xiiot_user.user.id}"
    role                        = "PROJECT_USER"
  }
}
