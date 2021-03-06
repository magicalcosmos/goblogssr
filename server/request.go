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
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/magicalcosmos/goblogssr/common/tlog"
	"github.com/magicalcosmos/goblogssr/common/util"
	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

// HandleSsrRequest handle ssr request
func HandleSsrRequest(c *gin.Context) {
	reqURL := c.Request.URL
	url := reqURL.Path
	if len(reqURL.RawQuery) > 0 {
		url += "?"
		url += reqURL.RawQuery
	}

	cookie := c.GetHeader("Cookie")
	if ThisServer.ClientCookie != "" {
		var clientID string
		cookieName := ThisServer.ClientCookie
		cookieVal, err := c.Request.Cookie(cookieName)
		if err == nil && len(cookieVal.Value) > 0 {
			clientID = cookieVal.Value
		} else {
			clientID = generateUUID() + strconv.FormatInt(int64(rand.Int31n(10)), 10)
			c.SetCookie(cookieName, clientID, 24*3600*365*10,
				"/", util.GetDomainFromHost(c.Request.Host), false, false)
			if len(cookie) > 0 {
				cookie = cookieName + "=" + clientID + "; " + cookie
			} else {
				cookie = cookieName + "=" + clientID
			}
		}
	}
	ssrCtx := map[string]string{"Cookie": cookie}
	for _, k := range ThisServer.SsrCtx {
		v := c.GetHeader(k)
		if v == "" && k == "X-Forwarded-For" {
			v = c.ClientIP()
		}
		ssrCtx[strings.ReplaceAll(k, "-", "_")] = v
	}

	tlog.Infof("http request: %s", url)
	result, bOK := generateSsrResult(url, ssrCtx)

	if !bOK && ThisServer.RedirectOnerror != "" && reqURL.Path != ThisServer.RedirectOnerror {
		tlog.Errorf("redirect: %s?%s", reqURL.Path, reqURL.RawQuery)
		c.Redirect(302, ThisServer.RedirectOnerror)
		return
	}

	outputHTML(c, result)
}

func outputHTML(c *gin.Context, result SsrResult) {
	templName := ThisServer.SsrTemplate
	templObj := gin.H{
		"Html":   template.HTML(result.HTML),
		"Css":    template.HTML(result.CSS),
		"UrlEnv": template.JS(ThisServer.TemplateURLEnv),
	}
	for k, v := range result.Meta {
		if v != "" {
			ktype := ThisServer.TemplateVars[k]
			switch ktype {
			case "js":
				templObj[k] = template.JS(v)
			case "html":
				templObj[k] = template.HTML(v)
			default:
				templObj[k] = v
			}
		}
	}
	c.HTML(http.StatusOK, templName, templObj)
}

func generateSsrResult(url string, ssrCtx map[string]string) (SsrResult, bool) {
	req := ThisServer.RequestMgr.NewRequest()

	headerJSON, _ := json.Marshal(ssrCtx)
	var jsCode strings.Builder
	jsCode.Grow(renderJsLength + len(headerJSON) + len(url) + 30)
	jsCode.WriteString(renderJsPart1)
	jsCode.WriteString(`{v8reqId:`)
	jsCode.WriteString(strconv.FormatInt(req.reqID, 10))
	jsCode.WriteString(`,url:"`)
	jsCode.WriteString(url)
	jsCode.WriteString(`",ssrCtx:`)
	jsCode.Write(headerJSON)
	jsCode.WriteString(`}`)
	jsCode.WriteString(renderJsPart2)

	//fmt.Println(jsCode.String())

	err := ThisServer.V8Mgr.Execute("bundle.js", jsCode.String())

	if err == nil {
		req.wg.Wait()
	} else {
		req.result.HTML = err.Error()
	}
	ThisServer.RequestMgr.DestroyRequest(req.reqID)

	return req.result, req.bOK
}

func generateUUID() string {
	uuid := uuid.NewV4().String()
	return strings.Replace(uuid, "-", "", -1)
}
