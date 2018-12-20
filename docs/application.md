# terraform-provider-xiiot

## Application reference

Terraform resource name : `xiiot_application`

## Example

```
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
```

## Arguments

* __id__ (_Computed_)
  * Type: string
  * Id of the resource, this is the id for both terraform and Xi IoT.
* __name__ (_Required_)
  * Type: string
  * Application name.
* __description__ (_Optional_)
  * Type: string
  * Application description.
* __yaml_data__ (_Required_)
  * Type: string
  * Application manifest.
* __project_id__ (_Required_)
  * Type: string
  * Id of the project the application is associated with.
* __edge_ids__ (_Optional_)
  * Type: set of strings
  * List of ids of edges this application should be deployed to. Only relevant if parent project EdgeSelectorType = 'Explicit'.
* __edge_selector__ (_Optional_)
  * Type: set of [catgory_info](catgory_info.md)
  * Edge selectors - CategoryInfo list. Only relevant when parent project EdgeSelectorType = 'Category'.
