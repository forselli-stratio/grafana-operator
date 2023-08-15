package grafana

import (
	"net/url"
	"net/http"
	"time"
	gapi "github.com/grafana/grafana-api-golang-client"
)

type grafanaAdminCredentials struct {
	username string
	password string
	apikey   string
}

func getAdminCredentials() (*grafanaAdminCredentials, error) {
	//TODO make logic
	credentials := &grafanaAdminCredentials{}
	credentials.username = "admin"
	credentials.password = "admin"
	return credentials, nil
}

func NewGrafanaClient(grafanaUrl string) (*gapi.Client, error) {
	credentials,err := getAdminCredentials()
	if err != nil {
		return nil, err
	}

	var timeout time.Duration = 10

	clientConfig := gapi.Config{
		HTTPHeaders: nil,
		Client: &http.Client{
			Timeout:   time.Second * timeout,
		},
		// TODO populate me
		OrgID: 0,
		// TODO populate me
		NumRetries: 0,
	}

	if credentials.apikey != "" {
		clientConfig.APIKey = credentials.apikey
	}

	if credentials.username != "" && credentials.password != "" {
		clientConfig.BasicAuth = url.UserPassword(credentials.username, credentials.password)
	}

	grafanaClient, err := gapi.New(grafanaUrl, clientConfig)
	if err != nil {
		return nil, err
	}

	return grafanaClient, nil
}
