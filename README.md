s3zipper
========

Golang tool for getting zipped list of files stored in Amazon S3.

s3zipper tries to minimize memory footprint by streaming compressed data as is being downloaded, instead of a) getting, files b) compress them, and c) send compressed data.

Installation
--
`go get github.com/gchaincl/s3zipper`

Usage:
--

```bash
AWS_ACCESS_KEY="AccessKey" AWS_SECRET_KEY="SecretKey" ./s3zipper -bucket yourBucket -port=8000
```
The default port will be 8000 if you don't specify another one.

Then you should specify a list of files you want to download separated by ",":
```bash
wget -O file http://localhost:8000\?files=foo.pdf,bar.pdf
```

Notes:
--
* This is a proof of concept, don't expect s3zipper to be an elaborated tool
* Unexisting files will be ignored
