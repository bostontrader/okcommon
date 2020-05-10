package utils

// Each response should have these headers.
var ExpectedResponseHeaders = map[string]string{
	"Cache-Control":                "",
	"Content-Length":               "",
	"Content-Type":                 "",
	"Date":                         "",
	"Expires":                      "",
	"Set-Cookie":                   "",
	"Pragma":                       "",
	"Via":                          "",
	"X-Brokerid":                   "",
	"X-Content-Type-Options":       "",
	"X-Frame-Options":              "",
	"X-Kong-Proxy-Latency":         "",
	"X-Kong-Upstream-Latency":      "",
	"X-Ratelimit-Limit-Second":     "",
	"X-Ratelimit-Remaining-Second": "",
	"X-Xss-Protection":             "",
}

// Some responses return all of the above but omit X-Ratelimit*
var ExpectedResponseHeadersB = map[string]string{
	"Cache-Control":           "",
	"Content-Length":          "",
	"Content-Type":            "",
	"Date":                    "",
	"Expires":                 "",
	"Set-Cookie":              "",
	"Pragma":                  "",
	"Via":                     "",
	"X-Brokerid":              "",
	"X-Content-Type-Options":  "",
	"X-Frame-Options":         "",
	"X-Kong-Proxy-Latency":    "",
	"X-Kong-Upstream-Latency": "",
	"X-Xss-Protection":        "",
}
