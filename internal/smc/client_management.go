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

// ClientManager keeps track of SMC clients. This allows to cleanup
// the resources when the process exits.

package smc

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"sync"
	"time"
)

// ClientInfo wraps a Client with additional live information
type ClientInfo struct {
	Client    *SmcClient
	CreatedAt time.Time
}

// ClientManager is a singleton that manages all client connections by
// MgtServerId. Note that multiple clients are not supported yet, but
// the singleton allows to check if the client is already logged in
// rather than "re login" each time.
type ClientManager struct {
	clients map[string]*ClientInfo
	lock    sync.RWMutex
}

// Singleton instance of ClientManager (Lazy initialization is not needed, it's empty on startup)
var managerInstance = &ClientManager{
	clients: make(map[string]*ClientInfo),
}

// GetClientManager returns the singleton instance of ClientManager
func GetClientManager() *ClientManager {
	return managerInstance
}

// GetClient returns a client for a given url -- return nil if no
// applicable client are available.

// param:
// validLogin: if true, only return a client that is logged in
func (clientMgr *ClientManager) GetClient(ctx context.Context,
	url string, validLogin bool) *SmcClient {
	clientMgr.lock.RLock()
	defer clientMgr.lock.RUnlock()

	clientInfo, exists := clientMgr.clients[url]
	if !exists {
		tflog.Debug(ctx, fmt.Sprintf("ClientManager: no client found for URL %s", url))
		return nil
	}

	if validLogin && !clientInfo.Client.LoginConfirmed {
		tflog.Debug(ctx, fmt.Sprintf("ClientManager: client found for %s was not logged in", url))
		return nil
	}

	if validLogin && !clientInfo.Client.ping(ctx) {
		tflog.Debug(ctx, fmt.Sprintf("ClientManager: existing session do not get any answer on URL %s", url))
		return nil
	}

	// Existing context may be "cancelled", upgrade to the new one
	clientInfo.Client.UpdateLogContext(ctx, url)

	return clientInfo.Client
}

// AddClient adds or updates a client for a given MgtServerId
func (clientMgr *ClientManager) AddClient(URL string, client *SmcClient) error {
	if URL == "" {
		return fmt.Errorf("MgtServerId cannot be empty")
	}
	if client == nil {
		return fmt.Errorf("client cannot be nil")
	}

	clientMgr.lock.Lock()
	defer clientMgr.lock.Unlock()

	// Update the log context to include the MgtServerId
	if client.SmcContext != nil {
		client.UpdateLogContext(client.SmcContext.Context(), URL)
	}
	clientMgr.clients[URL] = &ClientInfo{
		Client:    client,
		CreatedAt: time.Now(),
	}
	client.SmcContext.Debug(fmt.Sprintf("SMC API Connection for %s registered at %s", URL, clientMgr.clients[URL].CreatedAt))
	return nil
}

// RemoveClient removes a client by MgtServerId
func (clientMgr *ClientManager) RemoveClient(mgtServerId string) {
	clientMgr.lock.Lock()
	defer clientMgr.lock.Unlock()

	delete(clientMgr.clients, mgtServerId)
}

// HasClient checks if a client exists for the given MgtServerId
func (clientMgr *ClientManager) HasClient(mgtServerId string) bool {
	clientMgr.lock.RLock()
	defer clientMgr.lock.RUnlock()

	_, exists := clientMgr.clients[mgtServerId]
	return exists
}

// ListClients returns a list of all MgtServerIds currently managed
func (clientMgr *ClientManager) ListClients() []string {
	clientMgr.lock.RLock()
	defer clientMgr.lock.RUnlock()

	ids := make([]string, 0, len(clientMgr.clients))
	for id := range clientMgr.clients {
		ids = append(ids, id)
	}
	return ids
}

// CleanupAll logs out all clients and clears the manager
func (clientMgr *ClientManager) CleanupAll(ctx context.Context) {
	clientMgr.lock.Lock()
	defer clientMgr.lock.Unlock()

	for mgtServerId, clientInfo := range clientMgr.clients {
		if clientInfo.Client != nil && clientInfo.Client.LoginConfirmed {
			// Attempt to logout
			if err := clientInfo.Client.Logout(ctx); err != nil {
				// Log error but continue with other clients
				if clientInfo.Client.SmcContext != nil {
					clientInfo.Client.SmcContext.Warn("Failed to logout client during cleanup", map[string]interface{}{
						"mgt_server_id": mgtServerId,
						"error":         err.Error(),
					})
				}
			}
		}
	}

	// Clear all clients
	clientMgr.clients = make(map[string]*ClientInfo)
}

// GetOrCreateClient creates a new Client from an Auth struct
func GetOrCreateClient(ctx context.Context,
	url string, apiVersion string, apikey string,
	verifySSL bool, trustedCert string, domain string) (*SmcClient, error) {

	// todo normalize URL
	// Check if a client session already exists for this server
	if existingClient := GetClientManager().GetClient(ctx, url, true); existingClient != nil {
		existingClient.SmcContext.Info(fmt.Sprintf("Reusing existing SMC client session === %s", url), map[string]any{
			"url":         url,
			"api_version": apiVersion,
		})
		return existingClient, nil
	}

	client, err := NewClient(ctx, url, apiVersion, apikey,
		verifySSL, trustedCert, domain)
	if err != nil {
		return nil, err
	}
	GetClientManager().AddClient(url, client)
	return client, nil
}
