package resources

import (
	"log"

	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
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

func getContainerRegistry(d *schema.ResourceData) *api_models.ContainerRegistry {
	// name := d.Get("name").(string)
	// purpose := d.Get("purpose").(string)

	resource := api_models.ContainerRegistry{
		// ID:      d.Id(),
		// Name:    &name,
		// Purpose: &purpose,
		// Values:  utils.Convert_set_to_string_array(d.Get("values").(*schema.Set)),
	}

	return &resource
}

func setContainerRegistry(d *schema.ResourceData, resource *api_models.ContainerRegistry) {
	// d.Set("xi_id", resource.ID)
	// d.Set("name", resource.Name)
	// d.Set("purpose", resource.Purpose)
	// d.Set("values", resource.Values)
	// d.SetId(resource.ID)
}

func createContainerRegistry(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createContainerRegistry")

	read := readContainerRegistry(d, meta)

	if read != nil {
		config := meta.(configuration.Configuration)

		tfResource := getContainerRegistry(d)
		_, err := config.Client.Operations.ContainerRegistryCreate(api_operations.NewContainerRegistryCreateParams().WithBody(tfResource), config.Auth)

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		// TODO : should read back ?

		setContainerRegistry(d, tfResource)
	}

	return nil
}

func readContainerRegistry(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readContainerRegistry")

	id := d.Id()

	if id == "" {
		id = d.Get("xi_id").(string)
	}

	if id == "" {
		d.SetId("")
	} else {
		config := meta.(configuration.Configuration)

		xiResource, err := config.Client.Operations.ContainerRegistryGet(api_operations.NewContainerRegistryGetParams().WithID(id), config.Auth)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		setContainerRegistry(d, xiResource.Payload)
	}

	return nil
}

func updateContainerRegistry(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateContainerRegistry")

	config := meta.(configuration.Configuration)

	tfResource := getContainerRegistry(d)
	_, err := config.Client.Operations.ContainerRegistryCreate(api_operations.NewContainerRegistryCreateParams().WithBody(tfResource), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setContainerRegistry(d, tfResource)

	return nil
}

func deleteContainerRegistry(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteContainerRegistry")

	id := d.Id()

	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.ContainerRegistryDelete(api_operations.NewContainerRegistryDeleteParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to delete resource : ", err)

		return err
	}

	d.SetId("")

	return nil
}
