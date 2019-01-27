package resources

import (
	api_operations "github.com/doubret/terraform-provider-xiiot/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/doubret/terraform-provider-xiiot/xiiot/utils"
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
				Type:             schema.TypeString,
				Optional:         true,
				StateFunc:        utils.Id_to_state,
				DiffSuppressFunc: utils.Compare_ids,
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
	cloudCredsID := d.Get("cloud_creds_id").(string)

	resource := api_models.ContainerRegistry{
		ID:           d.Id(),
		Name:         &name,
		Description:  description,
		Type:         &registryType,
		Server:       &server,
		UserName:     userName,
		Email:        email,
		Pwd:          pwd,
		CloudCredsID: cloudCredsID,
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
}

func createContainerRegistry(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	result, err := config.Client.Operations.ContainerRegistryCreate(api_operations.NewContainerRegistryCreateParams().WithBody(getContainerRegistry(d)), config.Auth)

	if err != nil {
		return err
	}

	d.SetId(*result.Payload.ID)

	return readContainerRegistry(d, meta)
}

func readContainerRegistry(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	resource, err := config.Client.Operations.ContainerRegistryGet(api_operations.NewContainerRegistryGetParams().WithID(d.Id()), config.Auth)

	if err != nil {
		if httperr, ok := err.(*api_operations.ContainerRegistryGetDefault); ok {
			if httperr.Code() == 404 {
				d.SetId("")

				return nil
			}
		}

		return err
	}

	setContainerRegistry(d, resource.Payload)

	return nil
}

func updateContainerRegistry(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.ContainerRegistryUpdate(api_operations.NewContainerRegistryUpdateParams().WithID(d.Id()).WithBody(getContainerRegistry(d)), config.Auth)

	if err != nil {
		return err
	}

	return readContainerRegistry(d, meta)
}

func deleteContainerRegistry(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.ContainerRegistryDelete(api_operations.NewContainerRegistryDeleteParams().WithID(d.Id()), config.Auth)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
