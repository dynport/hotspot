package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	br "github.com/dynport/dgtk/browser"
	"github.com/dynport/gocli"
)

const (
	railControllerUrl = "http://2nd.railnet.train/railnet/php_logic/T_railcontroller.php?include_id=1&lang=de_DE"
	connectUrl        = "https://hotspot.t-mobile.net/wlan/start.do"
	disconnectUrl     = "https://hotspot.t-mobile.net/wlan/stop.do"
	checkHost         = "www.heise.de"
	envLogin          = "TMOBILE_LOGIN"
	envPwd            = "TMOBILE_PWD"
)

var logger = log.New(os.Stderr, "", 0)

func online() (bool, error) {
	rsp, e := http.Head("http://" + checkHost)
	if e != nil {
		return false, e
	}
	return rsp.Request.URL.Host == checkHost, nil
}

var cachedBrowser *br.Browser

func browser() *br.Browser {
	if cachedBrowser == nil {
		b, e := br.New()
		if e != nil {
			logger.Fatal("unable to initialize browser: %q", e)
		}
		cachedBrowser = b
		return cachedBrowser
	}
	return cachedBrowser
}

func getEnv(key string) (string, error) {
	if v := os.Getenv(key); v != "" {
		return v, nil
	} else {
		return "", fmt.Errorf("key %q not found in env", key)
	}
}

func red(s string) string {
	return gocli.Red(s)
}

func green(s string) string {
	return gocli.Green(s)
}

func yellow(s string) string {
	return gocli.Yellow(s)
}
