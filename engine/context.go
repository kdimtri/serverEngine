package engine

import (
	"context"
	"net/http"
	"time"
)

type key int

const (
	// ContextOriginalPath holds the original request URL
	ContextOriginalPath key = iota
	// ContextRequestStart holds the request start time
	ContextRequestStart
	// ContextFrontID holds sha1 from frontServer
	ContextFrontID
	// ContextPublicMountPath holds public path prefix value
	ContextPublicMountPath
	// ContextPrivateMountPath holds private path prefix value
	ContextPrivateMountPath
)

//SetContext sets context.Context from request
func SetContext(r *http.Request) *context.Context {
	ctx := context.WithValue(r.Context(), ContextRequestStart, time.Now())
	return &ctx
}
