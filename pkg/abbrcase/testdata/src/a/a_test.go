package a

import "testing"

var testHTTPValue = 1 // want `identifier "testHTTPValue" should use camel-case abbreviations: "testHttpValue"`

func TestAPI(t *testing.T) { // want `identifier "TestAPI" should use camel-case abbreviations: "TestApi"`
	t.Helper()
}
