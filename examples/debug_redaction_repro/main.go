// Package main reproduces the debug-log API key leak from the bounty report.
// Run on main vs PLT-2238 to compare behavior:
//
//	go run ./examples/debug_redaction_repro/
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/go-openapi/runtime"
	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	sdktransport "github.com/spectrocloud/palette-sdk-go/api/apiutil/transport"
)

const canaryAPIKey = "canary-secret-api-key-do-not-log-12345"

func main() {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Set(runtime.HeaderContentType, runtime.JSONMime)
		rw.Header().Set("Set-Cookie", "session=super-secret-cookie")
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	host := strings.TrimPrefix(server.URL, "http://")

	oldStderr := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stderr = w

	done := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		done <- buf.String()
	}()

	rt := sdktransport.New(host, "/", []string{"http"})
	rt.SetDebug(true)
	rt.DefaultAuthentication = openapiclient.APIKeyAuth("ApiKey", "header", canaryAPIKey)
	rt.AddSensitiveValue(canaryAPIKey)

	_, err = rt.Submit(&runtime.ClientOperation{
		Method:      http.MethodGet,
		PathPattern: "/",
		Params: runtime.ClientRequestWriterFunc(func(_ runtime.ClientRequest, _ strfmt.Registry) error {
			return nil
		}),
		Reader: runtime.ClientResponseReaderFunc(func(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
			var res map[string]string
			if err := consumer.Consume(response.Body(), &res); err != nil {
				return nil, err
			}
			return res, nil
		}),
	})

	_ = w.Close()
	os.Stderr = oldStderr
	logOutput := <-done

	if err != nil {
		fmt.Printf("request failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("=== Debug log output (stderr) ===")
	fmt.Println(logOutput)
	fmt.Println("=== Verdict ===")

	leakedKey := strings.Contains(logOutput, canaryAPIKey)
	leakedHeader := strings.Contains(strings.ToLower(logOutput), "apikey:")
	leakedCookie := strings.Contains(logOutput, "super-secret-cookie")

	if leakedKey {
		fmt.Println("FAIL: API key value appears in debug logs (VULNERABLE)")
	} else {
		fmt.Println("PASS: API key value not in debug logs")
	}
	if leakedHeader {
		fmt.Println("FAIL: ApiKey header appears in debug logs (VULNERABLE)")
	} else {
		fmt.Println("PASS: ApiKey header not in debug logs")
	}
	if leakedCookie {
		fmt.Println("FAIL: Set-Cookie value appears in debug logs (VULNERABLE)")
	} else {
		fmt.Println("PASS: Set-Cookie not in debug logs")
	}

	if leakedKey || leakedHeader || leakedCookie {
		os.Exit(1)
	}
}
