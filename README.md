s3zipper
========

Golang tool for getting zipped list of files stored in Amazon S3 

Installation
--
`go get github.com/gchaincl/s3zipper`

Usage:
--

```bash
AWS_ACCESS_KEY="AccessKey" AWS_SECRET_KEY="SecretKey" ./s3zip -bucket yourBucket
```
Server will be listening on port 8000 (change the code for a different port)

Then you should specify a list of files you want to download separated by ",":
```bash
wget -O file http://localhost:8000\?files=foo.pdf,bar.pdf
```

Notes:
--
* This is a proof of concept, don't expect s3zipper to be an elaborated tool
* Unexising files will be ignored
