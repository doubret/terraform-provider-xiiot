package resources

import (
	"log"

	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTContainerRegistry() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createContainerRegistry,
		Read:          readContainerRegistry,
		Update:        updateContainerRegistry,
		Delete:        deleteContainerRegistry,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pwd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_creds_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func getContainerRegistry(d *schema.ResourceData) *api_models.ContainerRegistry {
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	registryType := d.Get("type").(string)
	server := d.Get("server").(string)
	userName := d.Get("user_name").(string)
	email := d.Get("email").(string)
	pwd := d.Get("pwd").(string)
	cloudCredsId := d.Get("cloud_creds_id").(string)

	resource := api_models.ContainerRegistry{
		ID:           d.Id(),
		Name:         &name,
		Description:  description,
		Type:         &registryType,
		Server:       &server,
		UserName:     userName,
		Email:        email,
		Pwd:          pwd,
		CloudCredsID: cloudCredsId,
	}

	return &resource
}

func setContainerRegistry(d *schema.ResourceData, resource *api_models.ContainerRegistry) {
	d.Set("name", resource.Name)
	d.Set("description", resource.Description)
	d.Set("type", resource.Type)
	d.Set("server", resource.Server)
	d.Set("user_name", resource.UserName)
	d.Set("email", resource.Email)
	d.Set("pwd", resource.Pwd)
	d.Set("cloud_creds_id", resource.CloudCredsID)
	d.SetId(resource.ID)
}

func createContainerRegistry(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createContainerRegistry")

	config := meta.(configuration.Configuration)

	model := getContainerRegistry(d)
	_, err := config.Client.Operations.ContainerRegistryCreate(api_operations.NewContainerRegistryCreateParams().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setContainerRegistry(d, model)

	return nil
}

func readContainerRegistry(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readContainerRegistry")

	id := d.Id()

	config := meta.(configuration.Configuration)

	model, err := config.Client.Operations.ContainerRegistryGet(api_operations.NewContainerRegistryGetParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	}

	setContainerRegistry(d, model.Payload)

	return nil
}

func updateContainerRegistry(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateContainerRegistry")

	config := meta.(configuration.Configuration)

	model := getContainerRegistry(d)
	_, err := config.Client.Operations.ContainerRegistryUpdate(api_operations.NewContainerRegistryUpdateParams().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to update resource : ", err)

		return err
	}

	setContainerRegistry(d, model)

	return nil
}

func deleteContainerRegistry(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteContainerRegistry")

	id := d.Id()

	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.ContainerRegistryDelete(api_operations.NewContainerRegistryDeleteParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to delete resource : ", err)

		return err
	}

	d.SetId("")

	return nil
}
