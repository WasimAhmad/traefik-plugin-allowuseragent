// Package traefik_plugin_allowuseragent is a plugin to allow only specific User-Agents.
package traefik_plugin_allowuseragent

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "regexp"
)

// Config holds the plugin configuration.
type Config struct {
    AllowRegex []string `json:"allowRegex,omitempty"` // Regex patterns to allow
}

// CreateConfig creates and initializes the plugin configuration.
func CreateConfig() *Config {
    return &Config{AllowRegex: make([]string, 0)}
}

type allowUserAgent struct {
    name        string
    next        http.Handler
    allowRegexps []*regexp.Regexp
}

// New creates and returns a plugin instance.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
    allowRegexps := make([]*regexp.Regexp, len(config.AllowRegex))

    for i, regex := range config.AllowRegex {
        re, err := regexp.Compile(regex)
        if err != nil {
            return nil, fmt.Errorf("error compiling allow regex %q: %w", regex, err)
        }

        allowRegexps[i] = re
    }

    return &allowUserAgent{
        name:        name,
        next:        next,
        allowRegexps: allowRegexps,
    }, nil
}

func (a *allowUserAgent) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    if req != nil {
        userAgent := req.UserAgent()
        isAllowed := false

        for _, re := range a.allowRegexps {
            if re.MatchString(userAgent) {
                isAllowed = true
                break
            }
        }

        if !isAllowed {
            log.Printf("Blocked User-Agent: '%s'", userAgent)
            rw.WriteHeader(http.StatusForbidden)
            return
        }
    }

    a.next.ServeHTTP(rw, req)
}
