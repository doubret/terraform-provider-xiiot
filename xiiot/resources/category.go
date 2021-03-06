package resources

import (
	api_operations "github.com/doubret/terraform-provider-xiiot/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/models"
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
}

func createCategory(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	result, err := config.Client.Operations.CategoryCreate(api_operations.NewCategoryCreateParams().WithBody(getCategory(d)), config.Auth)

	if err != nil {
		return err
	}

	d.SetId(*result.Payload.ID)

	return readCategory(d, meta)
}

func readCategory(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	resource, err := config.Client.Operations.CategoryGet(api_operations.NewCategoryGetParams().WithID(d.Id()), config.Auth)

	if err != nil {
		if httperr, ok := err.(*api_operations.CategoryGetDefault); ok {
			if httperr.Code() == 404 {
				d.SetId("")

				return nil
			}
		}

		return err
	}

	setCategory(d, resource.Payload)

	return nil
}

func updateCategory(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.CategoryUpdateV2(api_operations.NewCategoryUpdateV2Params().WithID(d.Id()).WithBody(getCategory(d)), config.Auth)

	if err != nil {
		return err
	}

	return readCategory(d, meta)
}

func deleteCategory(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.CategoryDelete(api_operations.NewCategoryDeleteParams().WithID(d.Id()), config.Auth)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
