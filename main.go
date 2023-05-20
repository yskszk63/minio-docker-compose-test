package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	cx := context.Background()
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(cx)
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg, s3.WithEndpointResolver(s3.EndpointResolverFromURL("http://s3:9000")))

	o, err := client.ListObjectsV2(cx, &s3.ListObjectsV2Input{
		Bucket: aws.String("test"),
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, obj := range o.Contents {
		fmt.Printf("%s\n", *obj.Key)

		r, err := client.GetObject(cx, &s3.GetObjectInput{
			Bucket: aws.String("test"),
			Key: obj.Key,
		})
		if err != nil {
			log.Fatal(err)
		}

		func() {
			defer r.Body.Close()
			io.Copy(os.Stdout, r.Body)
		}()
	}
}
