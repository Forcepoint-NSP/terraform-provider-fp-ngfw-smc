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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
	"testing"
)

type GenericCRUDConfigMock struct {
	mock.Mock
}

func (m *GenericCRUDConfigMock) ReadResourceByHref(href string) (*smc.ResponseData, error) {
	args := m.Called(href)
	return args.Get(0).(*smc.ResponseData), args.Error(1)
}

func TestGetHrefFromLinks(t *testing.T) {
	const resp = `{
	"result": [
		{
			"href": "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/26843901",
			"name": "myfw1",
			"type": "internal_gateway"
		},
		{
			"href": "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/26843902",
			"name": "myfw2",
			"type": "internal_gateway"
		}
	]
}`
	href, err := getHrefFromResult([]byte(resp), "myfw1")
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/26843901", href)

	href, err = getHrefFromResult([]byte(resp), "*2")
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/26843902", href)

	href, err = getHrefFromResult([]byte(resp), "*3")
	assert.Error(t, err)
}

func TestGetHrefFromJP(t *testing.T) {
	const resp = `{
    "interface_id": "0",
    "interfaces": [
        {
            "single_node_interface": {
                "address": "192.168.100.14",
                "key": 268546893,
                "link": [
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268587524/physical_interface/268498965/interface/268546893/contact_addresses",
                        "rel": "contact_addresses"
                    }
                ],
                "network_value": "192.168.100.0/24",
                "nicid": "0",
                "nodeid": 1
            }
        }
    ],
    "key": 268498965,
    "link": [
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268587524/physical_interface/268498965",
            "rel": "self",
            "type": "physical_interface"
        }
    ]
}
`
	href, err := getHrefFromJP([]byte(resp), "interfaces[0].single_node_interface.link[0].href")
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8082/7.4/elements/single_fw/268587524/physical_interface/268498965/interface/268546893/contact_addresses", href)

	href, err = getHrefFromJP([]byte(resp), "interfaces[?single_node_interface.address=='192.168.100.14']|[0].single_node_interface.link[?rel=='contact_addresses']|[0].href")
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8082/7.4/elements/single_fw/268587524/physical_interface/268498965/interface/268546893/contact_addresses", href)

	href, err = getHrefFromJP([]byte(resp), "key")
	assert.EqualError(t, err, "result is not a string")

	href, err = getHrefFromJP([]byte(resp), "not_a_valid_field")
	assert.EqualError(t, err, "jmespath expression returned no result")

	href, err = getHrefFromJP([]byte(resp), "interfaces[0.single_node_interface")
	assert.ErrorContains(t, err, "jmespath search failed: SyntaxError")

}

func TestResolveRef(t *testing.T) {
	crudConfigMock := new(GenericCRUDConfigMock)
	crudConfigMock.On("ReadResourceByHref",
		"http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway").Return(
		&smc.ResponseData{
			Body: []byte(`{
					"result": [
						{
							"href": "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099",
							"name": "mygateway",
							"type": "internal_gateway"
						}
					]
				}`),
		}, nil)

	crudConfigMock.On("ReadResourceByHref",
		"http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099").Return(
		&smc.ResponseData{
			Body: []byte(`
				{
					"admin_domain": "http://localhost:8082/7.4/elements/admin_domain/1",
					"key": 268439099,
					"link": [
						{
							"href": "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099",
							"rel": "self",
							"type": "internal_gateway"
						},
						{
							"href": "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099/internal_endpoint",
							"rel": "internal_endpoint",
							"type": "internal_endpoint"
						}
					]
				}`),
		}, nil)

	crudConfigMock.On("ReadResourceByHref",
		"http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099/internal_endpoint").Return(
		&smc.ResponseData{
			Body: []byte(`
				{
					"result": [
						{
							"href": "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099/internal_endpoint/268475570",
							"name": "192.168.64.254",
							"type": "internal_endpoint"
						},
						{
							"href": "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099/internal_endpoint/268475193",
							"name": "192.168.51.254",
							"type": "internal_endpoint"
						}
					]
                 }
				`),
		}, nil)

	ref, err := ResolveRef(context.Background(), crudConfigMock,
		"http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway#mygateway")
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099", ref)

	ref, err = ResolveRef(context.Background(), crudConfigMock,
		"http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway#bad")
	assert.Error(t, err)
	assert.Equal(t, "", ref)
	assert.Contains(t, err.Error(), "failed to resolve name part 'bad'")

	ref, err = ResolveRef(context.Background(), crudConfigMock,
		"http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway#mygateway/internal_endpoint")
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099/internal_endpoint", ref)

	ref, err = ResolveRef(context.Background(), crudConfigMock,
		"http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway#mygateway/internal_endpoint/192.168.64.254")
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099/internal_endpoint/268475570", ref)

	// same with wildcard instead of name for internal_gateway
	ref, err = ResolveRef(context.Background(), crudConfigMock,
		"http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway#*/internal_endpoint/192.168.64.254")
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099/internal_endpoint/268475570", ref)

}
