package v1

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var expectedQuotas = &ProjectQuotasResponse{
	NetworkDirectPublicIPs: []Quota{
		{Value: 5, Used: 3},
	},
}

func TestPublicNetAPIClient_GetQuotas(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		httpClient := &testHTTPClient{
			response: &http.Response{
				Body:       io.NopCloser(strings.NewReader(apiQuotasJSON)),
				StatusCode: http.StatusOK,
			},
		}
		cfg := &Config{
			AuthToken:  "testToken",
			URL:        "http://test.com",
			HTTPClient: httpClient,
		}

		client, _ := NewPublicNetAPIClient(cfg)

		quotas, err := client.GetQuotas(context.Background(), "test_project")
		require.NoError(t, err)

		require.Equal(t, expectedQuotas, quotas)

		require.Equal(t, "http://test.com/v1/projects/test_project/quotas", httpClient.request.URL.String())
		require.Equal(t, http.MethodGet, httpClient.request.Method)
		require.Equal(t, cfg.AuthToken, httpClient.request.Header.Get(xAuthHeader))
	})

	t.Run("APIError", func(t *testing.T) {
		httpClient := &testHTTPClient{
			response: &http.Response{
				Body:       io.NopCloser(strings.NewReader(apiErrorJSON)),
				StatusCode: http.StatusNotFound,
			},
		}
		cfg := &Config{
			AuthToken:  "testToken",
			URL:        "http://test.com",
			HTTPClient: httpClient,
		}

		client, _ := NewPublicNetAPIClient(cfg)

		_, err := client.GetQuotas(context.Background(), "test_project")
		var apiErr *APIErr
		require.ErrorAs(t, err, &apiErr)
		require.Equal(t, "Not found", apiErr.Msg)
		require.Equal(t, apiErrorJSON, apiErr.Raw())

		require.Equal(t, "http://test.com/v1/projects/test_project/quotas", httpClient.request.URL.String())
		require.Equal(t, http.MethodGet, httpClient.request.Method)
		require.Equal(t, cfg.AuthToken, httpClient.request.Header.Get(xAuthHeader))
	})
}
