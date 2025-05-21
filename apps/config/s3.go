package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

// S3Config holds the configuration for S3
type S3Config struct {
	Session    *session.Session
	Service    *s3.S3
	BucketName string
}

// InitS3 initializes the S3 client and bucket name
func InitS3() (*S3Config, error) {
	accessKey := viper.GetString("S3_ACCESS_KEY")
	secretKey := viper.GetString("S3_SECRET_KEY")
	host := viper.GetString("S3_HOST")
	bucket := viper.GetString("S3_BUCKET")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"), // You can change the region if needed
		Endpoint:    aws.String(host),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	return &S3Config{
		Session:    sess,
		Service:    svc,
		BucketName: bucket,
	}, nil
}
