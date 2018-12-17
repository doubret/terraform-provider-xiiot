package resources

import (
	"log"

	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
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
				Type:     schema.TypeString,
				Required: true,
			},
			"topic_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func getSensor(d *schema.ResourceData) *api_models.Sensor {
	edgeId := d.Get("edge_id").(string)
	topicName := d.Get("topic_name").(string)

	resource := api_models.Sensor{
		ID:        d.Id(),
		EdgeID:    &edgeId,
		TopicName: &topicName,
	}

	return &resource
}

func setSensor(d *schema.ResourceData, resource *api_models.Sensor) {
	d.Set("edge_id", resource.EdgeID)
	d.Set("topic_name", resource.TopicName)
	d.SetId(resource.ID)
}

func createSensor(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createSensor")

	config := meta.(configuration.Configuration)

	model := getSensor(d)
	_, err := config.Client.Operations.SensorCreate(api_operations.NewSensorCreateParams().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setSensor(d, model)

	return nil
}

func readSensor(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readSensor")

	id := d.Id()

	config := meta.(configuration.Configuration)

	model, err := config.Client.Operations.SensorGet(api_operations.NewSensorGetParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	}

	setSensor(d, model.Payload)

	return nil
}

func updateSensor(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateSensor")

	config := meta.(configuration.Configuration)

	model := getSensor(d)
	_, err := config.Client.Operations.SensorUpdateV2(api_operations.NewSensorUpdateV2Params().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to update resource : ", err)

		return err
	}

	setSensor(d, model)

	return nil
}

func deleteSensor(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteSensor")

	id := d.Id()

	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.SensorDelete(api_operations.NewSensorDeleteParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to delete resource : ", err)

		return err
	}

	d.SetId("")

	return nil
}
