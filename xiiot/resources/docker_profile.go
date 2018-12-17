package resources

import (
	"log"

	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTDockerProfile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDockerProfile,
		Read:          readDockerProfile,
		Update:        updateDockerProfile,
		Delete:        deleteDockerProfile,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"credentials": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func getDockerProfile(d *schema.ResourceData) *api_models.DockerProfile {
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	registryType := d.Get("type").(string)
	server := d.Get("server").(string)
	userName := d.Get("user_name").(string)
	email := d.Get("email").(string)
	pwd := d.Get("pwd").(string)
	cloudCredsId := d.Get("cloud_creds_id").(string)
	credentials := d.Get("credentials").(string)

	resource := api_models.DockerProfile{
		ID:           d.Id(),
		Name:         &name,
		Description:  description,
		Type:         &registryType,
		Server:       &server,
		UserName:     userName,
		Email:        email,
		Pwd:          pwd,
		CloudCredsID: cloudCredsId,
		Credentials:  credentials,
	}

	return &resource
}

func setDockerProfile(d *schema.ResourceData, resource *api_models.DockerProfile) {
	d.Set("name", resource.Name)
	d.Set("description", resource.Description)
	d.Set("type", resource.Type)
	d.Set("server", resource.Server)
	d.Set("user_name", resource.UserName)
	d.Set("email", resource.Email)
	d.Set("pwd", resource.Pwd)
	d.Set("cloud_creds_id", resource.CloudCredsID)
	d.Set("credentials", resource.Credentials)
	d.SetId(resource.ID)
}

func createDockerProfile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createDockerProfile")

	config := meta.(configuration.Configuration)

	model := getDockerProfile(d)
	_, err := config.Client.Operations.DockerProfileCreate(api_operations.NewDockerProfileCreateParams().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setDockerProfile(d, model)

	return nil
}

func readDockerProfile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readDockerProfile")

	id := d.Id()

	config := meta.(configuration.Configuration)

	model, err := config.Client.Operations.DockerProfileGet(api_operations.NewDockerProfileGetParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	}

	setDockerProfile(d, model.Payload)

	return nil
}

func updateDockerProfile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateDockerProfile")

	config := meta.(configuration.Configuration)

	model := getDockerProfile(d)
	_, err := config.Client.Operations.DockerProfileUpdateV2(api_operations.NewDockerProfileUpdateV2Params().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to update resource : ", err)

		return err
	}

	setDockerProfile(d, model)

	return nil
}

func deleteDockerProfile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteDockerProfile")

	id := d.Id()

	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.DockerProfileDelete(api_operations.NewDockerProfileDeleteParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to delete resource : ", err)

		return err
	}

	d.SetId("")

	return nil
}
