package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	assert_ := assert.New(t)

	server := httptest.NewServer(allowOnlyGet(validateQueryParam(sayHello)))

	// request dengan http method yang tidak diizinkan
	res, err := http.Post(server.URL+"/hello?name=Dika", "plain/text", nil)
	assert_.Nil(err)
	assert_.Equal(http.StatusMethodNotAllowed, res.StatusCode)

	// request dengan query parameter yang tidak valid
	res, err = http.Get(server.URL + "/hello?foo=Dika")
	assert_.Nil(err)
	assert_.Equal(http.StatusBadRequest, res.StatusCode)

	res, err = http.Get(server.URL + "/hello?name=Dika")
	assert_.Nil(err)
	assert_.Equal(http.StatusOK, res.StatusCode)

	resposeBody, bodyError := io.ReadAll(res.Body)
	assert_.Nil(bodyError)
	assert_.Equal("Hello Dika", string(resposeBody))
}
