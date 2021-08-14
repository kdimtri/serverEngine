package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/kdimtri/serverEngine/engine"
)

// API it is first http.Handler that gets all server requests
type API struct {
	Debug string
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var next *engine.Router
	var head string
	var urlPath string
	ctx := engine.SetContext(r)
	head, urlPath = engine.ShiftPath(r.URL.Path)
	if head == "api" {
		head, urlPath = engine.ShiftPath(urlPath)
		if head == "endpoint" {
			next = endpoint(r)
		} else {
			next = newError(fmt.Errorf("Api page not found"), http.StatusNotFound)
		}
	} else if a.Debug != "" && head == "debug" {
		ts := fmt.Sprintf("%x", sha256.Sum256([]byte(a.Debug)))
		head, urlPath = engine.ShiftPath(urlPath)
		if head == ts {
			head, urlPath = engine.ShiftPath(urlPath)
			if head == "pprof" {
				head, urlPath = engine.ShiftPath(urlPath)
				next = pProf(head)
			} else {
				next = newError(fmt.Errorf("Debug profile page not found"), http.StatusNotFound)
			}
		} else {
			next = newError(fmt.Errorf("Debug page not found"), http.StatusNotFound)
		}
	} else {
		next = newError(fmt.Errorf("Page not found"), http.StatusNotFound)
	}
	if next.Logger {
		next.Handler = engine.Logger(next.Handler)
	}
	next.Handler.ServeHTTP(w, r.WithContext(ctx))
}

func newError(err error, statusCode int) *engine.Router {
	return &engine.Router{
		Logger: true,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			engine.Respond(w, r, statusCode, err)
		}),
	}
}
func endpoint(r *http.Request) *engine.Router {
	m := &model.Handler{}
	if err := engine.ParseBody(r.Body, p); err != nil {
		return newError(err, http.StatusUnprocessableEntity)
	}
	return &engine.Router{
		Logger:  true,
		Handler: m,
	}
}
func pProf(name string) *engine.Router {
	var p http.Handler
	p = pprof.Handler(name)
	return &engine.Router{
		Logger:  true,
		Handler: p,
	}
}
