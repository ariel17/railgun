package tests

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateAndStartTestServer receives an (ideally) already configured gin-gonic
// engine to use in a custom test server.
func CreateAndStartTestServer(r *gin.Engine) *http.Server {
	server := http.Server{
		Addr: ":8080",
		Handler: r,
	}
	go server.ListenAndServe()
	return &server
}

// StopTestServer creates a timeouted context to shutdown the server if it takes
// to long to stop.
func StopTestServer(server *http.Server) {
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	server.Shutdown(c)
	cancel()
}