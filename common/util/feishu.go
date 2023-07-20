package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/magicalcosmos/goblogssr/common/tlog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type RobotFeishu struct {
	env        string
	url        string
	secret     string
	httpClient *http.Client
}

// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func NewRobotFeishu(env string, app string, url string, secret string) *RobotFeishu {
	host, _ := os.Hostname()
	r := &RobotFeishu{
		env:        "env: " + env + "\n" + "host: " + host + "\n" + "app: " + app + "\n",
		url:        url,
		secret:     secret,
		httpClient: NewHttpClient(false),
	}
	return r
}

func (this *RobotFeishu) SendMsg(msg string) string {
	var timestamp string
	var sign string
	if this.secret != "" {
		timestamp = strconv.FormatInt(time.Now().UnixMilli(), 10)
		strSign := timestamp + "\n" + this.secret

		h := hmac.New(sha256.New, []byte(strSign))
		h.Write([]byte{})
		sign = base64.StdEncoding.EncodeToString(h.Sum(nil))
	}

	type msgText struct {
		Text string `json:"text"`
	}

	type msgData struct {
		Timestamp string  `json:"timestamp,omitempty"`
		Sign      string  `json:"sign,omitempty"`
		Msgtype   string  `json:"msg_type"`
		Content   msgText `json:"content"`
	}

	b := strings.Builder{}
	b.WriteString(this.env)
	b.WriteString("time: ")
	b.WriteString(FormatFullTime(time.Now()))
	b.WriteByte('\n')
	b.WriteString(msg)

	data := msgData{Timestamp: timestamp, Sign: sign,
		Msgtype: "text", Content: msgText{Text: b.String()}}

	var ret string
	err := HttpPost(this.httpClient, this.url, nil, data, &ret)
	if err != nil {
		tlog.Error(err)
	}
	return ret
}
