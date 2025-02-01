package handlers

import (
	"fmt"
	"net/url"
	"strings"
)

type PathVars map[string]string
type UrlQueries map[string]string

// RouteHandler function signature
type RouteHandler func(vars PathVars, queries UrlQueries)

// Router struct
type Router struct {
	routes map[string]RouteHandler
}

// NewRouter initializes a new router
func NewRouter() *Router {
	return &Router{routes: make(map[string]RouteHandler)}
}

// Handle registers a route
func (r *Router) Handle(pattern string, handler RouteHandler) {
	r.routes[pattern] = handler
}

// Parse extracts URL parameters and calls the correct handler
func (r *Router) Parse(input string) {
	parsedURL, _ := url.Parse(input)
	pathParts := strings.Split(parsedURL.Path, "/")[1:] // Remove leading slash
	queryParams := parseQueryParams(parsedURL.RawQuery)

	for pattern, handler := range r.routes {
		patternParts := strings.Split(pattern, "/")[1:] // Remove leading slash

		if len(patternParts) != len(pathParts) {
			continue // Skip if the pattern does not match length
		}

		params := map[string]string{}
		match := true

		for i, part := range patternParts {
			if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
				paramName := part[1 : len(part)-1]
				params[paramName] = pathParts[i]
			} else if part != pathParts[i] {
				match = false
				break
			}
		}

		if match {
			handler(params, queryParams)
			return
		}
	}

	fmt.Println("No matching route found!")
}

// parseQueryParams extracts query params
func parseQueryParams(query string) map[string]string {
	params := map[string]string{}
	parsedQuery, _ := url.ParseQuery(query)
	for key, values := range parsedQuery {
		params[key] = values[0]
	}
	return params
}
