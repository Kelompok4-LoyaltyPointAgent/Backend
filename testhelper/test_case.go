package testhelper

import "github.com/golang-jwt/jwt"

type HTTPTestCase struct {
	Name         string
	Request      Request
	Body         any
	Token        *jwt.Token
	ExpectedCode int
	ExpectedFunc func()
}

type PathParam struct {
	Names  []string
	Values []string
}

type Request struct {
	Method      string
	ContentType string
	Path        string
	PathParam   *PathParam
}
