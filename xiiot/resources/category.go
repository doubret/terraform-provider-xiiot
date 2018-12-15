package resources

import (
	"log"

	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/doubret/terraform-provider-xiiot/xiiot/utils"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTCategory() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_category,
		Read:          read_category,
		Update:        update_category,
		Delete:        delete_category,
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
			"purpose": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Optional: false,
				Computed: false,
				ForceNew: false,
			},
			"values": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
				Optional: false,
				Computed: false,
				ForceNew: false,
			},
		},
	}
}

func get_category(d *schema.ResourceData) *api_models.Category {
	name := d.Get("name").(string)
	purpose := d.Get("purpose").(string)

	resource := api_models.Category{
		ID:      d.Id(),
		Name:    &name,
		Purpose: &purpose,
		Values:  utils.Convert_set_to_string_array(d.Get("values").(*schema.Set)),
	}

	return &resource
}

func set_category(d *schema.ResourceData, resource *api_models.Category) {
	d.Set("xi_id", resource.ID)
	d.Set("name", resource.Name)
	d.Set("purpose", resource.Purpose)
	d.Set("values", resource.Values)
	d.SetId(resource.ID)
}

func create_category(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In create_category")

	read := read_category(d, meta)

	if read != nil {
		config := meta.(configuration.Configuration)

		tfResource := get_category(d)
		_, err := config.Client.Operations.CategoryCreate(api_operations.NewCategoryCreateParams().WithBody(tfResource), config.Auth)

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		// TODO : should read back

		set_category(d, tfResource)
	}

	return nil
}

func read_category(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In read_category")

	id := d.Id()

	if id == "" {
		id = d.Get("xi_id").(string)
	}

	if id == "" {
		d.SetId("")
	} else {
		config := meta.(configuration.Configuration)

		xiResource, err := config.Client.Operations.CategoryGet(api_operations.NewCategoryGetParams().WithID(id), config.Auth)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_category(d, xiResource.Payload)
	}

	return nil
}

func update_category(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In update_category")

	config := meta.(configuration.Configuration)

	tfResource := get_category(d)
	_, err := config.Client.Operations.CategoryCreate(api_operations.NewCategoryCreateParams().WithBody(tfResource), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	set_category(d, tfResource)

	return nil
}

func delete_category(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In delete_category")

	id := d.Id()

	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.CategoryDelete(api_operations.NewCategoryDeleteParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to delete resource : ", err)

		return err
	}

	d.SetId("")

	return nil
}
