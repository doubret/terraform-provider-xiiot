package resources

import (
	"log"

	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTCloudCredGcp() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCloudCredGcp,
		Read:          readCloudCredGcp,
		Update:        updateCloudCredGcp,
		Delete:        deleteCloudCredGcp,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_provider_x509_cert_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"auth_uri": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_x509_cert_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"private_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"private_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"token_uri": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func getCloudCredGcp(d *schema.ResourceData) *api_models.CloudCreds {
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	credType := "GCP"
	authProviderX509CertURL := d.Get("auth_provider_x509_cert_url").(string)
	authUri := d.Get("auth_uri").(string)
	clientEmail := d.Get("client_email").(string)
	clientId := d.Get("client_id").(string)
	clientX509CertUrl := d.Get("client_x509_cert_url").(string)
	privateKey := d.Get("private_key").(string)
	privateKeyId := d.Get("private_key_id").(string)
	projectId := d.Get("project_id").(string)
	tokenUri := d.Get("token_uri").(string)
	gcpType := d.Get("type").(string)

	resource := api_models.CloudCreds{
		ID:          d.Id(),
		Name:        &name,
		Description: &description,
		Type:        &credType,
		GcpCredential: &api_models.GCPCredential{
			AuthProviderX509CertURL: &authProviderX509CertURL,
			AuthURI:                 &authUri,
			ClientEmail:             &clientEmail,
			ClientID:                &clientId,
			ClientX509CertURL:       &clientX509CertUrl,
			PrivateKey:              &privateKey,
			PrivateKeyID:            &privateKeyId,
			ProjectID:               &projectId,
			TokenURI:                &tokenUri,
			Type:                    &gcpType,
		},
	}

	return &resource
}

func setCloudCredGcp(d *schema.ResourceData, resource *api_models.CloudCreds) {
	d.Set("name", resource.Name)
	d.Set("description", resource.Description)
	d.Set("auth_provider_x509_cert_url", resource.GcpCredential.AuthProviderX509CertURL)
	d.Set("auth_uri", resource.GcpCredential.AuthURI)
	d.Set("client_email", resource.GcpCredential.ClientEmail)
	d.Set("client_id", resource.GcpCredential.ClientID)
	d.Set("client_x509_cert_url", resource.GcpCredential.ClientX509CertURL)
	d.Set("private_key", resource.GcpCredential.PrivateKey)
	d.Set("private_key_id", resource.GcpCredential.PrivateKeyID)
	d.Set("project_id", resource.GcpCredential.ProjectID)
	d.Set("token_uri", resource.GcpCredential.TokenURI)
	d.Set("type", resource.GcpCredential.Type)
	d.SetId(resource.ID)
}

func createCloudCredGcp(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createCloudCredGcp")

	config := meta.(configuration.Configuration)

	model := getCloudCredGcp(d)
	_, err := config.Client.Operations.CloudCredsCreate(api_operations.NewCloudCredsCreateParams().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setCloudCredGcp(d, model)

	return nil
}

func readCloudCredGcp(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readCloudCredGcp")

	id := d.Id()

	config := meta.(configuration.Configuration)

	model, err := config.Client.Operations.CloudCredsGet(api_operations.NewCloudCredsGetParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	}

	setCloudCredGcp(d, model.Payload)

	return nil
}

func updateCloudCredGcp(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateCloudCredGcp")

	config := meta.(configuration.Configuration)

	model := getCloudCredGcp(d)
	_, err := config.Client.Operations.CloudCredsUpdateV2(api_operations.NewCloudCredsUpdateV2Params().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to update resource : ", err)

		return err
	}

	setCloudCredGcp(d, model)

	return nil
}

func deleteCloudCredGcp(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteCloudCredGcp")

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
