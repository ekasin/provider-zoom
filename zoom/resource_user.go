package zoom

import (
	"terraform-provider-zoom/client"
	"terraform-provider-zoom/server"
	"strings"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)




func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Exists: resourceExistsUser,
		Schema: map[string]*schema.Schema{

			"email": &schema.Schema{
				Type:        schema.TypeString,
				Description: "emailId of new user",
				Required:    true,
			},
			"first_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "first name of new user",
				Required:    true,
			},
			"last_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "last name of new user",
				Required:    true,
			},
		},
	}

}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	// Warning or errors can be collected in a slice type
	apiClient := m.(*client.Client)

	user := server.Item{
		EmailId:   d.Get("email").(string),
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
	}

	err := apiClient.NewItem(&user)

	if err != nil {
		return err
	}
	d.SetId(user.EmailId)
	return nil
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	// Warning or errors can be collected in a slice type

	return nil
}


func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	user := server.Item{
		EmailId:   d.Get("email").(string),
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
	}

	err := apiClient.UpdateItem(&user)
	if err != nil {
		return err
	}
	return nil
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	userId := d.Id()

	err := apiClient.DeleteItem(userId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}



func resourceExistsUser(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.GetItem(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}