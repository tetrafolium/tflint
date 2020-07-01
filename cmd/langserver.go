package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sourcegraph/jsonrpc2"
	"github.com/tetrafolium/tflint/langserver"
	"github.com/tetrafolium/tflint/tflint"
)

func (cli *CLI) startLanguageServer(configPath string, cliConfig *tflint.Config) int {
	log.Println("Starting language server...")

	handler, plugin, err := langserver.NewHandler(configPath, cliConfig)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to start language server: %s", err))
		return ExitCodeError
	}
	if plugin != nil {
		defer plugin.Clean()
	}

	var connOpt []jsonrpc2.ConnOpt
	<-jsonrpc2.NewConn(
		context.Background(),
		jsonrpc2.NewBufferedStream(langserver.NewConn(os.Stdin, os.Stdout), jsonrpc2.VSCodeObjectCodec{}),
		handler,
		connOpt...,
	).DisconnectNotify()
	log.Println("Shutting down...")

	return ExitCodeOK
}
