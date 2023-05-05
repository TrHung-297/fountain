
package s3_client

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gabriel-vasile/mimetype"
	"github.com/TrHung-297/fountain/baselib/g_log"
)

// s3FilePath is key
func (s *S3Client) UploadFile(file []byte, s3FilePath string, bucketNames ...string) (linkCDN string, err error) {
	bucketName := config.BucketName
	for _, bk := range bucketNames {
		if bk != "" {
			bucketName = bk
			break
		}
	}

	body := bytes.NewReader(file)
	mime := mimetype.Detect(file)

	params := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(s3FilePath),
		Body:        body,
		ContentType: aws.String(mime.Extension()),
		//ACL:         aws.String("public-read"),
	}

	_, err = s.Client.PutObject(params)
	if err != nil {
		g_log.WithError(err).Errorf("S3Client::UploadFile - PutObject Error: %+v", err)
	}
	linkCDN = fmt.Sprintf("%v/%v", s.config.CDN, s3FilePath)

	return
}

// UploadImage func;
// s3FilePath is key
func (s *S3Client) UploadImageBase64(imageBase64 string, s3FilePath string, bucketNames ...string) (linkCDN string, err error) {
	bucketName := config.BucketName
	for _, bk := range bucketNames {
		if bk != "" {
			bucketName = bk
			break
		}
	}

	imageBase64 = strings.TrimPrefix(imageBase64, "data:")
	imageBase64Arr := strings.Split(imageBase64, ";base64,")

	var (
		dataImage []byte
		imageType string
	)

	defaultEncoding := "base64"

	if len(imageBase64Arr) > 1 {
		imageType = imageBase64Arr[0]
		dataImage = []byte(imageBase64Arr[1])
	}

	body := bytes.NewReader(dataImage)

	params := &s3.PutObjectInput{
		Bucket:          aws.String(bucketName),
		Key:             aws.String(s3FilePath),
		Body:            body,
		ContentEncoding: &defaultEncoding,
		ContentType:     &imageType,
		ACL:             aws.String("public-read"),
	}

	_, err = s.Client.PutObject(params)
	if err != nil {
		g_log.WithError(err).Errorf("S3Client::UploadFile - PutObject Error: %+v", err)
	}
	linkCDN = fmt.Sprintf("%v/%v", s.config.CDN, s3FilePath)

	return
}
