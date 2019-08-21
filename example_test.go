package kit_test

import (
	kitadapter "logur.dev/adapter/kit"
)

func ExampleNew() {
	var l interface{}
	logger := kitadapter.New(l)

	// Output:
	_ = logger
}

// If logger is nil, a default instance is created.
func ExampleNew_default() {
	logger := kitadapter.New(nil)

	// Output:
	_ = logger
}
