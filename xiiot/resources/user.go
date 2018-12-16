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
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Optional: false,
				Computed: false,
				ForceNew: false,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Optional: false,
				Computed: false,
				ForceNew: false,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: false,
				ForceNew: false,
			},
		},
	}
}

func getUser(d *schema.ResourceData) *api_models.User {
	// name := d.Get("name").(string)
	// purpose := d.Get("purpose").(string)

	resource := api_models.User{
		// ID:      d.Id(),
		// Name:    &name,
		// Purpose: &purpose,
		// Values:  utils.Convert_set_to_string_array(d.Get("values").(*schema.Set)),
	}

	return &resource
}

func setUser(d *schema.ResourceData, resource *api_models.User) {
	// d.Set("xi_id", resource.ID)
	// d.Set("name", resource.Name)
	// d.Set("purpose", resource.Purpose)
	// d.Set("values", resource.Values)
	// d.SetId(resource.ID)
}

func createUser(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createUser")

	read := readUser(d, meta)

	if read != nil {
		config := meta.(configuration.Configuration)

		tfResource := getUser(d)
		_, err := config.Client.Operations.UserCreate(api_operations.NewUserCreateParams().WithBody(tfResource), config.Auth)

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		// TODO : should read back ?

		setUser(d, tfResource)
	}

	return nil
}

func readUser(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readUser")

	id := d.Id()

	if id == "" {
		id = d.Get("xi_id").(string)
	}

	if id == "" {
		d.SetId("")
	} else {
		config := meta.(configuration.Configuration)

		xiResource, err := config.Client.Operations.UserGet(api_operations.NewUserGetParams().WithID(id), config.Auth)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		setUser(d, xiResource.Payload)
	}

	return nil
}

func updateUser(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateUser")

	config := meta.(configuration.Configuration)

	tfResource := getUser(d)
	_, err := config.Client.Operations.UserCreate(api_operations.NewUserCreateParams().WithBody(tfResource), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setUser(d, tfResource)

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
