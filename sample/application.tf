resource "xiiot_application" "application" {
  name                          = "Application"
  description                   = "description"
  yaml_data                     = "---"
  project_id                    = "${xiiot_project.project.id}"
  edge_ids                      = [ "${xiiot_edge.edge.id}" ]

  edge_selector {
    id                          = "${xiiot_category.category.id}"
    value                       = "Value1"
  }
}
