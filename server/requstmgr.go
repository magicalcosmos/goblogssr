// Copyright 2021 brodyliao@gmail.com

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"sync"
)

// SsrResult ssr result
type SsrResult struct {
	HTML string            `json:"html"`
	CSS  string            `json:"css"`
	Meta map[string]string `json:"meta"`
}

// Request request
type Request struct {
	wg     sync.WaitGroup
	reqID  int64
	result SsrResult
	bOK    bool
}

// RequestMgr request manager
type RequestMgr struct {
	mutex    sync.Mutex
	requests map[int64]*Request
	maxID    int64
}

// NewRequestMgr new request manager
func NewRequestMgr() *RequestMgr {
	return &RequestMgr{requests: make(map[int64]*Request)}
}

// NewRequest new request
func (that *RequestMgr) NewRequest() *Request {
	req := &Request{}
	req.wg.Add(1)

	that.mutex.Lock()
	that.maxID++
	req.reqID = that.maxID
	that.requests[req.reqID] = req
	that.mutex.Unlock()
	return req
}

// DestroyRequest destroy request
func (that *RequestMgr) DestroyRequest(reqID int64) {
	that.mutex.Lock()
	if _, ok := that.requests[reqID]; ok {
		delete(that.requests, reqID)
	}
	that.mutex.Unlock()
}

// GetRequest get request
func (that *RequestMgr) GetRequest(reqID int64) *Request {
	var req *Request
	that.mutex.Lock()
	if v, ok := that.requests[reqID]; ok {
		req = v
	}
	that.mutex.Unlock()
	return req
}
