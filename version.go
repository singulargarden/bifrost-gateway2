package main

import (
	"net/http"
	"runtime/debug"
	"time"
)

var name = "bifrost-gateway"
var version = buildVersion()
var userAgent = name + "/" + version

func buildVersion() string {
	var revision string
	var day string
	var dirty bool

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "dev-build"
	}
	for _, kv := range info.Settings {
		switch kv.Key {
		case "vcs.revision":
			revision = kv.Value[:7]
		case "vcs.time":
			t, _ := time.Parse(time.RFC3339, kv.Value)
			day = t.UTC().Format("2006-01-02")
		case "vcs.modified":
			dirty = kv.Value == "true"
		}
	}
	if dirty {
		revision += "-dirty"
	}
	if revision != "" {
		return day + "-" + revision
	}
	return "dev-build"
}

type withUserAgent struct {
	http.RoundTripper
}

func (adt *withUserAgent) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", userAgent)
	return adt.RoundTripper.RoundTrip(req)
}
