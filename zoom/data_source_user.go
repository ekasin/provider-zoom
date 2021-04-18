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
	  "id": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      /*
      "type": &schema.Schema{
        Type:     schema.TypeInt,
        Computed: true,
      },
      "role_name": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "pmi": &schema.Schema{
        Type:     schema.TypeInt,
        Computed: true,
      },
      "use_pmi": &schema.Schema{
        Type:     schema.TypeBool,
        Computed: true,
      },
      "personal_meeting_url": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "timezone": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "verified": &schema.Schema{
        Type:     schema.TypeInt,
        Computed: true,
      },
      "dept": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "created_at": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "last_login_time": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "pic_url": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "host_key": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "cms_user_id": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "jid": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "account_id": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "language": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "phone_country": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "phone_number": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "status": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "job_title": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "location": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "role_id": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },*/
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

	
	d.SetId(user.EmailId)
	d.Set("email", user.EmailId)
	d.Set("first_name", user.FirstName)
	d.Set("last_name", user.LastName)
  /*
  d.Set("type", user.Type)
  d.Set("role_name", user.RoleName)
  d.Set("pmi", user.Pmi)
  d.Set("use_pmi", user.UsePmi)
  d.Set("personal_meeting_url", user.PersonalMeetingUrl)
  d.Set("timezone", user.TimeZone)
  d.Set("verified", user.Verified)
  d.Set("dept", user.Dept)
  d.Set("created_at", user.CreatedAt)
  d.Set("last_login_time", user.LastLoginTime)
  d.Set("pic_url", user.PicUrl)
  d.Set("host_key", user.HostKey)
  d.Set("cms_user_id", user.CmsUserId)
  d.Set("jid", user.Jid)
  d.Set("account_id", user.AccountId)
  d.Set("language", user.Language)
  d.Set("phone_country", user.PhoneCountry)
  d.Set("phone_number", user.PhoneNumber)
  d.Set("status", user.Status)
  d.Set("job_title", user.JobTitle)
  d.Set("location", user.Location)
  d.Set("role_id", user.RoleId)
	*/
  
	return nil
  }
  
  