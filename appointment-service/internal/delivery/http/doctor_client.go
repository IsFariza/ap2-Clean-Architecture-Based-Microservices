package delivery

import (
	"context"
	"net/http"
)

type doctorClient struct {
	baseURL string
	client  *http.Client
}

func NewDoctorClient(url string) *doctorClient {
	return &doctorClient{
		baseURL: url,
		client:  &http.Client{},
	}
}

func (dc *doctorClient) DoctorExists(ctx context.Context, doctorId string) (bool, error) {
	url := dc.baseURL + "/doctors/" + doctorId

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false, err
	}

	resp, err := dc.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}
