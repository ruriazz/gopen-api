package main

import (
	"fmt"
	"os"

	apiRoute "github.com/ruriazz/gopen-api/api"
	"github.com/ruriazz/gopen-api/package/manager"
)

func main() {
	manager, err := manager.NewManager()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	apiRoute.NewApiRoute(manager)
	manager.Server.HttpServer.ListenAndServe()
}
