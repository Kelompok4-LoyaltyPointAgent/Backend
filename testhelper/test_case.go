package testhelper

type HTTPTestCase struct {
	Name         string
	Request      Request
	Body         any
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
	URL         string
	PathParam   *PathParam
}
