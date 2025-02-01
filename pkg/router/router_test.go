package router

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

// Helper type to capture printed output.
type outputCapture struct {
	oldStdout *os.File
	reader    *os.File
	writer    *os.File
	out       chan string
}

func newOutputCapture() *outputCapture {
	r, w, _ := os.Pipe()
	return &outputCapture{
		oldStdout: os.Stdout,
		reader:    r,
		writer:    w,
		out:       make(chan string),
	}
}

func (c *outputCapture) start() {
	os.Stdout = c.writer
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, c.reader)
		c.out <- buf.String()
	}()
}

func (c *outputCapture) stop() string {
	_ = c.writer.Close()
	os.Stdout = c.oldStdout
	return <-c.out
}

func TestRouter_MatchedRoutes(t *testing.T) {
	// Define a list of test cases.
	tests := []struct {
		name         string
		routePattern string
		input        string
		expectedVars PathVars
		expectedQry  UrlQueries
	}{
		{
			name:         "About route with action get",
			routePattern: "/about/{action}",
			input:        "/about/get",
			expectedVars: PathVars{"action": "get"},
			expectedQry:  UrlQueries{},
		},
		{
			name:         "Menu route with no variables",
			routePattern: "/menu",
			input:        "/menu",
			expectedVars: PathVars{},
			expectedQry:  UrlQueries{},
		},
		{
			name:         "Support route with no variables",
			routePattern: "/support",
			input:        "/support",
			expectedVars: PathVars{},
			expectedQry:  UrlQueries{},
		},
		{
			name:         "Help route with update action and query param",
			routePattern: "/help/{action}",
			input:        "/help/update?foo=bar",
			expectedVars: PathVars{"action": "update"},
			expectedQry:  UrlQueries{"foo": "bar"},
		},
		{
			name:         "Search route with no variables",
			routePattern: "/search",
			input:        "/search",
			expectedVars: PathVars{},
			expectedQry:  UrlQueries{},
		},
		{
			name:         "Search route with query variable",
			routePattern: "/search/{query}",
			input:        "/search/title",
			expectedVars: PathVars{"query": "title"},
			expectedQry:  UrlQueries{},
		},
		{
			name:         "Manage route with product query",
			routePattern: "/manage/{query}",
			input:        "/manage/product",
			expectedVars: PathVars{"query": "product"},
			expectedQry:  UrlQueries{},
		},
		{
			name:         "Product route with action and product id with query",
			routePattern: "/product/{action}/{product}",
			input:        "/product/get/42?x=1",
			expectedVars: PathVars{"action": "get", "product": "42"},
			expectedQry:  UrlQueries{"x": "1"},
		},
		{
			name:         "FAQ route with no variables",
			routePattern: "/faq",
			input:        "/faq",
			expectedVars: PathVars{},
			expectedQry:  UrlQueries{},
		},
		{
			name:         "FAQ add route with no variables",
			routePattern: "/faq/add",
			input:        "/faq/add",
			expectedVars: PathVars{},
			expectedQry:  UrlQueries{},
		},
		{
			name:         "FAQ menu route with update action",
			routePattern: "/faq/menu/{action}",
			input:        "/faq/menu/update",
			expectedVars: PathVars{"action": "update"},
			expectedQry:  UrlQueries{},
		},
		{
			name:         "FAQ route with action and question id",
			routePattern: "/faq/{action}/{question}",
			input:        "/faq/get/42",
			expectedVars: PathVars{"action": "get", "question": "42"},
			expectedQry:  UrlQueries{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			router := NewRouter()
			called := false

			// Register the route with a handler that checks the captured variables.
			router.Handle(tc.routePattern, func(vars PathVars, queries UrlQueries) {
				called = true
				if !reflect.DeepEqual(vars, tc.expectedVars) {
					t.Errorf("expected path vars %v, got %v", tc.expectedVars, vars)
				}
				if !reflect.DeepEqual(queries, tc.expectedQry) {
					t.Errorf("expected query params %v, got %v", tc.expectedQry, queries)
				}
			})

			router.Parse(tc.input)
			if !called {
				t.Errorf("expected handler to be called for route %q with input %q", tc.routePattern, tc.input)
			}
		})
	}
}

func TestRouter_NoMatch(t *testing.T) {
	router := NewRouter()

	// Register a known route.
	router.Handle("/known/route", func(vars PathVars, queries UrlQueries) {
		t.Error("handler should not have been called")
	})

	// Capture the output to verify the "No matching route found!" message.
	capture := newOutputCapture()
	capture.start()

	// Use an input that does not match any registered route.
	router.Parse("/unknown/route")

	output := capture.stop()

	if !strings.Contains(output, "No matching route found!") {
		t.Errorf("expected output to contain \"No matching route found!\", got %q", output)
	}
}

func TestRouter_TrailingSlashHandling(t *testing.T) {
	// Test that a trailing slash changes the segment count and might not match.
	router := NewRouter()
	called := false
	router.Handle("/menu", func(vars PathVars, queries UrlQueries) {
		called = true
	})

	// When input has a trailing slash, the segments may be different.
	router.Parse("/menu/") // This results in ["menu", ""] for pathParts.

	if called {
		t.Error("expected handler NOT to be called for /menu/ because of segment mismatch")
	}
}

func TestRouter_QueryParameterParsing(t *testing.T) {
	// Test that multiple query parameters are correctly parsed.
	router := NewRouter()
	called := false
	router.Handle("/test", func(vars PathVars, queries UrlQueries) {
		called = true
		expected := UrlQueries{"a": "1", "b": "2"}
		if !reflect.DeepEqual(queries, expected) {
			t.Errorf("expected queries %v, got %v", expected, queries)
		}
	})

	router.Parse("/test?a=1&b=2")
	if !called {
		t.Error("expected handler to be called for /test")
	}
}
