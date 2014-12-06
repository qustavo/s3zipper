package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/s3"
	"net/http"
)

var aws_bucket *s3.Bucket

func initAwsBucket(name, region string) {
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err)
	}

	aws_bucket = s3.New(auth, aws.GetRegion(region)).Bucket(name)
}

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	files, ok := r.URL.Query()["files"]
	if !ok || len(files) < 1 {
		http.Error(w, "Specify at least one file with ?files='file1,file2,fileN'", 500)
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename="+aws_bucket.Name+".zip")
	w.Header().Add("Content-Type", "application/zip")

	zipWriter := zip.NewWriter(w)

	for _, file := range strings.Split(files[0], ",") {
		fmt.Printf("Processing '%s'\n", file)

		rdr, err := aws_bucket.GetReader(file)
		if err != nil {
			switch t := err.(type) {
			case *s3.Error:
				// skip non existing files
				if t.StatusCode == 404 {
					continue
				}
			}
			panic(err)
		}

		f, err := zipWriter.Create(aws_bucket.Name + "/" + file)
		if err != nil {
			panic(err)
		}

		io.Copy(f, rdr)
		rdr.Close()
	}

	zipWriter.Close()

	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
}

func main() {
	var port int
	bucket := flag.String("bucket", "/", "S3 Bucket")
	region := flag.String("region", "us-east-1", "S3 Region")
	flag.IntVar(&port, "port", 8000, "port to use for s3zipper")
	flag.Parse()

	initAwsBucket(*bucket, *region)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
