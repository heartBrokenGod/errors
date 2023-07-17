package errors_test

import (
	"testing"

	"github.com/heartBrokenGod/errors"
)

var err = errors.New("some error occured with argument %s")

func TestError(t *testing.T) {

	e := err.WithArgument("ARGUMENT")

	if !e.Is(err) {
		t.Error("e.Is(err) should return true")
	}

}
