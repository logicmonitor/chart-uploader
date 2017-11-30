package constants

var (
	// Version is the Chart Uploader version and is set at build time.
	Version string
)

const (
	// LocalIndexFilename is the name of the temp local index file to operate on
	LocalIndexFilename = "index.yaml"
)

const (
	// ChartExtension is the default chart file extension
	ChartExtension = ".tgz"
)

const (
	// DefaultChartDir is the default path containtin charts to upload
	DefaultChartDir = "/charts"
	// DefaultIndexPath is the default path for remote index files
	DefaultIndexPath = "/index.yaml"
	// DefaultRmtChartPath  is the default path for uploading charts
	DefaultRmtChartPath = "/"
)
