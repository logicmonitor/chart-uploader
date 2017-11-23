package config

// Config for uploader
type Config struct {
	ChartPath    string
	IndexPath    string
	RepoType     string
	RepoURL      string
	RmtChartPath string
	S3           S3Config
}

// S3Config for s3 uploader
type S3Config struct {
	Bucket string
	Region string
}
