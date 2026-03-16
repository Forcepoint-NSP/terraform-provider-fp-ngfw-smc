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

package smc

import (
	neturl "net/url"
	"sync"
)

type LockRequest struct {
	lock sync.RWMutex
}

// LockRequests manages locks for concurrent requests
type LockRequests struct {
	requestLocksByBaseUrl map[string]*LockRequest
}

// Singleton instance of LockRequests
var instance = &LockRequests{
	requestLocksByBaseUrl: make(map[string]*LockRequest),
}

var internalLock sync.Mutex

func GetLockRequests() *LockRequests {
	return instance
}

func (mgr *LockRequests) get(lockId string) *LockRequest {

	internalLock.Lock()
	defer internalLock.Unlock()

	lock := mgr.requestLocksByBaseUrl[lockId]
	if lock == nil {
		lock = &LockRequest{}
		mgr.requestLocksByBaseUrl[lockId] = lock
	}
	return lock
}

func GetLockId(url string) string {
	parsedURL, err := neturl.Parse(url)
	if err != nil {
		// TODO no tracing in this version - to be efficient we need
		// the tf.logs here
		return "DEFAULT_LOCK_FROM_URL"
	}
	return parsedURL.Scheme + "://" + parsedURL.Host
}

func (mgr *LockRequests) Enter(lockId string) {
	mgr.get(lockId).lock.Lock()
}

func (mgr *LockRequests) Leave(lockId string) {
	mgr.get(lockId).lock.Unlock()
}
