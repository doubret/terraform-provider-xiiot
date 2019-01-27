package resources

import (
	api_operations "github.com/doubret/terraform-provider-xiiot/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/doubret/terraform-provider-xiiot/xiiot/utils"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTSensor() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSensor,
		Read:          readSensor,
		Update:        updateSensor,
		Delete:        deleteSensor,
		Schema: map[string]*schema.Schema{
			"edge_id": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				StateFunc:        utils.Id_to_state,
				DiffSuppressFunc: utils.Compare_ids,
			},
			"topic_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func getSensor(d *schema.ResourceData) *api_models.Sensor {
	edgeID := d.Get("edge_id").(string)
	topicName := d.Get("topic_name").(string)

	resource := api_models.Sensor{
		ID:        d.Id(),
		EdgeID:    &edgeID,
		TopicName: &topicName,
	}

	return &resource
}

func setSensor(d *schema.ResourceData, resource *api_models.Sensor) {
	d.Set("edge_id", resource.EdgeID)
	d.Set("topic_name", resource.TopicName)
}

func createSensor(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	result, err := config.Client.Operations.SensorCreate(api_operations.NewSensorCreateParams().WithBody(getSensor(d)), config.Auth)

	if err != nil {
		return err
	}

	d.SetId(*result.Payload.ID)

	return readSensor(d, meta)
}

func readSensor(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	resource, err := config.Client.Operations.SensorGet(api_operations.NewSensorGetParams().WithID(d.Id()), config.Auth)

	if err != nil {
		if httperr, ok := err.(*api_operations.SensorGetDefault); ok {
			if httperr.Code() == 404 {
				d.SetId("")

				return nil
			}
		}

		return err
	}

	setSensor(d, resource.Payload)

	return nil
}

func updateSensor(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.SensorUpdateV2(api_operations.NewSensorUpdateV2Params().WithID(d.Id()).WithBody(getSensor(d)), config.Auth)

	if err != nil {
		return err
	}

	return readSensor(d, meta)
}

func deleteSensor(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.SensorDelete(api_operations.NewSensorDeleteParams().WithID(d.Id()), config.Auth)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
