package main

import (
	"fmt"
	"os"

	openapiRoute "github.com/ruriazz/gopen-api/openapi"
	"github.com/ruriazz/gopen-api/package/manager"
)

func main() {
	mgr, err := manager.CreateManager()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	openapiRoute.InitRoute(mgr)
	mgr.Server.HttpServer.ListenAndServe()
}
