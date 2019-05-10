package fictionbase

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/fictionbase/fictionbase"
	"github.com/spf13/viper"
)

// Cw Service Client Operator
type Cw struct {
	sess *session.Session
	svc  *cloudwatch.CloudWatch
}

// NewCw Create New cloudwatch struct
func NewCw() *Cw {
	SetViperConfig()
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(viper.GetString("region"))},
	))
	svc := cloudwatch.New(sess)
	c := &Cw{
		sess: sess,
		svc:  svc,
	}
	return c
}

// SendCloudWatch send data to CloudWatch
func (Cw Cw) SendCloudWatch(input *cloudwatch.PutMetricDataInput) error {
	_, err := Cw.svc.PutMetricData(input)
	if err != nil {
		return err
	}
	return nil
}

// GetCloudWatch get data to CloudWatch
func (Cw Cw) GetCloudWatch(nameSpace string, metricName string, dimensionsName string) *cloudwatch.ListMetricsOutput {
	result, err := Cw.svc.ListMetrics(&cloudwatch.ListMetricsInput{
		Namespace:  aws.String(nameSpace),
		MetricName: aws.String(metricName),
		Dimensions: []*cloudwatch.DimensionFilter{
			&cloudwatch.DimensionFilter{
				Name: aws.String(dimensionsName),
			},
		},
	})
	if err != nil {
		fictionbase.Logger.Error(err.Error())
	}
	return result
}
