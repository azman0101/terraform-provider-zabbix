package provider

import (
	"github.com/claranet/go-zabbix-api"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider define the provider and his resources
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ZABBIX_USER", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ZABBIX_PASSWORD", nil),
			},
			"server_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ZABBIX_SERVER_URL", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"zabbix_host":       resourceZabbixHost(),
			"zabbix_host_group": resourceZabbixHostGroup(),
			"zabbix_item":       resourceZabbixItem(),
			"zabbix_trigger":    resourceZabbixTrigger(),
			"zabbix_template":   resourceZabbixTemplate(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	api := zabbix.NewAPI(d.Get("server_url").(string))

	if _, err := api.Login(d.Get("user").(string), d.Get("password").(string)); err != nil {
		return nil, err
	}

	return api, nil
}
