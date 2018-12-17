package resources

import (
	"log"

	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/doubret/terraform-provider-xiiot/xiiot/utils"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTApplication() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createApplication,
		Read:          readApplication,
		Update:        updateApplication,
		Delete:        deleteApplication,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"yaml_data": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"edge_ids": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"edge_selectors": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Optional: true,
			},
		},
	}
}

func getApplication(d *schema.ResourceData) *api_models.Application {
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	yamlData := d.Get("yaml_data").(string)
	projectId := d.Get("project_id").(string)

	resource := api_models.Application{
		ID:          d.Id(),
		Name:        &name,
		Description: description,
		YamlData:    &yamlData,
		ProjectID:   projectId,
		EdgeIds:     utils.Convert_set_to_string_array(d.Get("edge_ids").(*schema.Set)),
		// EdgeSelectors
	}

	return &resource
}

func setApplication(d *schema.ResourceData, resource *api_models.Application) {
	d.Set("name", resource.Name)
	d.Set("description", resource.Description)
	d.Set("yaml_data", resource.YamlData)
	d.Set("project_id", resource.ProjectID)
	d.Set("edge_ids", resource.EdgeIds)
	// EdgeSelectors
	d.SetId(resource.ID)
}

func createApplication(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createApplication")

	config := meta.(configuration.Configuration)

	model := getApplication(d)
	_, err := config.Client.Operations.ApplicationCreate(api_operations.NewApplicationCreateParams().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setApplication(d, model)

	return nil
}

func readApplication(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readApplication")

	id := d.Id()

	config := meta.(configuration.Configuration)

	model, err := config.Client.Operations.ApplicationGet(api_operations.NewApplicationGetParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to read resource : ", err)

		return err
	}

	setApplication(d, model.Payload)

	return nil
}

func updateApplication(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateApplication")

	config := meta.(configuration.Configuration)

	model := getApplication(d)
	_, err := config.Client.Operations.ApplicationUpdateV2(api_operations.NewApplicationUpdateV2Params().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to update resource : ", err)

		return err
	}

	setApplication(d, model)

	return nil
}

func deleteApplication(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteApplication")

	id := d.Id()

	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.ApplicationDelete(api_operations.NewApplicationDeleteParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to delete resource : ", err)

		return err
	}

	d.SetId("")

	return nil
}
