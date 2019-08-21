package kit_test

import (
	"github.com/go-kit/kit/log"

	kitadapter "logur.dev/adapter/kit"
)

func ExampleNew() {
	logger := kitadapter.New(log.NewNopLogger())

	// Output:
	_ = logger
}

// If logger is nil, a default instance is created.
func ExampleNew_default() {
	logger := kitadapter.New(nil)

	// Output:
	_ = logger
}
