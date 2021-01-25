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
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/magicalcosmos/goblogssr/common/tlog"
	v8 "github.com/magicalcosmos/goblogssr/v8"
)

func v8SendCallback(mimeType int, msg string, reqID int64) {
	switch mimeType {
	case v8.MSGTYPE_SET_URL:
		setTemplateEnv(msg)
	default:
		req := ThisServer.RequestMgr.GetRequest(reqID)
		if req != nil {
			switch mimeType {
			case v8.MSGTYPE_SSR_HTML_OK:
				req.bOK = true
				fallthrough
			case v8.MSGTYPE_SSR_HTML_FAIL:
				req.result.HTML = msg
				req.wg.Done()
			case v8.MSGTYPE_SSR_CSS:
				req.result.CSS = msg
			case v8.MSGTYPE_SSR_META:
				var meta map[string]string
				err := json.Unmarshal([]byte(msg), &meta)
				if err != nil {
					tlog.Error(err)
				} else {
					req.result.Meta = meta
				}
			}
		}
	}
}

func setTemplateEnv(msg string) {
	if len(ThisServer.TemplateURLEnv) == 0 {
		var dat map[string]string
		err := json.Unmarshal([]byte(msg), &dat)
		var baseURL string
		var apiURL string
		if err == nil {
			if v, ok := dat["base"]; ok {
				baseURL = v
			}
			if v, ok := dat["api"]; ok {
				apiURL = v
			}
		}
		if baseURL != "" && apiURL != "" {
			if ThisServer.IsAPIDelegate {
				u, err := url.Parse(apiURL)
				if err != nil {
					tlog.Error(err)
					return
				}
				u.Host = u.Hostname() + ":" + strconv.FormatInt(int64(ThisServer.HostPort), 10)
				apiURL = u.String()
			}

			ThisServer.TemplateURLEnv = fmt.Sprintf(`window.APP_ENV="%s";
			window.BASE_URL="%s";
			window.API_BASE_URL="%s";`, ThisServer.Env, baseURL, apiURL)
		}
	}
}
