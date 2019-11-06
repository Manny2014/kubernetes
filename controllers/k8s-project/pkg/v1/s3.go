package v1

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io/ioutil"
	"os"
)

type S3Source struct {
	Bucket string `json:"bucket"`
	Prefix string `json:"prefix"`
	Object string `json:"object"`
}

func NewS3Source(config SourceConfig) (Source, error) {
	src := &S3Source{
		Bucket: config.BasePath,
		Prefix: config.Prefix,
		Object: config.Object,
	}

	return src, nil
}

func (s *S3Source) String() string {
	return fmt.Sprintf("%s/%s/%s", s.Bucket, s.Prefix, s.Object)
}

func (s *S3Source) GetData() (map[string]string, error) {

	sess, _ := session.NewSessionWithOptions(session.Options{
		// Provide SDK Config options, such as Region.
		Config: aws.Config{
			Region: aws.String("us-east-1"), // TODO: Don't Hard code
		},
	})

	downloader := s3manager.NewDownloader(sess)
	downloader.Concurrency = 1

	f, _ := os.Create(s.Object)
	defer func() {
		_ = os.Remove(s.Object)
	}()

	_, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", s.Prefix, s.Object)),
	})

	if err != nil {
		return nil, err
	}

	dat, err := ioutil.ReadFile(s.Object)

	var data map[string]string

	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal([]byte(dat), &data)

	data[s.Object] = string(dat)

	return data, nil
}
