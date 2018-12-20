package resources

import (
	api_operations "github.com/doubret/terraform-provider-xiiot/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTCloudCredsAws() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCloudCredsAws,
		Read:          readCloudCredsAws,
		Update:        updateCloudCredsAws,
		Delete:        deleteCloudCredsAws,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"access_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func getCloudCredsAws(d *schema.ResourceData) *api_models.CloudCreds {
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	credType := "AWS"
	accessKey := d.Get("access_key").(string)
	secretKey := d.Get("secret_key").(string)

	resource := api_models.CloudCreds{
		ID:          d.Id(),
		Name:        &name,
		Description: &description,
		Type:        &credType,
		AwsCredential: &api_models.AWSCredential{
			AccessKey: &accessKey,
			Secret:    &secretKey,
		},
	}

	return &resource
}

func setCloudCredsAws(d *schema.ResourceData, resource *api_models.CloudCreds) {
	d.Set("name", resource.Name)
	d.Set("description", resource.Description)
	d.Set("access_key", resource.AwsCredential.AccessKey)
	d.Set("secret_key", resource.AwsCredential.Secret)
}

func createCloudCredsAws(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	result, err := config.Client.Operations.CloudCredsCreate(api_operations.NewCloudCredsCreateParams().WithBody(getCloudCredsAws(d)), config.Auth)

	if err != nil {
		return err
	}

	d.SetId(*result.Payload.ID)

	return readCloudCredsAws(d, meta)
}

func readCloudCredsAws(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	resource, err := config.Client.Operations.CloudCredsGet(api_operations.NewCloudCredsGetParams().WithID(d.Id()), config.Auth)

	if err != nil {
		return err
	}

	setCloudCredsAws(d, resource.Payload)

	return nil
}

func updateCloudCredsAws(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.CloudCredsUpdateV2(api_operations.NewCloudCredsUpdateV2Params().WithBody(getCloudCredsAws(d)), config.Auth)

	if err != nil {
		return err
	}

	return readCloudCredsAws(d, meta)
}

func deleteCloudCredsAws(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.CloudCredsDelete(api_operations.NewCloudCredsDeleteParams().WithID(d.Id()), config.Auth)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
