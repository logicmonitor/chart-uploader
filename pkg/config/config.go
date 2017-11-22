package config

// Config for uploader
type Config struct {
	RepoType  string
	RepoURL   string
	IndexPath string
	ChartPath string
	S3        S3Config
}

// S3Config for s3 uploader
type S3Config struct {
	Bucket string
	Region string
}
