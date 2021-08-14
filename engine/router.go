package engine

import "net/http"

// Router struct describes next server controller for a request
// Whith options to hendle request thrue Loger or Tester midleware interfaces
type Router struct {
	Tester  bool
	Logger  bool
	Handler http.Handler
}
