// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package main

type Group struct {
	ID string `json:"id"`
}

type User struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Groups []Group `json:"groups"`
}
