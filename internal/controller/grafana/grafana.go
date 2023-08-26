package grafana

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/forselli-stratio/grafana-operator/internal/controller/options"
	gapi "github.com/grafana/grafana-api-golang-client"
)

type grafanaAdminCredentials struct {
	username string
	password string
	apikey   string
}

func NewGrafanaClient(grafanaUrl string) (*gapi.Client, error) {
	opts := options.Parse()

	credentials := &grafanaAdminCredentials{
		username: opts.GrafanaUser,
		password: opts.GrafanaPass,
		apikey: opts.GrafanaApiKey,
	}

	timeoutStr, _ := strconv.Atoi(opts.GrafanaTimeoutSeconds)
	var timeout time.Duration
	if timeoutStr != 0 {
		timeout = time.Duration(timeoutStr)
		if timeout < 0 {
			timeout = 0
		}
	} else {
		timeout = 10
	}

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
