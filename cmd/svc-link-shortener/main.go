package main

import (
	"github.com/arizard/lab-less-coffee/cmd/svc-link-shortener/application"
)


func main() {
	var svc *application.LinkShortener

	svc = application.NewLinkShortener(
		application.NewLinkShortenerHTTPListener,
		func() application.LinkRepository {
			return application.NewLocalLinkRepository()
		},
	)

	svc.Run()
}


