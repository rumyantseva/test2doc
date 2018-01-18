package foos

import (
	"testing"

	"github.com/adams-sarah/prettytest"
	"github.com/gorilla/mux"
	"github.com/rumyantseva/test2doc/test"
)

var router *mux.Router
var server *test.Server

type mainSuite struct {
	prettytest.Suite
}

func TestRunner(t *testing.T) {
	var err error

	router = mux.NewRouter()
	AddRoutes(router)
	router.KeepContext = true

	test.RegisterURLVarExtractor(mux.Vars)

	server, err = test.NewServer(router, "test")
	if err != nil {
		panic(err.Error())
	}
	defer server.Finish()

	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(mainSuite),
	)
}
