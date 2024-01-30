package config

var config = map[string]any{
	"APP_PORT":     8081,
	"APP_DEBUG":    true,
	"APP_TIMEZONE": "Asia/Shanghai",

	// log
	"LOG_ENABLE": true,
	"LOG_PATH":   RootPath() + "/log/",

	// session.
	"SESSION_ENABLE":          true,
	"SESSION_COOKIE_NAME":     "p_sid",
	"SESSION_COOKIE_LIFETIME": 3600,

	// basic auth.
	"BASIC_AUTH_ENABLE":   true,
	"BASIC_AUTH_USERNAME": "admin",
	"BASIC_AUTH_PASSWORD": "admin",

	// mb
	"REQUEST_BODY_LIMIT": 8,

	// database
	"MYSQL_HOST":     "yes.server",
	"MYSQL_PORT":     3306,
	"MYSQL_DATABASE": "test_main",
	"MYSQL_USERNAME": "root",
	"MYSQL_PASSWORD": "root",

	// template
	"TEMPLATE_WRAPPER_STR": "__WRAPPER__",
	"TEMPLATE_PATH":        RootPath() + "/app/http/view/",
	"TEMPLATE_EXT":         "html",
}
