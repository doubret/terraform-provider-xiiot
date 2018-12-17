package resources

import (
	"log"

	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTAwsCloudCred() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAwsCloudCred,
		Read:          readAwsCloudCred,
		Update:        updateAwsCloudCred,
		Delete:        deleteAwsCloudCred,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func getAwsCloudCred(d *schema.ResourceData) *api_models.CloudCreds {
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

func setAwsCloudCred(d *schema.ResourceData, resource *api_models.CloudCreds) {
	d.Set("name", resource.Name)
	d.Set("description", resource.Description)
	d.Set("access_key", resource.AwsCredential.AccessKey)
	d.Set("secret_key", resource.AwsCredential.Secret)
	d.SetId(resource.ID)
}

func createAwsCloudCred(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createAwsCloudCred")

	config := meta.(configuration.Configuration)

	model := getAwsCloudCred(d)
	_, err := config.Client.Operations.CloudCredsCreate(api_operations.NewCloudCredsCreateParams().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setAwsCloudCred(d, model)

	return nil
}

func readAwsCloudCred(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readAwsCloudCred")

	id := d.Id()

	config := meta.(configuration.Configuration)

	model, err := config.Client.Operations.CloudCredsGet(api_operations.NewCloudCredsGetParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	}

	setAwsCloudCred(d, model.Payload)

	return nil
}

func updateAwsCloudCred(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateAwsCloudCred")

	config := meta.(configuration.Configuration)

	model := getAwsCloudCred(d)
	_, err := config.Client.Operations.CloudCredsUpdateV2(api_operations.NewCloudCredsUpdateV2Params().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to update resource : ", err)

		return err
	}

	setAwsCloudCred(d, model)

	return nil
}

func deleteAwsCloudCred(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteAwsCloudCred")

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
