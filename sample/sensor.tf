resource "xiiot_sensor" "sensor" {
  edge_id                       = "${xiiot_edge.edge.id}"
  topic_name                    = "topic_name"
}
