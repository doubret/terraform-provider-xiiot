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
				Required: false,
				Optional: true,
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
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Optional: false,
				Computed: false,
				ForceNew: false,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: false,
				ForceNew: false,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: false,
				ForceNew: false,
			},
			"pwd": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: false,
				ForceNew: false,
			},
			"cloud_creds_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: false,
				ForceNew: false,
			},
			"credentials": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: false,
				ForceNew: false,
			},
		},
	}
}

func getDockerProfile(d *schema.ResourceData) *api_models.DockerProfile {
	// name := d.Get("name").(string)
	// purpose := d.Get("purpose").(string)

	resource := api_models.DockerProfile{
		// ID:      d.Id(),
		// Name:    &name,
		// Purpose: &purpose,
		// Values:  utils.Convert_set_to_string_array(d.Get("values").(*schema.Set)),
	}

	return &resource
}

func setDockerProfile(d *schema.ResourceData, resource *api_models.DockerProfile) {
	// d.Set("xi_id", resource.ID)
	// d.Set("name", resource.Name)
	// d.Set("purpose", resource.Purpose)
	// d.Set("values", resource.Values)
	// d.SetId(resource.ID)
}

func createDockerProfile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createDockerProfile")

	read := readDockerProfile(d, meta)

	if read != nil {
		config := meta.(configuration.Configuration)

		tfResource := getDockerProfile(d)
		_, err := config.Client.Operations.DockerProfileCreate(api_operations.NewDockerProfileCreateParams().WithBody(tfResource), config.Auth)

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		// TODO : should read back ?

		setDockerProfile(d, tfResource)
	}

	return nil
}

func readDockerProfile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readDockerProfile")

	id := d.Id()

	if id == "" {
		id = d.Get("xi_id").(string)
	}

	if id == "" {
		d.SetId("")
	} else {
		config := meta.(configuration.Configuration)

		xiResource, err := config.Client.Operations.DockerProfileGet(api_operations.NewDockerProfileGetParams().WithID(id), config.Auth)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		setDockerProfile(d, xiResource.Payload)
	}

	return nil
}

func updateDockerProfile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateDockerProfile")

	config := meta.(configuration.Configuration)

	tfResource := getDockerProfile(d)
	_, err := config.Client.Operations.DockerProfileCreate(api_operations.NewDockerProfileCreateParams().WithBody(tfResource), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setDockerProfile(d, tfResource)

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
