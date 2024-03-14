package api

import (
	"context"
	"log"
	"os"
	"strconv"

	db "github.com/mahdikarami0111/cloud-computing-1/db/sqlc"
	"github.com/mahdikarami0111/cloud-computing-1/object"
	amqp "github.com/rabbitmq/amqp091-go"
)

func RunServiceTwo(url string, queries *db.Queries) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Q",   // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	var forever chan struct{}

	go func() {
		for d := range msgs {
			songId := songRecogniztion(string(d.Body), queries)
			log.Println("songID is: " + songId)
		}
	}()
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func songRecogniztion(id string, queries *db.Queries) string {
	bucket := "cc1-hw1-mk"
	filename := id + ".mp3"
	intId, _ := strconv.ParseInt(id, 10, 64)
	err := object.DownloadObject(bucket, filename)
	errArg := db.UpdateStatusParams{
		ID:     intId,
		Status: "failure",
	}
	if err != nil {
		log.Fatal(err)
		_, _ = queries.UpdateStatus(context.Background(), errArg)
	}

	songTitle := shazamApi(filename).Track.Title
	log.Println("song title is: " + songTitle)
	songId, err := spotifySearch(songTitle)
	if err != nil {
		log.Fatal(err)
		_, _ = queries.UpdateStatus(context.Background(), errArg)
	}
	arg := db.UpdateSongIDParams{
		ID:     intId,
		Songid: songId,
	}
	req, err := queries.UpdateSongID(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
		_, _ = queries.UpdateStatus(context.Background(), errArg)
	}
	arg2 := db.UpdateStatusParams{
		ID:     intId,
		Status: "ready",
	}
	req, err = queries.UpdateStatus(context.Background(), arg2)
	if err != nil {
		log.Fatal(err)
		_, _ = queries.UpdateStatus(context.Background(), errArg)
	}
	err = os.Remove(filename)
	if err != nil {
		log.Fatal(err)
		_, _ = queries.UpdateStatus(context.Background(), errArg)
	}
	return req.Songid

}
