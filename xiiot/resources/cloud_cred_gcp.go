package resources

import (
	"log"

	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTGcpCloudCred() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGcpCloudCred,
		Read:          readGcpCloudCred,
		Update:        updateGcpCloudCred,
		Delete:        deleteGcpCloudCred,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func getGcpCloudCred(d *schema.ResourceData) *api_models.CloudCreds {
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	credType := "GCP"
	authProviderX509CertURL := d.Get("auth_provider_x509_cert_url").(string)
	auth_uri := d.Get("auth_uri").(string)
	client_email := d.Get("client_email").(string)
	client_id := d.Get("client_id").(string)
	client_x509_cert_url := d.Get("client_x509_cert_url").(string)
	private_key := d.Get("private_key").(string)
	private_key_id := d.Get("private_key_id").(string)
	project_id := d.Get("project_id").(string)
	token_uri := d.Get("token_uri").(string)
	gcpType := d.Get("type").(string)

	resource := api_models.CloudCreds{
		ID:          d.Id(),
		Name:        &name,
		Description: &description,
		Type:        &credType,
		GcpCredential: &api_models.GCPCredential{
			AuthProviderX509CertURL: &authProviderX509CertURL,
			AuthURI:                 &auth_uri,
			ClientEmail:             &client_email,
			ClientID:                &client_id,
			ClientX509CertURL:       &client_x509_cert_url,
			PrivateKey:              &private_key,
			PrivateKeyID:            &private_key_id,
			ProjectID:               &project_id,
			TokenURI:                &token_uri,
			Type:                    &gcpType,
		},
	}

	return &resource
}

func setGcpCloudCred(d *schema.ResourceData, resource *api_models.CloudCreds) {
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

func createGcpCloudCred(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createGcpCloudCred")

	config := meta.(configuration.Configuration)

	model := getGcpCloudCred(d)
	_, err := config.Client.Operations.CloudCredsCreate(api_operations.NewCloudCredsCreateParams().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setGcpCloudCred(d, model)

	return nil
}

func readGcpCloudCred(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readGcpCloudCred")

	id := d.Id()

	config := meta.(configuration.Configuration)

	model, err := config.Client.Operations.CloudCredsGet(api_operations.NewCloudCredsGetParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	}

	setGcpCloudCred(d, model.Payload)

	return nil
}

func updateGcpCloudCred(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateGcpCloudCred")

	config := meta.(configuration.Configuration)

	model := getGcpCloudCred(d)
	_, err := config.Client.Operations.CloudCredsUpdateV2(api_operations.NewCloudCredsUpdateV2Params().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to update resource : ", err)

		return err
	}

	setGcpCloudCred(d, model)

	return nil
}

func deleteGcpCloudCred(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteGcpCloudCred")

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
