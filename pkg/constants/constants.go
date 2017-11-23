package constants

const (
	// LocalIndexFilename is the name of the temp local index file to operate on
	LocalIndexFilename = "index.yaml"
)

const (
	// RepoTypeS3 is the config value for specifying an S3 repo
	RepoTypeS3 = "s3"
)

const (
	// ChartExtension is the default chart file extension
	ChartExtension = ".tgz"
)

const (
	// DefaultChartDir is the default path containting charts to upload
	DefaultChartDir = "/charts"
	// DefaultIndexPath is the default path for remote index files
	DefaultIndexPath = "/index.yaml"
	// DefaultRmtChartPath  is the default path for uploading charts
	DefaultRmtChartPath = "/"
)
