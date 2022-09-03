package tracerr_test

import (
	"os"
	"testing"
)

var wd string

func TestMain(m *testing.M) {
	var err error
	if wd, err = os.Getwd(); err != nil {
		panic(err)
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}


