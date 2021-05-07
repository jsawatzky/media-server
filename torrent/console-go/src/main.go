/*
 * Torrent Console
 *
 * Specification for the torrent server console to provide an interface to tools
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/jsawatzky/server/torrent/console-go/openapi"
	"github.com/jsawatzky/server/torrent/console-go/service"
)

func main() {
	log.Printf("Server started")

	filebotChan := make(chan *exec.Cmd)

	go service.ProcessFilebot(filebotChan)

	cmd := exec.Command("filebot", "-script", "fn:properties", "--def", "net.filebot.xattr.store=.xattr")
	filebotChan <- cmd

	ApiService := service.NewApiService(filebotChan)
	DefaultApiController := openapi.NewDefaultApiController(ApiService)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":80", router))
}
