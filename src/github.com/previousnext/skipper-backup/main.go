package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	tokenFreq = "%FREQUENCY%"
	tokenTime = "%TIMESTAMP%"
)

var (
	cliRegion = kingpin.Flag("region", "Region which the S3 bucket resides").Default("ap-southeast-2").String()
	cliLocal  = kingpin.Arg("local", "Local file to upload to S3").Required().String()
	cliBucket = kingpin.Arg("bucket", "S3 bucket").Required().String()
	cliRemote = kingpin.Arg("remote", "Remote path in the AWS S3 bucket").Required().String()
)

func main() {
	kingpin.CommandLine.Help = "Backup Utility"
	kingpin.Parse()

	// Load the Skipper configuration.
	key, secret, err := obtainCredentialsSkipper()
	if err != nil {
		panic(err)
	}

	// Load the file to upload.
	file, err := os.Open(*cliLocal)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	uploader := s3manager.NewUploader(session.New(&aws.Config{
		Region:      cliRegion,
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
	}))

	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   file,
		Bucket: cliBucket,
		Key:    aws.String(tokens(*cliRemote)),
		ACL:    aws.String("private"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully uploaded:", result.Location)
}

// Helper function to obtain AWS credentials configured with skipper.
func obtainCredentialsSkipper() (string, string, error) {
	key, err := config("backup.id")
	if err != nil {
		return "", "", err
	}

	secret, err := config("backup.secret")
	if err != nil {
		return "", "", err
	}

	return key, secret, err
}

func config(name string) (string, error) {
	b, err := ioutil.ReadFile(fmt.Sprintf("/etc/skpr/%s", name))
	if err != nil {
		return "", err
	}
	return string(b), err
}

// Find and replace
func tokens(r string) string {
	var (
		day     = time.Now().Day()
		weekday = time.Now().Weekday()
		current = time.Now().Local()
	)

	return tokensReplace(r, current, day, weekday)
}

func tokensReplace(r string, current time.Time, day int, weekday time.Weekday) string {
	r = strings.Replace(r, tokenTime, current.Format("2006-01-02_15-04-05"), -1)

	// Place backup into monthly directory if taken on first of the month.
	if day == 1 {
		return strings.Replace(r, tokenFreq, "monthly", -1)
	}

	// Place backup into weekly directory if taken on first Sunday.
	if weekday == time.Sunday {
		return strings.Replace(r, tokenFreq, "weekly", -1)
	}

	// Default to daily.
	return strings.Replace(r, tokenFreq, "daily", -1)
}
