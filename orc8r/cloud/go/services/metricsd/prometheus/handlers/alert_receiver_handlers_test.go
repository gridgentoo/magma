package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"magma/orc8r/cloud/go/services/metricsd/prometheus/configmanager/alertmanager/config"

	"github.com/labstack/echo"
	amconfig "github.com/prometheus/alertmanager/config"
	"github.com/stretchr/testify/assert"
)

const (
	webhookURL = "http://test.com"
	slackURL   = "http://slack.com"
)

var (
	testWebhookConfigBody = fmt.Sprintf(`{
      "name": "test",
      "webhook_configs": [
      {
		  "send_resolved": true,
		  "url": "%s"
      }
      ],
      "slack_configs": [
      {
         "api_url": "%s"
      }
      ]
    }`, webhookURL, slackURL)

	testWebhookURL, _ = url.Parse(webhookURL)
	testWebhookConfig = config.WebhookConfig{
		NotifierConfig: amconfig.NotifierConfig{
			VSendResolved: true,
		},
		URL: &amconfig.URL{
			URL: testWebhookURL,
		},
	}
	testSlackConfig = config.SlackConfig{
		APIURL: slackURL,
	}
)

func TestBuildReceiverFromContext(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(testWebhookConfigBody))
	rec := httptest.NewRecorder()

	c := echo.New().NewContext(req, rec)

	receiver, err := buildReceiverFromContext(c)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(receiver.SlackConfigs))
	assert.Equal(t, 1, len(receiver.WebhookConfigs))
	assert.Equal(t, testWebhookConfig, *receiver.WebhookConfigs[0])
	assert.Equal(t, testSlackConfig, *receiver.SlackConfigs[0])
}
