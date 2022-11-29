//go:build !windows

package core

import (
	"net/http"
	"time"

	"github.com/fvbock/endless"
)

// func initServer(address string, router *gin.Engine) server {
func initServer(address string, router http.Handler) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
