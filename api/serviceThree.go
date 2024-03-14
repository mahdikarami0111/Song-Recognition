package api

import (
	"context"
	"log"
	"strings"
	"time"

	db "github.com/mahdikarami0111/cloud-computing-1/db/sqlc"
)

const (
	apiKey = "45f3ba214ce75a269d6c9ad10be25618-b02bcf9f-781f5621"
	domain = "sandbox7de32744eb5e46c9921f51799d0aca30.mailgun.org"
)

func RunServiceThree(queries *db.Queries) {
	for {
		time.Sleep(5 * time.Second)
		reqs, err := queries.GetByStatus(context.Background(), "ready")
		if err != nil {
			log.Fatal(err)
		}

		for _, req := range reqs {
			errArg := db.UpdateStatusParams{
				ID:     req.ID,
				Status: "failure",
			}
			recom, err := spotifyRecommend(req.Songid)
			if err != nil {
				log.Fatal(err)
				_, _ = queries.UpdateStatus(context.Background(), errArg)
			}
			log.Println(recom)
			arg := db.UpdateStatusParams{
				ID:     req.ID,
				Status: "done",
			}
			_, err = queries.UpdateStatus(context.Background(), arg)
			if err != nil {
				log.Fatal(err)
				_, _ = queries.UpdateStatus(context.Background(), errArg)
			}
			_, err = SendEmail(domain, apiKey, req.Email, strings.Join(recom, ", "))
			if err != nil {
				log.Fatal(err)
				_, _ = queries.UpdateStatus(context.Background(), errArg)
			}
		}
	}
}
