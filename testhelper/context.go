package testhelper

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func NewContext(testCase HTTPTestCase) (echo.Context, *httptest.ResponseRecorder) {
	bodyReader := new(bytes.Reader)
	if testCase.Body != nil {
		body, err := json.Marshal(testCase.Body)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		bodyReader = bytes.NewReader(body)
	}

	req := httptest.NewRequest(testCase.Request.Method, testCase.Request.Path, bodyReader)
	req.Header.Set("Content-Type", testCase.Request.ContentType)

	rec := httptest.NewRecorder()

	ctx := echo.New().NewContext(req, rec)
	ctx.SetPath(testCase.Request.Path)
	if testCase.Request.PathParam != nil {
		ctx.SetParamNames(testCase.Request.PathParam.Names...)
		ctx.SetParamValues(testCase.Request.PathParam.Values...)
	}
	ctx.Set("token", testCase.Token)

	return ctx, rec
}
