package fictionbase

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/spf13/viper"
)

// Sqs Service Client Operator
type Sqs struct {
	sess *session.Session
	svc  *sqs.SQS
}

// NewSqs Create New sqs struct
func NewSqs() *Sqs {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(viper.GetString("region"))},
	))
	svc := sqs.New(sess)
	s := &Sqs{
		sess: sess,
		svc:  svc,
	}
	return s
}

// SendFictionbaseMessage Send Message To FictionBase
func SendFictionbaseMessage(fb interface{}) (err error) {
	jsonByte, err := json.Marshal(fb)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		"POST",
		viper.GetString("sqs.url"),
		strings.NewReader(string(jsonByte)),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "Content-Type:application/json")
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

// GetFictionbaseMessage Get Message From FictionBase
func (s Sqs) GetFictionbaseMessage() ([]*sqs.Message, error) {
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(viper.GetString("sqs.sqsurl")),
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(20),
	}
	resp, err := s.svc.ReceiveMessage(params)
	if err != nil {
		return nil, err
	}
	return resp.Messages, nil
}

// DeleteFictionbaseMessage Delete Message From FictionBase
func (s Sqs) DeleteFictionbaseMessage(msg *sqs.Message) error {
	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(viper.GetString("sqs.sqsurl")),
		ReceiptHandle: aws.String(*msg.ReceiptHandle),
	}
	_, err := s.svc.DeleteMessage(params)

	if err != nil {
		return err
	}
	return nil
}
