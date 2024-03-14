package object

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"fmt"
	"os"
)

func DownloadObject(bucket string, item string) error {
	file, err := os.Create(item)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", item, err)
	}

	defer file.Close()
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, _ := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("a6ec1e1e-b806-4195-a636-5e1d16f7a2d1", "fa1d1750c62689380a54cab6d21ec7702ec9e768292236a0fc521970fde1dfec", ""),
		Region:      aws.String("default"),
		Endpoint:    aws.String("s3.ir-thr-at1.arvanstorage.ir"),
	})

	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		exitErrorf("Unable to download item %q, %v", item, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	return err

}
