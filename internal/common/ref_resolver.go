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

package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmespath/go-jmespath"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
	"github.com/tidwall/gjson"
	"net/url"
	"strings"
)

// getHrefFromResult extracts the first href from a JSON response body based on a name wildcard.
// @param body: JSON response body as byte slice
// @param nameWildcard: Wildcard pattern to match the name field (supports '*' and '?' wildcards)
//
// Example JSON body:
//
//	   "result": [
//	    {
//	        "href": "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099",
//	        "name": "f-eu-hc1b8-fw",
//	        "type": "internal_gateway"
//	    }
//	]
func getHrefFromResult(body []byte, nameWildcard string) (string, error) {
	jsonPath := fmt.Sprintf("result.#(name%%\"%s\").href", nameWildcard)
	value := gjson.Get(string(body), jsonPath)
	if !value.Exists() {
		return "", fmt.Errorf("no resource found matching name wildcard: %s", nameWildcard)
	}
	return value.String(), nil
}

// getHrefFromLink extracts href from JSON body "link" element
// @param body: JSON response body as byte slice
// @param relWildCard: Wildcard pattern to match the rel field (supports '*' and '?' wildcards)
//
// Example JSON body:
//
//	"link": [
//	    {
//	        "rel": "self",
//	        "href": "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099"
//	    },
func getHrefFromLink(body []byte, relWildCard string) (string, error) {
	jsonPath := fmt.Sprintf("link.#(rel%%\"%s\").href", relWildCard)
	value := gjson.Get(string(body), jsonPath)
	if !value.Exists() {
		return "", fmt.Errorf("no resource found matching rel wildcard: %s", relWildCard)
	}
	return value.String(), nil
}

// getHrefFromJP extracts href from JSON body using a JMESPath expression
func getHrefFromJP(body []byte, jmespathExpr string) (string, error) {
	var data any
	err := json.Unmarshal(body, &data)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON body: %w", err)
	}
	result, err := jmespath.Search(jmespathExpr, data)
	if err != nil {
		return "", fmt.Errorf("jmespath search failed: %w", err)
	}

	if result == nil {
		return "", fmt.Errorf("jmespath expression returned no result")
	}

	href, ok := result.(string)
	if !ok {
		return "", fmt.Errorf("result is not a string")
	}
	return href, nil
}

/*
 * resolveRef2 resolves a URL by searching for hierarchical name parts
 *
 *  e.g. search for endpoint inside internal_gateway
 *
 *  from_ref = "http://mysmc:8082/7.4/elements/single_fw/268587556/internal_gateway"
 *  name = "mygateway/internal_endpoint/192.168.100.14"
 *  endpoint, err := resolveRef2(ctx, configReader, from_ref, name)
 *
 * return resolved href or error if not found
 */
func ResolveRef2(_ context.Context, configReader smc.IResourceReader,
	from_ref string, search_fragment string) (string, error) {
	// todo escape '/' in fragment parts
	parts := strings.Split(search_fragment, "/")
	href := from_ref
	for _, part := range parts {
		resp, err := configReader.ReadResourceByHref(href)
		if err != nil {
			return "", err
		}

		href, err = getHrefFromResult(resp.Body, part)
		if err == nil {
			continue
		}
		href, err = getHrefFromLink(resp.Body, part)
		if err == nil {
			continue
		}
		href, err = getHrefFromJP(resp.Body, part)
		if err == nil {
			continue
		}
		return "", fmt.Errorf("failed to resolve name part '%s': %w", part, err)
	}
	return href, nil
}

// resolveRef resolves a reference URL with optional fragment part
// containing the relative hierarchical name to resolve.

// return resolved href or error if not found
//
// e.g.
// name = "http://mysmc:8082/7.4/elements/single_fw/268587556/internal_gateway#mygateway/internal_endpoint/192.168.100.14"
// href, err := resolveRef2(ctx, configReader, name)
//
// returns e.g.
//
// http://localhost:8082/7.4/elements/single_fw/12345/internal_gateway/6789/internal_endpoint/101112

func ResolveRef(ctx context.Context, configReader smc.IResourceReader, url_with_search_fragment string) (string, error) {
	u, err := url.Parse(url_with_search_fragment)
	if err != nil {
		return "", err
	}

	search_fragment := u.Fragment
	// no fragment, nothing to resolve
	if search_fragment == "" {
		return url_with_search_fragment, nil
	}

	u.Fragment = ""
	from_ref := u.String()

	return ResolveRef2(ctx, configReader, from_ref, search_fragment)
}
