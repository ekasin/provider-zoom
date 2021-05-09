package zoom

import (
	"terraform-provider-zoom/client"
	"strings"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"fmt"
	"regexp"
	"log"
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	nameRegex := regexp.MustCompile("^[A-Za-z]\\w{5,29}$")

	if !(nameRegex.MatchString(k)) {
		errs = append(errs, fmt.Errorf("Expected name is not valid .Got %s", value))
		return warns, errs
	}
	return
}

func validateEmail(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value := v.(string)

	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !(emailRegex.MatchString(value)) {
		errs = append(errs, fmt.Errorf("Expected EmailId is not valid  %s", k))
		return warns, errs
	}
	return
}

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{

			"email": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validateEmail,
			},
			"first_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validateName,
			},
			"last_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validateName,
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Optional :   true,
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed :   true,
			},
			"type": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
			},
			"role_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			},
			"pmi": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional :   true,
			},
			"use_pmi": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Optional :   true,
			},
			  "personal_meeting_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			  },
			  "timezone": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			  },
			  "verified": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional :   true,
			  },
			  "dept": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			  },
			  "host_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			  },
			  "cms_user_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			  },
			  "jid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			  },
			  "account_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			  },
			  "language": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			  },
			  "phone_country": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			  },
			  "phone_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			  },
			  "job_title": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			  },
			  "location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			  },
			  "role_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional :   true,
			  },
		},
	}
}

func resourceUserCreate(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	user := client.User{
		Email:   d.Get("email").(string),
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
		Type:      d.Get("type").(int),
	}
	err := apiClient.NewItem(&user)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return diag.FromErr(err)
	}
	d.SetId(user.Email)
	resourceUserRead(ctx,d,m)
	return diags
}

func resourceUserRead(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	userId := d.Id()
	user, err := apiClient.GetItem(userId)
	if err != nil {
		log.Println("[ERROR]: ",err)
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return diag.FromErr(err)
		}
	}
	if len(user.Email) > 0{
		d.SetId(user.Email)
		d.Set("email", user.Email)
		d.Set("id", user.Id)
		d.Set("first_name", user.FirstName)
		d.Set("last_name", user.LastName)
		d.Set("type", user.Type)
		d.Set("role_name", user.RoleName)
		d.Set("pmi", user.Pmi)
		d.Set("use_pmi", user.UsePmi)
		d.Set("timezone", user.TimeZone)
		d.Set("verified", user.Verified)
		d.Set("dept", user.Dept)
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
	}
	return diags
}

func resourceUserUpdate(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var _ diag.Diagnostics
	apiClient := m.(*client.Client)
	var diags diag.Diagnostics
	if d.HasChange("email") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "User not allowed to change email",
			Detail:   "User not allowed to change email",
		})

		return diags
	}
	user := client.User{
		Email:   d.Get("email").(string),
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
	}
	status := d.Get("status").(string)
	errDeac := apiClient.DeactivateUser(user.Email, status)
	log.Println(errDeac)
	err := apiClient.UpdateItem(&user)
	if err != nil {
		log.Printf("[Error] Error updating user :%s", err)
		return diag.FromErr(err)
	}
	return resourceUserRead(ctx,d,m)
}

func resourceUserDelete(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	userId := d.Id()
	err := apiClient.DeleteItem(userId)
	if err != nil {
		log.Printf("[Error] Error deleting user :%s", err)
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

