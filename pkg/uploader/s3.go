package uploader

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	log "github.com/sirupsen/logrus"
)

func getAwsSess(region string) *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
}

func downloadS3Object(bucket string, key string, filename string, sess *session.Session) error {
	log.Debugf("Downloading S3 object %s%s to %s", bucket, key, filename)
	downloader := s3manager.NewDownloader(sess)

	// Create a file to write the S3 Object contents to.
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	// Write the contents of S3 Object to the file
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	return nil
}

func uploadS3Object(bucket string, key string, filename string, sess *session.Session) error {
	log.Debugf("Uploading file %s to S3 %s%s", filename, bucket, key)
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	// Upload the file to S3.
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   f,
	})
	return err
}
