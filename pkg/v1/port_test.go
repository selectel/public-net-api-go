package v1

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	testPortDescription = "test port"
	expectedPort        = &Port{
		ID:          "test_port",
		ProjectID:   "test_project",
		NetworkID:   "test_network",
		IPAddress:   "10.10.10.5",
		CreatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		Description: &testPortDescription,
	}
)

var expectedPortsList = []*Port{{
	ID:          "test_port",
	ProjectID:   "test_project",
	NetworkID:   "test_network",
	IPAddress:   "10.10.10.5",
	CreatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	UpdatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	Description: &testPortDescription,
}}

func TestPublicNetAPIClient_ListPorts(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		httpClient := &testHTTPClient{
			response: &http.Response{
				Body:       io.NopCloser(strings.NewReader(apiPortListJSON)),
				StatusCode: http.StatusOK,
			},
		}
		cfg := &Config{
			AuthToken:  "testToken",
			URL:        "http://test.com",
			HTTPClient: httpClient,
		}

		client, _ := NewPublicNetAPIClient(cfg)

		ports, err := client.ListPorts(context.Background(), &PortsQuery{ProjectID: "test_project"})
		require.NoError(t, err)

		require.Equal(t, expectedPortsList, ports)

		require.Equal(t, "http://test.com/v1/public_ports?project_id=test_project", httpClient.request.URL.String())
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

		_, err := client.ListPorts(context.Background(), nil)
		var apiErr *APIErr
		require.ErrorAs(t, err, &apiErr)
		require.Equal(t, "Not found", apiErr.Msg)
		require.Equal(t, apiErrorJSON, apiErr.Raw())

		require.Equal(t, "http://test.com/v1/public_ports", httpClient.request.URL.String())
		require.Equal(t, http.MethodGet, httpClient.request.Method)
		require.Equal(t, cfg.AuthToken, httpClient.request.Header.Get(xAuthHeader))
	})
}

func TestPublicNetAPIClient_GetPort(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		httpClient := &testHTTPClient{
			response: &http.Response{
				Body:       io.NopCloser(strings.NewReader(apiPortDetailJSON)),
				StatusCode: http.StatusOK,
			},
		}
		cfg := &Config{
			AuthToken:  "testToken",
			URL:        "http://test.com",
			HTTPClient: httpClient,
		}

		client, _ := NewPublicNetAPIClient(cfg)
		port, err := client.GetPort(context.Background(), "test_port")
		require.NoError(t, err)

		require.Equal(t, expectedPort, port)

		require.Equal(t, "http://test.com/v1/public_ports/test_port", httpClient.request.URL.String())
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

		_, err := client.GetPort(context.Background(), "test_port")
		var apiErr *APIErr
		require.ErrorAs(t, err, &apiErr)
		require.Equal(t, "Not found", apiErr.Msg)
		require.Equal(t, apiErrorJSON, apiErr.Raw())

		require.Equal(t, "http://test.com/v1/public_ports/test_port", httpClient.request.URL.String())
		require.Equal(t, http.MethodGet, httpClient.request.Method)
		require.Equal(t, cfg.AuthToken, httpClient.request.Header.Get(xAuthHeader))
	})
}

func TestPublicNetAPIClient_CreatePort(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		httpClient := &testHTTPClient{
			response: &http.Response{
				Body:       io.NopCloser(strings.NewReader(apiPortDetailJSON)),
				StatusCode: http.StatusCreated,
			},
		}
		cfg := &Config{
			AuthToken:  "testToken",
			URL:        "http://test.com",
			HTTPClient: httpClient,
		}

		client, _ := NewPublicNetAPIClient(cfg)

		dto := &PortCreateDTO{
			ProjectID: "test_project",
		}

		port, err := client.CreatePort(context.Background(), dto)
		require.NoError(t, err)

		require.Equal(t, expectedPort, port)

		require.Equal(t, "http://test.com/v1/public_ports", httpClient.request.URL.String())
		require.Equal(t, http.MethodPost, httpClient.request.Method)
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
		dto := &PortCreateDTO{
			ProjectID: "test_project",
		}

		_, err := client.CreatePort(context.Background(), dto)
		var apiErr *APIErr
		require.ErrorAs(t, err, &apiErr)
		require.Equal(t, "Not found", apiErr.Msg)
		require.Equal(t, apiErrorJSON, apiErr.Raw())

		require.Equal(t, "http://test.com/v1/public_ports", httpClient.request.URL.String())
		require.Equal(t, http.MethodPost, httpClient.request.Method)
		require.Equal(t, cfg.AuthToken, httpClient.request.Header.Get(xAuthHeader))
	})
}

func TestPublicNetAPIClient_UpdatePort(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		httpClient := &testHTTPClient{
			response: &http.Response{
				Body:       io.NopCloser(strings.NewReader(apiPortDetailJSON)),
				StatusCode: http.StatusOK,
			},
		}
		cfg := &Config{
			AuthToken:  "testToken",
			URL:        "http://test.com",
			HTTPClient: httpClient,
		}

		client, _ := NewPublicNetAPIClient(cfg)

		desc := "updated"
		port, err := client.UpdatePort(context.Background(), "test_port", &PortUpdateDTO{Description: &desc})
		require.NoError(t, err)

		require.Equal(t, expectedPort, port)

		require.Equal(t, "http://test.com/v1/public_ports/test_port", httpClient.request.URL.String())
		require.Equal(t, http.MethodPatch, httpClient.request.Method)
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

		_, err := client.UpdatePort(context.Background(), "test_port", &PortUpdateDTO{})
		var apiErr *APIErr
		require.ErrorAs(t, err, &apiErr)
		require.Equal(t, "Not found", apiErr.Msg)
		require.Equal(t, apiErrorJSON, apiErr.Raw())

		require.Equal(t, "http://test.com/v1/public_ports/test_port", httpClient.request.URL.String())
		require.Equal(t, http.MethodPatch, httpClient.request.Method)
		require.Equal(t, cfg.AuthToken, httpClient.request.Header.Get(xAuthHeader))
	})
}

func TestPublicNetAPIClient_DeletePort(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		httpClient := &testHTTPClient{
			response: &http.Response{
				Body:       io.NopCloser(strings.NewReader("")),
				StatusCode: http.StatusNoContent,
			},
		}
		cfg := &Config{
			AuthToken:  "testToken",
			URL:        "http://test.com",
			HTTPClient: httpClient,
		}

		client, _ := NewPublicNetAPIClient(cfg)
		err := client.DeletePort(context.Background(), "test_port", nil)
		require.NoError(t, err)

		require.Equal(t, "http://test.com/v1/public_ports/test_port", httpClient.request.URL.String())
		require.Equal(t, http.MethodDelete, httpClient.request.Method)
		require.Equal(t, cfg.AuthToken, httpClient.request.Header.Get(xAuthHeader))
	})

	t.Run("WithForce", func(t *testing.T) {
		httpClient := &testHTTPClient{
			response: &http.Response{
				Body:       io.NopCloser(strings.NewReader("")),
				StatusCode: http.StatusNoContent,
			},
		}
		cfg := &Config{
			AuthToken:  "testToken",
			URL:        "http://test.com",
			HTTPClient: httpClient,
		}

		client, _ := NewPublicNetAPIClient(cfg)
		err := client.DeletePort(context.Background(), "test_port", &PortDeleteQuery{Force: true})
		require.NoError(t, err)

		require.Equal(t, "http://test.com/v1/public_ports/test_port?force=true", httpClient.request.URL.String())
		require.Equal(t, http.MethodDelete, httpClient.request.Method)
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

		err := client.DeletePort(context.Background(), "test_port", nil)
		var apiErr *APIErr
		require.ErrorAs(t, err, &apiErr)
		require.Equal(t, "Not found", apiErr.Msg)
		require.Equal(t, apiErrorJSON, apiErr.Raw())

		require.Equal(t, "http://test.com/v1/public_ports/test_port", httpClient.request.URL.String())
		require.Equal(t, http.MethodDelete, httpClient.request.Method)
		require.Equal(t, cfg.AuthToken, httpClient.request.Header.Get(xAuthHeader))
	})
}
