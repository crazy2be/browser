// A simple package to implement browser checking to allow mobile-specific, browser specific (Helllllo IE!), and platform specific logic.
package browser

import (
	"strings"
	"http"
	"fmt"
)

var mobileAgentSubstrings = [...]string{"iPod", "iPhone", "Mobile", "Phone", "Android"}[:]
var ieAgentSubstrings = [...]string{"MSIE"}[:]
var firefoxAgentSubstrings = [...]string{"Firefox"}[:]

func hasAnySubstring(agent string, substrings []string) bool {
	for _, substr := range substrings {
		if strings.Contains(agent, substr) {
			return true
		}
	}
	return false
}

func IsMobile(r *http.Request) bool {
	is := hasAnySubstring(r.UserAgent, mobileAgentSubstrings)
	if is {
		fmt.Println("Mobile browser detected! User-Agent:", r.UserAgent)
	}
	return is
}

func IsIE(r *http.Request) bool {
	is := hasAnySubstring(r.UserAgent, ieAgentSubstrings)
	if is {
		fmt.Println("IE detected! User-Agent:", r.UserAgent)
	}
	return is
}

func IsFirefox(r *http.Request) bool {
	is := hasAnySubstring(r.UserAgent, firefoxAgentSubstrings)
	if is {
		fmt.Println("Firefox detected! User-Agent:", r.UserAgent)
	}
	return is
}
