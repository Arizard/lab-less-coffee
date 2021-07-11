package application

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/arizard/lab-less-coffee/pkg/services"
	"github.com/arizard/lab-less-coffee/pkg/services/httplistener"
	"io"
	"io/ioutil"
	"net/http"
	"path"
)

func readJSONBody(body io.Reader) (out map[string]interface{}, err error) {
	raw, err := ioutil.ReadAll(body)
	if err != nil {
		return out, err
	}
	if err := json.Unmarshal(raw, &out); err != nil {
		return out, err
	}
	return out, nil
}

func badRequest(err error) services.Response {
	return services.Response{
		Error:     err,
		ErrorCode: services.ListenerErrorBadRequest,
	}
}

func NewLinkShortenerHTTPListener(app *LinkShortener) services.Listener {
	listener := httplistener.NewHTTPListener()
	listener.AddMiddleware(httplistener.DefaultLogMiddleware)
	listener.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) services.Response {
		//_, err := readJSONBody(req.Body)
		//if err != nil {
		//	return badRequest(errors.New("could not unmarshal json body"))
		//}
		return app.Ping(context.Background())
	},
		[]string{"GET"},
	)

	listener.HandleFunc("/link/", func(w http.ResponseWriter, req *http.Request) services.Response {
		url := req.URL.Path
		uid := path.Base(url)
		fmt.Printf("getting destination for uid: %s\n", uid)
		return app.GetDestination(context.Background(), uid)
	},
		[]string{"GET"},
	)

	listener.HandleFunc("/link", func(w http.ResponseWriter, req *http.Request) services.Response {
		body, err := readJSONBody(req.Body)
		if err != nil {
			return badRequest(errors.New("could not unmarshal json body"))
		}
		if src, ok := body["destination"]; ok {
			return app.CreateShortLink(context.Background(), src.(string))
		} else {
			return badRequest(errors.New("field `destination` is required"))
		}
	},
		[]string{"POST"},
	)

	return listener
}
