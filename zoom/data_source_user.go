package zoom

import (
   "strings"
   "fmt"
   "terraform-provider-zoom/client"
   "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
  return &schema.Resource{
    Read: dataSourceUserRead,
    Schema: map[string]*schema.Schema{
	  "first_name": &schema.Schema{
		  Type:     schema.TypeString,
		  Computed: true,
	  },
	  "last_name": &schema.Schema{
        	Type:     schema.TypeString,
        	Computed: true,
      	  },
	  "email": &schema.Schema{
        	Type:     schema.TypeString,
        	Computed: true,
      	  },
	  "status": &schema.Schema{
		Type:        schema.TypeString,
		Computed :   true,
	   },
	   "id": &schema.Schema{
        	Type:     schema.TypeString,
        	Required: true,
      	   },
      	   "license_type": &schema.Schema{
        	Type:     schema.TypeInt,
        	Computed: true,
           },
      	   "pmi": &schema.Schema{
		Type:        schema.TypeInt,
		Computed:    true,
	    },
      	    "role_name": &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
			},
            "department": &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
	     },
	     "job_title": &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
	      },
	      "location": &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
	      },
    },
  }
}

func dataSourceUserRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
  	userId := d.Get("id").(string)
  	user, err := apiClient.GetItem(userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Item with ID %s", userId)
		}
	}
  	d.SetId(user.Email)
	d.Set("email", user.Email)
	d.Set("first_name", user.FirstName)
	d.Set("last_name", user.LastName)
  	d.Set("license_type",user.Type)
  	d.Set("pmi",user.Pmi)
  	d.Set("status",user.Status)
  	d.Set("role_name",user.RoleName)
  	d.Set("department",user.Department)
	d.Set("job_title", user.JobTitle)
	d.Set("location", user.Location)
  	return nil
}
  
  
