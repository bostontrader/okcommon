package utils

// Each response should have these headers.
var ExpectedResponseHeaders = map[string]string{
	"X-Ratelimit-Remaining-Second": "",
	"X-Kong-Proxy-Latency":         "",
	"X-Ratelimit-Limit-Second":     "",
	"X-Kong-Upstream-Latency":      "",
	"Set-Cookie":                   "",
	"X-Xss-Protection":             "",
	"X-Frame-Options":              "",
	"Via":                          "",
	"Content-Type":                 "",
	"Pragma":                       "",
	"X-Brokerid":                   "",
	"Content-Length":               "",
	"X-Content-Type-Options":       "",
	"Date":                         "",
	"Cache-Control":                "",
	"Expires":                      "",
}
