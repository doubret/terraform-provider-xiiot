# terraform-provider-xiiot

## Project reference

Terraform resource name : `xiiot_project`

## Example

```
resource "xiiot_project" "project" {
  name                          = "name"
  description                   = "description"
  cloud_credential_ids          = [ "${xiiot_cloudcredsgcp.cloudcredsgcp.id}" ]
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
```

## Arguments

* __id__ (_Computed_)
  * Type: string
  * Id of the resource, this is the id for both terraform and Xi IoT.
* __name__ (_Required_)
  * Type: string
  * Project name.
* __description__ (_Required_)
  * Type: string
  * Project description.
* __cloud_credential_ids__ (_Required_)
  * Type: set of strings
  * List of ids of cloud credentials this project has access to.
* __docker_profile_ids__ (_Required_)
  * Type: set of strings
  * List of ids of docker profiles this project has access to.
* __edge_ids__ (_Optional_)
  * Type: set of strings
  * List of ids of edges belong to this project. Only relevant when edge selector type === 'Explicit'.
* __edge_selector_type__ (_Required_)
  * Type: string
  * Allowed values : Category, Explicit
  * Type of edge selector. Either 'Category' or 'Explicit' Specify whether edges belonging to this project are given by edgeIds ('Explicit') or edgeSelectors ('Category').
* __edge_selector__ (_Optional_)
  * Type: set of [catgory_info](category_info.md)
  * Only relevant when parent project edge selector type = 'Category'.
