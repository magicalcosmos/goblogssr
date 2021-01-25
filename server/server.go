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
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/magicalcosmos/goblogssr/common/conf"
	"github.com/magicalcosmos/goblogssr/common/tlog"
	"github.com/magicalcosmos/goblogssr/common/util"
	v8 "github.com/magicalcosmos/goblogssr/v8"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
)

const (
	DIST_DIR_WWW    = "public"
	DIST_DIR_SERVER = "server_dist"
)

// Server server
type Server struct {
	RequestMgr      *RequestMgr
	V8Mgr           *v8.V8Mgr
	HostPort        int
	ClientPath      string
	Env             string
	IsAPIDelegate   bool
	ClientCookie    string
	RedirectOnerror string
	SsrTemplate     string
	SsrCtx          []string
	TemplateVars    map[string]string
	TemplateURLEnv  string
}

// ThisServer this server
var ThisServer *Server

// NewServer new server
func NewServer(c *conf.Config) error {
	templateVars := map[string]string{"State": "js"}
	for _, v := range c.TemplateVars {
		if _, ok := templateVars[v.Key]; !ok {
			templateVars[v.Key] = v.Type
		}
	}

	clientPath := getClientPath(c.ClientPath)
	if clientPath == "" {
		return errors.New("Error: the path of js project is empty")
	}
	c.ClientPath = clientPath

	tmp := strings.Index(c.Host, ":")
	hostPort := int(util.StringToInt64(c.Host[tmp+1:], 0))

	if c.InternalApiPort == 0 {
		c.InternalApiPort = 80
	}

	ThisServer = &Server{
		RequestMgr:      NewRequestMgr(),
		HostPort:        hostPort,
		ClientPath:      c.ClientPath,
		SsrTemplate:     c.TemplateName,
		ClientCookie:    c.ClientCookie,
		RedirectOnerror: c.RedirectOnerror,
		Env:             c.Env,
		SsrCtx:          c.SsrCtx,
		IsAPIDelegate:   c.IsAPIDelegate,
		TemplateVars:    templateVars,
	}

	v8mgr, err := newV8Mgr(c)
	if err != nil {
		tlog.Fatal(err)
		return err
	}
	ThisServer.V8Mgr = v8mgr

	handler := getHTTPHandler(c)
	fmt.Println(util.FormatFullTime(time.Now()), "running ...")
	return gracehttp.Serve(&http.Server{Addr: c.Host, Handler: handler})
}

func getHTTPHandler(c *conf.Config) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()

	localStaticPath := ThisServer.ClientPath + DIST_DIR_WWW
	e.Use(GetStaticAndProxyHandler(c.StaticUrlPath, localStaticPath))
	e.StaticFile("/favicon.ico", localStaticPath+"/favicon.ico")
	e.LoadHTMLGlob(c.ClientPath + DIST_DIR_SERVER + "/template/*")

	e.NoRoute(HandleSsrRequest)
	return e
}

func getClientPath(clientPath string) string {
	if clientPath == "" {
		return ""
	}
	if clientPath[0] != '/' {
		basepath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return ""
		}
		clientPath = basepath + "/" + clientPath
	}
	if clientPath[len(clientPath)-1] != '/' {
		clientPath += "/"
	}
	return clientPath
}

func newV8Mgr(c *conf.Config) (*v8.V8Mgr, error) {
	serverPath := c.ClientPath + DIST_DIR_SERVER + "/"
	serverPathMain := serverPath + "g/"
	vuePath := c.ClientPath + "node_modules/"

	v8Conf := v8.V8MgrConfig{
		JsPaths:         []string{serverPathMain, serverPath, vuePath},
		SendCallback:    v8SendCallback,
		Env:             c.Env,
		MaxWorkerCount:  c.V8MaxCount,
		WorkerLifeTime:  c.V8LifeTime,
		InternalApiHost: c.InternalAPIHost,
		InternalApiIp:   c.InternalAPIIp,
		InternalApiPort: c.InternalAPIPort}
	return v8.NewV8Mgr(&v8Conf)
}
