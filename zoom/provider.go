package zoom

import (
	"terraform-provider-zoom/client"
	tkn "terraform-provider-zoom/token"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apisecret": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				DefaultFunc: schema.EnvDefaultFunc("API_SECRET", ""),
			},
			"apikey": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				DefaultFunc: schema.EnvDefaultFunc("API_KEY", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"zoom_user": resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"zoom_user": dataSourceUser(),
		},
		ConfigureFunc:  providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	token := tkn.GenerateToken(d.Get("apisecret").(string),d.Get("apikey").(string))
	return client.NewClient(token), nil
}
