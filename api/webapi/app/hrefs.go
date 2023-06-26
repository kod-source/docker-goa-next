// Code generated by goagen v1.5.14, DO NOT EDIT.
//
// API "docker_goa_next": Application Resource Href Factories
//
// Command:
// $ main

package app

import (
	"fmt"
	"strings"
)

// PostsHref returns the resource href.
func PostsHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/posts/%v", paramid)
}

// RoomsHref returns the resource href.
func RoomsHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/rooms/%v", paramid)
}
