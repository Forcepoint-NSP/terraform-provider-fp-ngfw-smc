// Copyright 2026 Forcepoint LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package request provides HTTP request utilities for the SMC provider.
package smc

import (
	"encoding/json"
	"fmt"
	"net/http"
	neturl "net/url"
	"time"
)

const (
	MAX_REQUEST_TIMEOUT_SEC = 300
)

// SearchResponse represents the response structure for search operations
type SearchResponse struct {
	Result []SearchResult `json:"result"`
}

// SearchResult represents an individual search result item
type SearchResult struct {
	Href string `json:"href"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// GenericCRUDConfig holds configuration for generic CRUD operations
type GenericCRUDConfig struct {
	ResourceType string // e.g., "host", "network", etc.
	Client       *SmcClient
}

type IResourceReader interface {
	ReadResourceByHref(href string) (*ResponseData, error)
}

func getBody(resource any) ([]byte, error) {
	if resourceStr, ok := resource.(string); ok {
		return []byte(resourceStr), nil
	} else if resourceBytes, ok := resource.([]byte); ok {
		return resourceBytes, nil
	} else {
		return json.Marshal(resource)
	}
}

// CreateResource creates a new resource in the SMC system using
// generic CRUD operations
func CreateResource(config *GenericCRUDConfig, resource any) (*ResponseData, error) {
	return _createResourceInternal(config, resource, "", nil)
}

// CreateSubResource creates a new resource in the SMC system using
// generic CRUD operations
func CreateSubResource(config *GenericCRUDConfig, resource any, url string) (*ResponseData, error) {
	return _createResourceInternal(config, resource, url, nil)
}

// CreateSubResourceWithQueryParams creates a new sub-resource in the
// SMC system with query parameters
func CreateSubResourceWithQueryParams(config *GenericCRUDConfig, resource any,
	url string, queryParams map[string]string) (*ResponseData, error) {
	return _createResourceInternal(config, resource, url, queryParams)
}

func _createResourceInternal(config *GenericCRUDConfig, resource any,
	url string, queryParams map[string]string) (*ResponseData, error) {

	client := config.Client

	reqBody, err := getBody(resource)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource: %w", err)
	}

	if url == "" {
		url = fmt.Sprintf("%s/%s/elements/%s", client.BaseUrl, client.APIVersion, config.ResourceType)
	}

	var finalURL string

	if queryParams == nil || len(queryParams) == 0 {
		finalURL = url
	} else {
		parsedURL, err := neturl.Parse(url)
		if err != nil {
			return nil, fmt.Errorf("failed to parse URL: %w", err)
		}

		query := parsedURL.Query()
		for key, value := range queryParams {
			query.Set(key, value)
		}
		parsedURL.RawQuery = query.Encode()
		finalURL = parsedURL.String()
	}

	headers := client.GetJSONHeaders()

	resp, err := client.DoRequest(Options{
		Method:  "POST",
		URL:     finalURL,
		Headers: headers,
		Body:    reqBody,
		Timeout: MAX_REQUEST_TIMEOUT_SEC * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	if resp.StatusCode != http.StatusCreated &&
		resp.StatusCode != http.StatusOK &&
		resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("create %s failed with status %d: %s", config.ResourceType, resp.StatusCode, string(resp.Body))
	}

	return resp, nil
}

// ReadResourceByName retrieves a resource by name from the SMC system
// using generic CRUD operations
func ReadResourceByName(config *GenericCRUDConfig, name string) (*ResponseData, error) {
	client := config.Client
	encodedName := neturl.QueryEscape(name)
	searchURL := fmt.Sprintf("%s/%s/elements/%s?filter=%s", client.BaseUrl, client.APIVersion, config.ResourceType, encodedName)
	headers := client.GetJSONHeaders()
	searchResp, err := client.DoRequest(Options{
		Method:  "GET",
		URL:     searchURL,
		Headers: headers,
		Timeout: MAX_REQUEST_TIMEOUT_SEC * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("search request failed: %w", err)
	}

	if searchResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search failed with status %d: %s", searchResp.StatusCode, string(searchResp.Body))
	}

	var searchResponse SearchResponse
	if err := json.Unmarshal(searchResp.Body, &searchResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal search response: %w", err)
	}

	if len(searchResponse.Result) == 0 {
		return nil, fmt.Errorf("no %s found with name: %s", config.ResourceType, name)
	}

	href := searchResponse.Result[0].Href
	resp, err := client.DoRequest(Options{
		Method:  "GET",
		URL:     href,
		Headers: client.GetAuthHeaders(),
		Timeout: MAX_REQUEST_TIMEOUT_SEC * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("read failed: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("read %s failed with status %d: %s", config.ResourceType, resp.StatusCode, string(resp.Body))
	}

	// Add href and etag to response for caller convenience
	resp.Location = href
	return resp, nil
}

// ReadResourceByHref retrieves a resource by href from the SMC system
// if the object does not exist, returns a StatusError with
// 404 status.
func (config *GenericCRUDConfig) ReadResourceByHref(href string) (*ResponseData, error) {
	if href == "" {
		return nil, fmt.Errorf("href cannot be empty")
	}
	client := config.Client
	resp, err := client.DoRequest(Options{
		Method:  "GET",
		URL:     href,
		Headers: client.GetAuthHeaders(),
		Timeout: MAX_REQUEST_TIMEOUT_SEC * time.Second,
	})
	if err != nil {
		// propagate error if any
		return nil, fmt.Errorf("read resource failed: %w", err)
	}

	if resp.StatusCode == http.StatusOK {
		return resp, nil
	}

	// turn any other status code into a StatusError
	return nil, NewStatusError(resp.StatusCode, fmt.Sprintf(
		"read %s failed with status %d: %s",
		config.ResourceType, resp.StatusCode, string(resp.Body)))
}

// UpdateResource updates an existing resource in the SMC system using
// generic CRUD operations
func UpdateResource(config *GenericCRUDConfig, resource any,
	href string, etag string) (*ResponseData, error) {

	client := config.Client

	if etag == "" {
		currentResp, err := config.ReadResourceByHref(href)
		if err != nil {
			return nil, fmt.Errorf("failed to read current %s for update. href='%s': %w",
				config.ResourceType, href, err)
		}
		if currentResp == nil {
			return nil, fmt.Errorf("%s not found for update: %s", config.ResourceType, href)
		}
		etag = currentResp.ETag
	}
	reqBody, err := getBody(resource)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource: %w", err)
	}
	headers := client.GetJSONHeadersWithEtag(etag)
	resp, err := client.DoRequest(Options{
		Method:  "PUT",
		URL:     href,
		Headers: headers,
		Body:    reqBody,
		Timeout: MAX_REQUEST_TIMEOUT_SEC * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("update request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		statusError := NewStatusError(resp.StatusCode, string(resp.Body))
		return nil, fmt.Errorf("update %s failed: %w", config.ResourceType, statusError)
	}

	return resp, nil
}

// DeleteResource deletes a resource from the SMC system by name using
// generic CRUD operations
func DeleteResource(config *GenericCRUDConfig, resourceName string) error {
	// First, read the resource to get ETag and Href
	client := config.Client

	resourceResp, err := ReadResourceByName(config, resourceName)
	if err != nil {
		return fmt.Errorf("failed to read %s for deletion: %w", config.ResourceType, err)
	}
	if resourceResp == nil {
		return fmt.Errorf("%s not found for deletion: %s", config.ResourceType, resourceName)
	}

	headers := client.GetAuthHeadersWithEtag(resourceResp.ETag)
	resp, err := client.DoRequest(Options{
		Method:  "DELETE",
		URL:     resourceResp.Location,
		Headers: headers,
		Timeout: MAX_REQUEST_TIMEOUT_SEC * time.Second,
	})
	if err != nil {
		return fmt.Errorf("delete request failed: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("delete %s failed with status %d: %s", config.ResourceType, resp.StatusCode, string(resp.Body))
	}

	return nil
}

// DeleteResourceByHref deletes a resource from the SMC system by href
// using generic CRUD operations.
// On error (eg 409) a StatusError is returned.
func DeleteResourceByHref(config *GenericCRUDConfig, href, etag string) error {
	client := config.Client
	headers := client.GetAuthHeadersWithEtag(etag)
	resp, err := client.DoRequest(Options{
		Method:  "DELETE",
		URL:     href,
		Headers: headers,
		Timeout: MAX_REQUEST_TIMEOUT_SEC * time.Second,
	})
	if err != nil {
		return fmt.Errorf("delete request failed: %w", err)
	}

	if resp.StatusCode == http.StatusOK ||
		resp.StatusCode == http.StatusNoContent {
		return nil
	}

	return NewStatusError(resp.StatusCode, fmt.Sprintf("delete %s failed: %s", config.ResourceType, string(resp.Body)))
}

// SearchElements performs a search and returns all matching elements
// without fetching their details
func SearchElements(config *GenericCRUDConfig, namePattern string) ([]SearchResult, error) {
	client := config.Client

	encodedName := neturl.QueryEscape(namePattern)
	searchURL := fmt.Sprintf("%s/%s/elements/%s?filter=%s&exact_match=true", client.BaseUrl, client.APIVersion, config.ResourceType, encodedName)
	headers := client.GetJSONHeaders()

	searchResp, err := client.DoRequest(Options{
		Method:  "GET",
		URL:     searchURL,
		Headers: headers,
		Timeout: MAX_REQUEST_TIMEOUT_SEC * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("search request failed: %w", err)
	}

	if searchResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search failed with status %d: %s", searchResp.StatusCode, string(searchResp.Body))
	}

	var searchResponse SearchResponse
	if err := json.Unmarshal(searchResp.Body, &searchResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal search response: %w", err)
	}

	return searchResponse.Result, nil
}
