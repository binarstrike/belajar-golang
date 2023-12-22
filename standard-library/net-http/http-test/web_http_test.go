package testhttp

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	_ "os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var helloHandler http.HandlerFunc = func(writer http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")

	fmt.Fprintf(writer, "Hello %s", p[len(p)-1])
	// writer.Write([]byte(fmt.Sprint("Hello ", p[len(p)-1])))
}

func TestHelloHandler(t *testing.T) {
	assert_ := assert.New(t)
	request := httptest.NewRequest(http.MethodGet, "/hello/Otong", nil)
	rec := httptest.NewRecorder()

	helloHandler(rec, request)

	response := rec.Result()
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)

	t.Log(string(data))

	assert_.Nil(err)
	assert_.Equal("Hello Otong", string(data))
}

func TestClientRequestHello(t *testing.T) {
	assert_ := assert.New(t)
	expected := "Hello Dika"

	server := httptest.NewServer(helloHandler)

	defer server.Close()

	client := NewClient(server.URL)
	responseData, err := client.RequestHello("Dika")

	t.Log(responseData)

	assert_.Nil(err)
	assert_.Equal(expected, responseData)
}
