# Chart Uploader

This application is designed to upload a Helm chart or charts to a Helm
repository and re-index the repository

## Usage

```
Upload a chart to an S3 helm repository.
  This command relies on your local AWS authentication configuration. See:
  https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html

  The configured credentials must have read/write access to the S3 bucket
  hosting the Helm repository

Usage:
  chart-uploader s3 [flags]

Flags:
      --bucket string            Helm repo s3 bucket
      --chartdir string          Local path to the directory containing chart(s) to upload (Defaults to cwd)
  -h, --help                     help for s3
      --indexpath string         Path to index.yaml in the remote repository (Defaults to /index.yaml)
      --region string            S3 bucket region
      --remotechartpath string   Remote path to upload chart(s) (Defaults to /)
      --repo string              The URL of the remote repository
```

## Example
```
chart-uploader s3 \
  --repo http://my-repo.example.com \
  --bucket my-repo.example.com \
  --region us-west-1
```

## Docker Example
```
docker run --rm \
    -v "$(pwd)":/charts \
    -e AWS_SECRET_ACCESS_KEY \
    -e AWS_ACCESS_KEY_ID \
  logicmonitor/chart-uploader \
    s3 \
      --repo http://my-repo.example.com \
      --bucket my-repo.example.com \
      --region us-west-1
      --chartdir /charts
```
