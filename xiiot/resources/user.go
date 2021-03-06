package resources

import (
	api_operations "github.com/doubret/terraform-provider-xiiot/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/doubret/terraform-provider-xiiot/xiiot/utils"
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
				Type:      schema.TypeString,
				Required:  true,
				StateFunc: utils.Hash,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func getUser(d *schema.ResourceData) *api_models.User {
	name := d.Get("name").(string)
	email := d.Get("email").(string)
	password := d.Get("password").(string)
	tenant_id := d.Get("tenant_id").(string)

	resource := api_models.User{
		ID:       d.Id(),
		Name:     &name,
		Email:    &email,
		Password: &password,
		TenantID: &tenant_id,
		Role:     d.Get("role").(string),
	}

	return &resource
}

func setUser(d *schema.ResourceData, resource *api_models.User) {
	d.Set("name", resource.Name)
	d.Set("email", resource.Email)
	d.Set("tenant_id", resource.TenantID)
	d.Set("role", resource.Role)
}

func createUser(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	result, err := config.Client.Operations.UserCreate(api_operations.NewUserCreateParams().WithBody(getUser(d)), config.Auth)

	if err != nil {
		return err
	}

	d.SetId(*result.Payload.ID)

	return readUser(d, meta)
}

func readUser(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	resource, err := config.Client.Operations.UserGet(api_operations.NewUserGetParams().WithID(d.Id()), config.Auth)

	if err != nil {
		if httperr, ok := err.(*api_operations.UserGetDefault); ok {
			if httperr.Code() == 404 {
				d.SetId("")

				return nil
			}
		}

		return err
	}

	setUser(d, resource.Payload)

	return nil
}

func updateUser(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.UserUpdateV2(api_operations.NewUserUpdateV2Params().WithID(d.Id()).WithBody(getUser(d)), config.Auth)

	if err != nil {
		return err
	}

	return readUser(d, meta)
}

func deleteUser(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.UserDelete(api_operations.NewUserDeleteParams().WithID(d.Id()), config.Auth)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
