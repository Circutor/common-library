// Copyright (c) 2021 Circutor S.A. All rights reserved.

package request

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/circutor/common-library/internal/request/config"
	"github.com/jbrodriguez/mlog"
)

const (
	contentType     = "Content-Type"
	contentTypeJSON = "application/json"
	authorization   = "x-authorization"
	// Constants error.
	errCreateNewRequest = "error in create request"
	errSendRequest      = "error in send request"
	errReadBody         = "error in read response body"
)

// CreateNewRequest method create request.
func CreateNewRequest(method, url, token string, body io.Reader, query map[string]interface{}) ([]byte, int, error) {
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("%s : %w", errCreateNewRequest, err)
	}

	if body != nil {
		req.Header.Set(contentType, contentTypeJSON)
	}

	if token != "" {
		req.Header.Set(authorization, token)
	}

	if query != nil {
		addQueryParameters(req, query)
	}

	resp, err := makeRequest(req)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()

	respBody, err := getBody(resp)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("%w", err)
	}

	return respBody, resp.StatusCode, nil
}

// addQueryParameters method aggregate queries in to the request.
func addQueryParameters(req *http.Request, queryParameters map[string]interface{}) {
	query := req.URL.Query()

	for key, element := range queryParameters {
		query.Add(key, fmt.Sprintf("%v", element))
	}

	req.URL.RawQuery = query.Encode()
}

// Method to make the request and return the response.
func makeRequest(req *http.Request) (*http.Response, error) {
	client := config.CreateHTTPClient()

	start := time.Now()

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", errSendRequest, err)
	}

	timestamp := time.Now()
	latency := timestamp.Sub(start)

	mlog.Info("%s request to %s take %13v with result %d",
		req.Method,
		req.URL,
		latency,
		resp.StatusCode)

	return resp, nil
}

// getBody get content body.
func getBody(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", errReadBody, err)
	}

	return body, nil
}
