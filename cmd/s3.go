// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
  "github.com/logicmonitor/chart-uploader/pkg/config"
  "github.com/logicmonitor/chart-uploader/pkg/uploader"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var bucket string
var chartPath string
var indexPath string
var region string
var repoType string
var repoURL string

// s3cmd represents thes3  upload command
var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Upload a chart to an S3 helm repository",
  Long: `Upload a chart to an S3 helm repository.
  This command relies on your local AWS authentication configuration. See:
  https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html

  The configured credentials must have read/write access to the S3 bucket
  hosting the Helm repository`,

	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve the application configuration.
		upldconfig := initS3Config()
    err := uploader.UploadS3(upldconfig)
    if err != nil {
      log.Fatalf("Failed to upload chart: %v", err)
    }
	},
}

func init() {
	log.SetLevel(log.DebugLevel)
  s3Cmd.Flags().StringVar(&bucket, "bucket", "", "Helm repo s3 bucket")
  s3Cmd.Flags().StringVar(&chartPath, "chartdir", "", "Local path to the directory containing chart(s) to upload (Defaults to /charts)")
  s3Cmd.Flags().StringVar(&indexPath, "indexpath", "", "Path to index.yaml in the remote repository (Defaults to /index.yaml)")
  s3Cmd.Flags().StringVar(&region, "region", "", "S3 bucket region")
  s3Cmd.Flags().StringVar(&repoURL, "repo", "", "The URL of the remote repository")
  viper.SetDefault("chartdir", "/charts")
  viper.SetDefault("indexpath", "/index.yaml")
  RootCmd.AddCommand(s3Cmd)
}

func initS3Config() (*config.Config) {
  return &config.Config{
    ChartPath: chartPath,
    IndexPath: indexPath,
    RepoURL: repoURL,
    S3: config.S3Config{
      Bucket: bucket,
      Region: region,
    },
  }
}
