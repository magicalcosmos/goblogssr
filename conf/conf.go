package conf

import "github.com/magicalcosmos/goblogssr/common/tlog"

// TemplateVar template variable
type TemplateVar struct {
	Key  string `toml:"key"`
	Type string `toml:"type"`
}

// Config config
type Config struct {
	Host            string        `toml:"host"`
	Log             tlog.Config   `toml:"Log"`
	Env             string        `toml:"env"`
	V8MaxCount      int32         `toml:"v8_maxcount"`
	V8LifeTime      int           `toml:"v8_lifetime"`
	ClientPath      string        `toml:"client_path"`
	StaticURLPath   string        `toml:"static_url_path"`
	InternalAPIHost string        `toml:"internal_api_host"`
	InternalAPIIp   string        `toml:"internal_api_ip"`
	InternalAPIPort int32         `toml:"internal_api_port"`
	IsAPIDelegate   bool          `toml:"is_api_delegate"`
	TemplateName    string        `toml:"template_name"`
	ClientCookie    string        `toml:"client_cookie"`
	RedirectOnerror string        `toml:"redirect_onerror"`
	SsrCtx          []string      `toml:"ssr_ctx"`
	TemplateVars    []TemplateVar `toml:"template_vars"`
}
