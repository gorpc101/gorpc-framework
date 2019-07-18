package main

import (
	"git.code.oa.com/go-neat/core/nserver/default_nserver"
	_ "{{.ProtoSpec.Handler}}"
	{{- if .HttpOn}}
	_ "git.code.oa.com/go-neat/core/proto/http/dft_httpsvr"
	{{- end}}
	"net/http"

	_ "{{.ServerName}}/exec"
    _ "net/http/pprof"
)

func main() {

	// goto `http://<server-ip>:6060/debug/pprof` to find the bottleneck,
	// or, run `go tool pprof http://<server-ip>:6060/debug/pprof/profile` to get a interactive console.
	// this can be deleted before deploying to productive environment.
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	default_nserver.Serve()
}