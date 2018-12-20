# terraform-provider-xiiot

## Sensor reference

Terraform resource name : `xiiot_sensor`

## Example

```
resource "xiiot_sensor" "sensor" {
  edge_id                       = "${xiiot_edge.edge.id}"
  topic_name                    = "topic_name"
}
```

## Arguments

* __id__ (_Computed_)
  * Type: string
  * Id of the resource, this is the id for both terraform and Xi IoT
* __edge_id__ (_Required_)
  * Type: string
  * Id of the edge
* __topic_name__ (_Required_)
  * Type: string
  * Name of the mqtt topic
