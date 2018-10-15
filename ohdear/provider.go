package ohdear

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OHDEAR_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ohdear_site": resourceOhdearSite(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{Token: d.Get("api_token").(string)}
	return config, nil
}
