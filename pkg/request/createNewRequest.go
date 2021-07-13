// Copyright (c) 2021 Circutor S.A. All rights reserved.

package request

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func (r *Request) CreateNewRequest(method, url, token string, body io.Reader,
	query map[string]interface{}) ([]byte, int, error) {
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
