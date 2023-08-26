package options

import (
	"flag"
	"os"
	"strconv"
	"strings"
	"time"

	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	setupLog = ctrl.Log.WithName("setup")
	o        Options
)

// Options are flags/env
type Options struct {
	GrafanaUser         string
	GrafanaPass         string
	GrafanaApiKey		string
	GrafanaTimeoutSeconds	string
}

// -- []string Value
type StringSliceValue []string

func newStringSliceValue(p *[]string, val []string) *StringSliceValue {
	*p = val
	return (*StringSliceValue)(p)
}

func (ss *StringSliceValue) Set(s string) error {
	if s != "" {
		*ss = StringSliceValue(strings.Split(s, ","))
	}
	return nil
}

func (ss *StringSliceValue) String() string { return strings.Join(([]string)(*ss), ",") }

// Parse returns the config struct
func Parse() *Options {
	flag.StringVar(&o.GrafanaUser, "grafana-user", lookupEnvOrString("GRAFANA_USER", ""), "Grafana user to connect to its API.")
	flag.StringVar(&o.GrafanaPass, "grafana-pass", lookupEnvOrString("GRAFANA_PASS", ""), "Grafana user password to connect to its API.")
	flag.StringVar(&o.GrafanaApiKey, "grafana-apikey", lookupEnvOrString("GRAFANA_API_KEY", ""), "Grafana api key to connect to its API.")
	flag.StringVar(&o.GrafanaTimeoutSeconds, "grafana-timeout-seconds", lookupEnvOrString("GRAFANA_TIMEOUT_SECONDS", ""), "Grafana timeout in seconds.")
	flag.Parse()

	return &o
}

func lookupEnvOrDuration(key string, defaultVal time.Duration) time.Duration {
	if val, ok := os.LookupEnv(key); ok {
		duration, err := time.ParseDuration(val)
		if err != nil {
			setupLog.Error(err, "lookupEnvOrDuration", "key", key, "value", val)
		}
		return duration
	}
	return defaultVal
}

func lookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func lookupEnvOrStringSlice(key string, defaultVal []string) []string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return strings.Split(val, ",")
	}
	return defaultVal
}

func lookupEnvOrBool(key string, defaultVal bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseBool(val)
		if err != nil {
			setupLog.Error(err, "LookupEnvOrBool", "key", key, "value", val)
		}
		return v
	}
	return defaultVal
}