package resources

import (
	api_operations "github.com/doubret/terraform-provider-xiiot/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/models"
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
	cloudCredsID := d.Get("cloud_creds_id").(string)
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
		CloudCredsID: cloudCredsID,
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
}

func createDockerProfile(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	result, err := config.Client.Operations.DockerProfileCreate(api_operations.NewDockerProfileCreateParams().WithBody(getDockerProfile(d)), config.Auth)

	if err != nil {
		return err
	}

	d.SetId(*result.Payload.ID)

	return readDockerProfile(d, meta)
}

func readDockerProfile(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	resource, err := config.Client.Operations.DockerProfileGet(api_operations.NewDockerProfileGetParams().WithID(d.Id()), config.Auth)

	if err != nil {
		if httperr, ok := err.(*api_operations.DockerProfileGetDefault); ok {
			if httperr.Code() == 404 {
				d.SetId("")

				return nil
			}
		}

		return err
	}

	setDockerProfile(d, resource.Payload)

	return nil
}

func updateDockerProfile(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.DockerProfileUpdateV2(api_operations.NewDockerProfileUpdateV2Params().WithID(d.Id()).WithBody(getDockerProfile(d)), config.Auth)

	if err != nil {
		return err
	}

	return readDockerProfile(d, meta)
}

func deleteDockerProfile(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.DockerProfileDelete(api_operations.NewDockerProfileDeleteParams().WithID(d.Id()), config.Auth)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
