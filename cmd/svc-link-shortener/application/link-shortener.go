package application

import (
	"context"
	"fmt"
	"github.com/arizard/lab-less-coffee/pkg/services"
	"net/url"
)

type LinkShortener struct {
	Listener services.Listener
	LinkRepo LinkRepository
}

func NewLinkShortener(listenerFactory func(svc *LinkShortener) services.Listener, linkRepoFactory func() LinkRepository) *LinkShortener {
	app := &LinkShortener{
		LinkRepo: linkRepoFactory(),
	}

	app.Listener = listenerFactory(app)

	return app
}

func (app *LinkShortener) Ping(ctx context.Context) services.Response {
	return services.Response{
		Error:     nil,
		ErrorCode: services.ListenerErrorNone,
		Data: map[string]string{
			"message": "pong",
		},
	}
}

func (app *LinkShortener) GetDestination(ctx context.Context, uid string) services.Response {
	link := &Link{}
	if err := app.LinkRepo.Get(uid, link); err != nil {
		return services.Response{
			Error:     err,
			ErrorCode: services.ListenerErrorInternal,
		}
	} else {
		return services.Response{
			Data: map[string]string{
				"destination": link.Destination,
			},
		}

	}
}

func (app *LinkShortener) CreateShortLink(ctx context.Context, dst string) services.Response {
	dstUrl, err := url.Parse(dst)
	if err != nil {
		return services.Response{
			Error: err,
			ErrorCode: services.ListenerErrorBadRequest,
		}
	}

	fmt.Printf("shortening url %s\n", dstUrl.String())

	newLink := Link{
		Destination: dstUrl.String(),
		DeleteCode:  "magic",
	}
	if err := app.LinkRepo.Insert(&newLink); err != nil {
		return services.Response{
			Error: err,
			ErrorCode: services.ListenerErrorInternal,
		}
	} else {
		return services.Response{
			Data: newLink,
		}
	}
}

func (app *LinkShortener) Run() {
	app.Listener.Listen()
}

type Link struct {
	Uid         string `json:"uid"`
	Destination string `json:"destination"`
	DeleteCode  string `json:"deleteCode"`
}

type LinkRepository interface {
	Get(uid string, link *Link) error
	Insert(newLink *Link) error
	Delete(uid string) error
}
