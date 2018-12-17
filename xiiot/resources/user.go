package resources

import (
	"log"

	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTUser() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createUser,
		Read:          readUser,
		Update:        updateUser,
		Delete:        deleteUser,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
			},
		},
	}
}

func getUser(d *schema.ResourceData) *api_models.User {
	name := d.Get("name").(string)
	email := d.Get("email").(string)
	password := d.Get("password").(string)

	resource := api_models.User{
		ID:       d.Id(),
		Name:     &name,
		Email:    &email,
		Password: &password,
		Role:     d.Get("role").(string),
	}

	return &resource
}

func setUser(d *schema.ResourceData, resource *api_models.User) {
	d.Set("name", resource.Name)
	d.Set("email", resource.Email)
	d.Set("password", resource.Password)
	d.Set("role", resource.Role)
	d.SetId(resource.ID)
}

func createUser(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createUser")

	config := meta.(configuration.Configuration)

	model := getUser(d)
	_, err := config.Client.Operations.UserCreate(api_operations.NewUserCreateParams().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setUser(d, model)

	return nil
}

func readUser(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readUser")

	id := d.Id()

	config := meta.(configuration.Configuration)

	model, err := config.Client.Operations.UserGet(api_operations.NewUserGetParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	}

	setUser(d, model.Payload)

	return nil
}

func updateUser(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateUser")

	config := meta.(configuration.Configuration)

	model := getUser(d)
	_, err := config.Client.Operations.UserUpdateV2(api_operations.NewUserUpdateV2Params().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to update resource : ", err)

		return err
	}

	setUser(d, model)

	return nil
}

func deleteUser(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteUser")

	id := d.Id()

	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.UserDelete(api_operations.NewUserDeleteParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to delete resource : ", err)

		return err
	}

	d.SetId("")

	return nil
}
