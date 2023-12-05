/*
you need to install amqp client for RabbitMQ
1) go get -t github.com/streadway/amqp/
2) install docker
3) run this to start your rabbitmq server
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.9-management

*) you may need to have go.mod file
*/

/*
USAGE
go run receive_logs_topic.go *.info # receives from anything having info in the end
go run receive_logs_topic.go anonymous.*
go run receive_logs_topic.go kernel.error

USAGE:
1) you can push to ONLY ONE TOPIC, you may modify this code to emit to multiple topics, loop on publish
2) THIS CAN'T HAVE * OR #

go run emit_logs_topic.go anonymous.info # this sends to anonymous.info
go run emit_logs_topic.go kernel.info  # this sends to kernel.info
*/
package main

import (
	"context"
	"fmt"
	"os"
	"section7/utils"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	utils.FailOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	topic := severityFrom(os.Args)
	for i := 0; ; {
		var hi string
		fmt.Print(" ( " + topic + " ) > ")
		fmt.Scan(&hi)
		body := " ( " + topic + " SAYS): > " + hi

		err = ch.PublishWithContext(ctx,
			"logs_topic", // exchange
			topic,        // routing key (THIS CAN'T HAVE * OR #)
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		utils.FailOnError(err, "Failed to publish a message")

		//log.Printf(" ( %s ) > %s ", topic, body)
		time.Sleep(400 * time.Millisecond)
		i += 1
	}
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
}
