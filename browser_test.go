package browser

import (
	"http"
	"testing"
)

var ieUserAgentTests = [...]string{
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 7.1; Trident/5.0)",
}

var firefoxUserAgentTests = [...]string{
	"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.7) Gecko/20100809 Fedora/3.6.7-1.fc14 Firefox/3.6.7",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:2.0b11pre) Gecko/20110128 Firefox/4.0b11pre",
}

func makeReq(userAgent string) *http.Request {
	r := new(http.Request)
	r.UserAgent = userAgent
	return r
}

func TestIsIE(t *testing.T) {
	for _, agent := range ieUserAgentTests {
		if !IsIE(makeReq(agent)) {
			t.Fatal(agent, "Incorrectly identified as a non-IE user agent.")
		}
	}
}

func TestIsFF(t *testing.T) {
	for _, agent := range firefoxUserAgentTests {
		if !IsFirefox(makeReq(agent)) {
			t.Fatal(agent, "Incorrectly identified as a non-Firefox user agent")
		}
	}
}
