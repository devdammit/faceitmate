package faceit

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (fc *API) request(method string, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, fc.hostURI+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", fc.apiKey))

	return req, nil
}

func (fc *API) doRequest(req *http.Request) (*http.Response, error) {
	resp, err := fc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		faceitReqErr := &ResponseError{
			StatusCode: resp.StatusCode,
		}

		switch {
		case faceitReqErr.NotAuthorized():
			faceitReqErr.Err = errors.New("not authorized")
		case faceitReqErr.NotFound():
			faceitReqErr.Err = errors.New("not found")
		case faceitReqErr.Temporary():
			faceitReqErr.Err = errors.New("faceit not available")
		default:
			faceitReqErr.Err = errors.New("unavailable")
		}

		return nil, faceitReqErr
	}

	return resp, nil
}
