// Package authdemo a demo plugin.
package auth_demo

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

// Config the plugin configuration.
type Config struct {
	AuthTarget    string `json:"auth_target,omitempty"`
	AuthCookie    string `json:"auth_cookie,omitempty"`
	ForwardHeader string `json:"forward_header,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		AuthTarget:    "http://host.docker.internal:8080/auth",
		AuthCookie:    "login",
		ForwardHeader: "X-Auth-Token",
	}
}

// Demo a Demo plugin.
type Demo struct {
	client *http.Client
	config *Config
	name   string
	next   http.Handler
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Demo{
		client: &http.Client{},
		config: config,
		name:   name,
		next:   next,
	}, nil
}

func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := a.addForwardHeader(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	a.next.ServeHTTP(rw, req)
}

func (a *Demo) addForwardHeader(forwardRequest *http.Request) error {
	authCookie, err := forwardRequest.Cookie(a.config.AuthCookie)
	if err != nil {
		return err
	}

	authRequest, err := http.NewRequest("GET", a.config.AuthTarget, nil)
	if err != nil {
		return err
	}

	authRequest.AddCookie(authCookie)
	authResponse, err := a.client.Do(authRequest)
	if err != nil {
		return err
	}

	forwardRequest.Header.Add(a.config.ForwardHeader, authResponse.Status)

	return nil
}
