package resources

import (
	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTCloudCredsGcp() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCloudCredsGcp,
		Read:          readCloudCredsGcp,
		Update:        updateCloudCredsGcp,
		Delete:        deleteCloudCredsGcp,
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

func getCloudCredsGcp(d *schema.ResourceData) *api_models.CloudCreds {
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	credType := "GCP"
	authProviderX509CertURL := d.Get("auth_provider_x509_cert_url").(string)
	authURI := d.Get("auth_uri").(string)
	clientEmail := d.Get("client_email").(string)
	clientID := d.Get("client_id").(string)
	clientX509CertURL := d.Get("client_x509_cert_url").(string)
	privateKey := d.Get("private_key").(string)
	privateKeyID := d.Get("private_key_id").(string)
	projectID := d.Get("project_id").(string)
	tokenURI := d.Get("token_uri").(string)
	gcpType := d.Get("type").(string)

	resource := api_models.CloudCreds{
		ID:          d.Id(),
		Name:        &name,
		Description: &description,
		Type:        &credType,
		GcpCredential: &api_models.GCPCredential{
			AuthProviderX509CertURL: &authProviderX509CertURL,
			AuthURI:                 &authURI,
			ClientEmail:             &clientEmail,
			ClientID:                &clientID,
			ClientX509CertURL:       &clientX509CertURL,
			PrivateKey:              &privateKey,
			PrivateKeyID:            &privateKeyID,
			ProjectID:               &projectID,
			TokenURI:                &tokenURI,
			Type:                    &gcpType,
		},
	}

	return &resource
}

func setCloudCredsGcp(d *schema.ResourceData, resource *api_models.CloudCreds) {
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
}

func createCloudCredsGcp(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	result, err := config.Client.Operations.CloudCredsCreate(api_operations.NewCloudCredsCreateParams().WithBody(getCloudCredsGcp(d)), config.Auth)

	if err != nil {
		return err
	}

	d.SetId(*result.Payload.ID)

	return readCloudCredsGcp(d, meta)
}

func readCloudCredsGcp(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	model, err := config.Client.Operations.CloudCredsGet(api_operations.NewCloudCredsGetParams().WithID(d.Id()), config.Auth)

	if err != nil {
		return err
	}

	setCloudCredsGcp(d, model.Payload)

	return nil
}

func updateCloudCredsGcp(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.CloudCredsUpdateV2(api_operations.NewCloudCredsUpdateV2Params().WithBody(getCloudCredsGcp(d)), config.Auth)

	if err != nil {
		return err
	}

	return readCloudCredsGcp(d, meta)
}

func deleteCloudCredsGcp(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.CloudCredsDelete(api_operations.NewCloudCredsDeleteParams().WithID(d.Id()), config.Auth)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
