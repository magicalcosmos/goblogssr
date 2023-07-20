package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/magicalcosmos/goblogssr/common/tlog"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type RobotDingDing struct {
	env        string
	url        string
	secret     string
	httpClient *http.Client
}

// https://developers.dingtalk.com/document/app/custom-robot-access
func NewRobotDingDing(env string, app string, url string, secret string) *RobotDingDing {
	host, _ := os.Hostname()
	r := &RobotDingDing{
		env:        "env: " + env + "\n" + "host: " + host + "\n" + "app: " + app + "\n",
		url:        url,
		secret:     secret,
		httpClient: NewHttpClient(false),
	}
	return r
}

func (this *RobotDingDing) SendMsg(msg string) string {
	var ddurl string
	if this.secret != "" {
		timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
		strSign := timestamp + "\n" + this.secret

		h := hmac.New(sha256.New, []byte(this.secret))
		h.Write([]byte(strSign))
		sign := url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
		ddurl = this.url + "&timestamp=" + timestamp + "&sign=" + sign
	} else {
		ddurl = this.url
	}

	type msgText struct {
		Content string `json:"content"`
	}
	type msgData struct {
		Msgtype string  `json:"msgtype"`
		Text    msgText `json:"text"`
	}

	b := strings.Builder{}
	b.WriteString(this.env)
	b.WriteString("time: ")
	b.WriteString(FormatFullTime(time.Now()))
	b.WriteByte('\n')
	b.WriteString(msg)

	data := msgData{Msgtype: "text", Text: msgText{Content: b.String()}}
	var ret string
	err := HttpPost(this.httpClient, ddurl, nil, data, &ret)
	if err != nil {
		tlog.Error(err)
	}
	return ret
}
