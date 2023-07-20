// Copyright 2020-present, lizc2003@gmail.com
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

package v8

import (
	"fmt"
	"github.com/magicalcosmos/goblogssr/common/tlog"
	"github.com/magicalcosmos/goblogssr/v8worker"
)

const (
	MSGTYPE_PRINT_DEBUG   = 0
	MSGTYPE_PRINT_INFO    = 1
	MSGTYPE_PRINT_WARN    = 2
	MSGTYPE_PRINT_ERROR   = 3
	MSGTYPE_REQUIRE       = 10
	MSGTYPE_HTTP          = 11
	MSGTYPE_HTTP_CALLBACK = 20

	MSGTYPE_SSR_HTML_OK   = 80
	MSGTYPE_SSR_HTML_FAIL = 81
	MSGTYPE_SSR_CSS       = 82
	MSGTYPE_SSR_META      = 83

	MSGTYPE_SET_URL = 101
)

func IsDevEnvironment(env string) bool {
	return env == "dev"
}

var gInitJs string

func initV8NewJs() {
	gInitJs = globalJsContent + callbackJsContent + xmlHttpRequestJsContent + nativeModuleJsContent
}

func newV8Worker(appEnv string) (*v8worker.Worker, error) {
	w := v8worker.New(v8WorkerSendCallback, v8WorkerRequestCallback)
	w.Acquire()

	nodeEnv := "production"
	if IsDevEnvironment(appEnv) {
		nodeEnv = "development"
	}
	err := w.Execute("env.js", fmt.Sprintf(`
		this.process = { env: { VUE_ENV: "server", NODE_ENV: "%s" }};
		this.global = { process: process};
		this.APP_ENV = "%s";
	`, nodeEnv, appEnv))

	if err != nil {
		w.Dispose()
		return nil, err
	}

	err = w.Execute("init.js", gInitJs)
	if err != nil {
		w.Dispose()
		return nil, err
	}

	err = loadMainModule(w, "v8main.js")
	if err != nil {
		w.Dispose()
		return nil, err
	}

	w.Release()
	return w, nil
}

func v8WorkerSendCallback(w *v8worker.Worker, mtype int, msg string, userdata int64) {
	switch mtype {
	case MSGTYPE_PRINT_DEBUG:
		tlog.Debug(msg)
	case MSGTYPE_PRINT_INFO:
		tlog.Info(msg)
	case MSGTYPE_PRINT_WARN:
		tlog.Warning(msg)
	case MSGTYPE_PRINT_ERROR:
		tlog.Error(msg)
	default:
		TheV8Mgr.SendCallback(mtype, msg, userdata)
	}
}

func v8WorkerRequestCallback(w *v8worker.Worker, mtype int, msg string) string {
	switch mtype {
	case MSGTYPE_HTTP:
		return processXMLHttpRequestCmd(w, msg)
	case MSGTYPE_REQUIRE:
		return requireModule(msg)
	}
	return ""
}
