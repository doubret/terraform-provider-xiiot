package resources

import (
	"log"

	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTCloudCreds() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCloudCreds,
		Read:          readCloudCreds,
		Update:        updateCloudCreds,
		Delete:        deleteCloudCreds,
		Schema: map[string]*schema.Schema{
			"xi_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Optional: false,
				Computed: false,
				ForceNew: false,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Optional: false,
				Computed: false,
				ForceNew: false,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Optional: false,
				Computed: false,
				ForceNew: false,
			},
			"aws_credential": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: false,
				Optional: true,
				Computed: false,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_key": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Optional: false,
							Computed: false,
							ForceNew: false,
						},
						"secret_key": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Optional: false,
							Computed: false,
							ForceNew: false,
						},
					},
				},
			},
			"gcp_credential": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: false,
				Optional: true,
				Computed: false,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_provider_x509_cert_url": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Optional: false,
							Computed: false,
							ForceNew: false,
						},
						"auth_uri": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Optional: false,
							Computed: false,
							ForceNew: false,
						},
						"client_email": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Optional: false,
							Computed: false,
							ForceNew: false,
						},
						"client_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Optional: false,
							Computed: false,
							ForceNew: false,
						},
						"client_x509_cert_url": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Optional: false,
							Computed: false,
							ForceNew: false,
						},
						"private_key": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Optional: false,
							Computed: false,
							ForceNew: false,
						},
						"private_key_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Optional: false,
							Computed: false,
							ForceNew: false,
						},
						"project_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Optional: false,
							Computed: false,
							ForceNew: false,
						},
						"token_uri": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Optional: false,
							Computed: false,
							ForceNew: false,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Optional: false,
							Computed: false,
							ForceNew: false,
						},
					},
				},
			},
		},
	}
}

func getCloudCreds(d *schema.ResourceData) *api_models.CloudCreds {
	// name := d.Get("name").(string)
	// purpose := d.Get("purpose").(string)

	resource := api_models.CloudCreds{
		// ID:      d.Id(),
		// Name:    &name,
		// Purpose: &purpose,
		// Values:  utils.Convert_set_to_string_array(d.Get("values").(*schema.Set)),
	}

	return &resource
}

func setCloudCreds(d *schema.ResourceData, resource *api_models.CloudCreds) {
	// d.Set("xi_id", resource.ID)
	// d.Set("name", resource.Name)
	// d.Set("purpose", resource.Purpose)
	// d.Set("values", resource.Values)
	d.SetId(resource.ID)
}

func createCloudCreds(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createCloudCreds")

	read := readCloudCreds(d, meta)

	if read != nil {
		config := meta.(configuration.Configuration)

		tfResource := getCloudCreds(d)
		_, err := config.Client.Operations.CloudCredsCreate(api_operations.NewCloudCredsCreateParams().WithBody(tfResource), config.Auth)

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		// TODO : should read back ?

		setCloudCreds(d, tfResource)
	}

	return nil
}

func readCloudCreds(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readCloudCreds")

	id := d.Id()

	if id == "" {
		id = d.Get("xi_id").(string)
	}

	if id == "" {
		d.SetId("")
	} else {
		config := meta.(configuration.Configuration)

		xiResource, err := config.Client.Operations.CloudCredsGet(api_operations.NewCloudCredsGetParams().WithID(id), config.Auth)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		setCloudCreds(d, xiResource.Payload)
	}

	return nil
}

func updateCloudCreds(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateCloudCreds")

	config := meta.(configuration.Configuration)

	tfResource := getCloudCreds(d)
	_, err := config.Client.Operations.CloudCredsCreate(api_operations.NewCloudCredsCreateParams().WithBody(tfResource), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setCloudCreds(d, tfResource)

	return nil
}

func deleteCloudCreds(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteCloudCreds")

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
