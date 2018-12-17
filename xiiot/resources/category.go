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
		Create:        createCategory,
		Read:          readCategory,
		Update:        updateCategory,
		Delete:        deleteCategory,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"purpose": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"values": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
		},
	}
}

func getCategory(d *schema.ResourceData) *api_models.Category {
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

func setCategory(d *schema.ResourceData, resource *api_models.Category) {
	d.Set("name", resource.Name)
	d.Set("purpose", resource.Purpose)
	d.Set("values", resource.Values)
	d.SetId(resource.ID)
}

func createCategory(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createCategory")

	config := meta.(configuration.Configuration)

	model := getCategory(d)
	_, err := config.Client.Operations.CategoryCreate(api_operations.NewCategoryCreateParams().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setCategory(d, model)

	return nil
}

func readCategory(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readCategory")

	id := d.Id()

	config := meta.(configuration.Configuration)

	model, err := config.Client.Operations.CategoryGet(api_operations.NewCategoryGetParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to read resource : ", err)

		return err
	}

	setCategory(d, model.Payload)

	return nil
}

func updateCategory(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateCategory")

	config := meta.(configuration.Configuration)

	model := getCategory(d)
	_, err := config.Client.Operations.CategoryUpdateV2(api_operations.NewCategoryUpdateV2Params().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to update resource : ", err)

		return err
	}

	setCategory(d, model)

	return nil
}

func deleteCategory(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteCategory")

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
