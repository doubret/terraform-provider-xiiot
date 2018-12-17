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
				Optional: false,
				Computed: false,
				ForceNew: false,
			},
			"topic_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Optional: false,
				Computed: false,
				ForceNew: false,
			},
		},
	}
}

func getSensor(d *schema.ResourceData) *api_models.Sensor {
	// name := d.Get("name").(string)
	// purpose := d.Get("purpose").(string)

	resource := api_models.Sensor{
		// ID:      d.Id(),
		// Name:    &name,
		// Purpose: &purpose,
		// Values:  utils.Convert_set_to_string_array(d.Get("values").(*schema.Set)),
	}

	return &resource
}

func setSensor(d *schema.ResourceData, resource *api_models.Sensor) {
	// d.Set("xi_id", resource.ID)
	// d.Set("name", resource.Name)
	// d.Set("purpose", resource.Purpose)
	// d.Set("values", resource.Values)
	// d.SetId(resource.ID)
}

func createSensor(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createSensor")

	read := readSensor(d, meta)

	if read != nil {
		config := meta.(configuration.Configuration)

		tfResource := getSensor(d)
		_, err := config.Client.Operations.SensorCreate(api_operations.NewSensorCreateParams().WithBody(tfResource), config.Auth)

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		// TODO : should read back ?

		setSensor(d, tfResource)
	}

	return nil
}

func readSensor(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readSensor")

	id := d.Id()

	if id == "" {
		id = d.Get("xi_id").(string)
	}

	if id == "" {
		d.SetId("")
	} else {
		config := meta.(configuration.Configuration)

		xiResource, err := config.Client.Operations.SensorGet(api_operations.NewSensorGetParams().WithID(id), config.Auth)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		setSensor(d, xiResource.Payload)
	}

	return nil
}

func updateSensor(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateSensor")

	config := meta.(configuration.Configuration)

	tfResource := getSensor(d)
	_, err := config.Client.Operations.SensorCreate(api_operations.NewSensorCreateParams().WithBody(tfResource), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setSensor(d, tfResource)

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
