package httperror_test

import (
	"os"
	"testing"

	"github.com/IamNator/httperror"
)

func TestReply(t *testing.T) {
	testEr := httperror.New(nil)
	testEr.WriteToWriter(os.Stdout)
}
