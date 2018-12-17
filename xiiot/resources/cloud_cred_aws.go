package resources

import (
	"log"

	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTCloudCredAws() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCloudCredAws,
		Read:          readCloudCredAws,
		Update:        updateCloudCredAws,
		Delete:        deleteCloudCredAws,
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

func getCloudCredAws(d *schema.ResourceData) *api_models.CloudCreds {
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

func setCloudCredAws(d *schema.ResourceData, resource *api_models.CloudCreds) {
	d.Set("name", resource.Name)
	d.Set("description", resource.Description)
	d.Set("access_key", resource.AwsCredential.AccessKey)
	d.Set("secret_key", resource.AwsCredential.Secret)
	d.SetId(resource.ID)
}

func createCloudCredAws(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createCloudCredAws")

	config := meta.(configuration.Configuration)

	model := getCloudCredAws(d)
	_, err := config.Client.Operations.CloudCredsCreate(api_operations.NewCloudCredsCreateParams().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setCloudCredAws(d, model)

	return nil
}

func readCloudCredAws(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readCloudCredAws")

	id := d.Id()

	config := meta.(configuration.Configuration)

	model, err := config.Client.Operations.CloudCredsGet(api_operations.NewCloudCredsGetParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	}

	setCloudCredAws(d, model.Payload)

	return nil
}

func updateCloudCredAws(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateCloudCredAws")

	config := meta.(configuration.Configuration)

	model := getCloudCredAws(d)
	_, err := config.Client.Operations.CloudCredsUpdateV2(api_operations.NewCloudCredsUpdateV2Params().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to update resource : ", err)

		return err
	}

	setCloudCredAws(d, model)

	return nil
}

func deleteCloudCredAws(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteCloudCredAws")

	id := d.Id()

	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.CloudCredsDelete(api_operations.NewCloudCredsDeleteParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to delete resource : ", err)

		return err
	}

	d.SetId("")

	return nil
}
