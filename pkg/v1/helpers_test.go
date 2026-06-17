package v1

import "net/http"

const (
	apiErrorJSON      = `{"error":{"code":404,"message":"Not found"}}`
	apiPortListJSON   = `{"ports":[{"id":"test_port","project_id":"test_project","network_id":"test_network","ip_address":"10.10.10.5","created_at":"2024-01-01T00:00:00Z","updated_at":"2024-01-01T00:00:00Z","description":"test port"}]}`
	apiPortDetailJSON = `{"port":{"id":"test_port","project_id":"test_project","network_id":"test_network","ip_address":"10.10.10.5","created_at":"2024-01-01T00:00:00Z","updated_at":"2024-01-01T00:00:00Z","description":"test port"}}`
	apiQuotasJSON     = `{"network_direct_public_ips":[{"value":5,"used":3}]}`
)

type testHTTPClient struct {
	request  *http.Request
	response *http.Response
	err      error
}

func (client *testHTTPClient) Do(req *http.Request) (*http.Response, error) {
	client.request = req

	return client.response, client.err
}
