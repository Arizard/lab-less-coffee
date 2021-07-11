package httplistener

import (
	"encoding/json"
	"fmt"
	"github.com/arizard/lab-less-coffee/pkg/services"
	"net/http"
	"time"
)

var StatusCodeByErrorCode = map[services.ListenerErrorCode]int{
	services.ListenerErrorNone:       200,
	services.ListenerErrorInternal:   500,
	services.ListenerErrorBadRequest: 400,
	services.ListenerErrorBadAuth:    401,
	services.ListenerErrorNotFound:   404,
}

func ConvertErrorCodeToStatusCode(ec services.ListenerErrorCode) int {
	if status, ok := StatusCodeByErrorCode[ec]; ok {
		return status
	} else {
		return 500
	}
}

func HTTPHandleResponseError(response services.Response, w http.ResponseWriter, r *http.Request) (ok bool) {
	if response.Error != nil {
		w.WriteHeader(ConvertErrorCodeToStatusCode(response.ErrorCode))
		data := map[string]string{
			"error": response.Error.Error(),
		}
		raw, _ := json.Marshal(data)
		w.Write(raw)
		return false
	}

	return true
}

// HTTPListener is a listener that implements the http package.
type HTTPListener struct {
	Mux *http.ServeMux
	Server *http.Server
	Middleware []MiddlewareFunc
}

type MiddlewareFunc func(inner http.HandlerFunc) http.HandlerFunc

func DefaultLogMiddleware(inner http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("%s %s from remote %s\n", req.Method, req.URL, req.RemoteAddr)
		inner(w, req)
	}
}

func (l *HTTPListener) Listen() {
	err := l.Server.ListenAndServe()
	fmt.Println(err)
}

func (l *HTTPListener) WithMiddleware(inner http.HandlerFunc) (stack http.HandlerFunc) {
	stack = inner
	for _, middleware := range l.Middleware {
		stack = middleware(stack)
	}

	return stack
}

type getResponse func(w http.ResponseWriter, req *http.Request) services.Response

func createHandler(get getResponse, allowedMethods []string) http.HandlerFunc {
	allowedMethodsMap := map[string]bool{}
	for _, method := range allowedMethods {
		allowedMethodsMap[method] = true
	}
	return func(w http.ResponseWriter, req *http.Request) {
		if _, ok := allowedMethodsMap[req.Method]; !ok {
			w.WriteHeader(405)
			return
		}
		response := get(w, req)
		if !HTTPHandleResponseError(response, w, req) {
			return
		}
		raw, err := json.Marshal(response.Data)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		if _, err = w.Write(raw); err != nil {
			fmt.Println(err)
		}
		return
	}
}

func (l *HTTPListener) HandleFunc(pattern string, handle func(w http.ResponseWriter, r *http.Request) services.Response, allowedMethods []string) {
	l.Mux.HandleFunc(pattern, l.WithMiddleware(createHandler(handle, allowedMethods)))
}

func (l *HTTPListener) AddMiddleware(middleware MiddlewareFunc) {
	l.Middleware = append(l.Middleware, middleware)
}

func NewHTTPListener() *HTTPListener {
	mux := &http.ServeMux{}
	return &HTTPListener{
		Mux: mux,
		Server: &http.Server{
			Addr: ":8080",
			Handler: mux,
			ReadTimeout: 10 * time.Second,
			WriteTimeout: 10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}
