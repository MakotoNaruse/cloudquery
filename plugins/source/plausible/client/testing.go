package client

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func MockTestHelper(t *testing.T, table *schema.Table, createServices func(*mux.Router) error) {
	t.Helper()

	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	table.IgnoreInTests = false
	version := "vDev"
	router := mux.NewRouter()
	h := httptest.NewServer(router)
	defer h.Close()

	configureTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		err := createServices(router)
		if err != nil {
			return nil, err
		}
		var pluginSpec Spec
		if err = spec.UnmarshalSpec(&pluginSpec); err != nil {
			return nil, err
		}
		pluginSpec.SetDefaults()
		cqClient := Client{
			logger:     logger,
			Client:     h.Client(),
			PluginSpec: pluginSpec,
		}
		return &cqClient, nil
	}

	p := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{table},
		configureTestExecutionClient,
	)
	p.SetLogger(logger)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
		Spec: Spec{
			SiteId:  "testSiteId",
			ApiKey:  "testApiKey",
			BaseURL: h.URL,
		},
	})
}
