package server

import (
	"encoding/json"
	"fmt"
	"github.com/magicalcosmos/goblogssr/common/tlog"
	v8 "github.com/magicalcosmos/goblogssr/v8"
	"net/url"
	"strconv"
)

func v8SendCallback(mtype int, msg string, reqId int64) {
	switch mtype {
	case v8.MSGTYPE_SET_URL:
		setTemplateEnv(msg)
	default:
		req := ThisServer.RequstMgr.GetRequest(reqId)
		if req != nil {
			switch mtype {
			case v8.MSGTYPE_SSR_HTML_OK:
				req.bOK = true
				fallthrough
			case v8.MSGTYPE_SSR_HTML_FAIL:
				req.result.Html = msg
				req.wg.Done()
			case v8.MSGTYPE_SSR_CSS:
				req.result.Css = msg
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
	if len(ThisServer.TemplateUrlEnv) == 0 {
		var dat map[string]string
		err := json.Unmarshal([]byte(msg), &dat)
		var baseUrl string
		var apiUrl string
		if err == nil {
			if v, ok := dat["base"]; ok {
				baseUrl = v
			}
			if v, ok := dat["api"]; ok {
				apiUrl = v
			}
		}
		if baseUrl != "" && apiUrl != "" {
			if ThisServer.IsApiDelegate {
				u, err := url.Parse(apiUrl)
				if err != nil {
					tlog.Error(err)
					return
				}
				u.Host = u.Hostname() + ":" + strconv.FormatInt(int64(ThisServer.HostPort), 10)
				apiUrl = u.String()
			}

			ThisServer.TemplateUrlEnv = fmt.Sprintf(`window.APP_ENV="%s";
			window.BASE_URL="%s";
			window.API_BASE_URL="%s";`, ThisServer.Env, baseUrl, apiUrl)
		}
	}
}
