package main

import (
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
)

type Payload struct {
	Text string `json:"text"`
}

func main() {
	messagePtr := flag.String("message", "", "Message to send")
	snsPtr := flag.String("sns", "", "SNS ARN to push to")

	flag.Parse()

	if len(*messagePtr) > 0 && len(*snsPtr) > 0 {
		svc := sns.New(session.New())

		input := &sns.PublishInput{
			Message:  aws.String(*messagePtr),
			TopicArn: aws.String(*snsPtr),
		}

		_, err := svc.Publish(input)

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					log.Fatal(aerr)
				}
			} else {
				log.Fatal(err)
			}

			return
		}
	} else {
		log.Fatal("missing parameters")
	}
}
