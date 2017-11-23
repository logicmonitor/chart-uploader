package uploader

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/logicmonitor/chart-uploader/pkg/config"
	"github.com/logicmonitor/chart-uploader/pkg/constants"
	log "github.com/sirupsen/logrus"
)

// UploadS3 the configured chart to the specified helm repo and re-index the repo
func UploadS3(upldConfig *config.Config) error {
	sess := getAwsSess(upldConfig.S3.Region)

	err := downloadS3Object(upldConfig.S3.Bucket, upldConfig.IndexPath, constants.LocalIndexFilename, sess)
	if err != nil {
		return err
	}

	err = updateAndMergeIndex(upldConfig.ChartPath, upldConfig.RepoURL, constants.LocalIndexFilename)
	if err != nil {
		return err
	}

	err = uploadS3Object(upldConfig.S3.Bucket, upldConfig.IndexPath, constants.LocalIndexFilename, sess)
	if err != nil {
		return err
	}

	log.Infof("Successfully uploaded charts in %s to %s", upldConfig.ChartPath, upldConfig.RepoURL)
	cleanup(constants.LocalIndexFilename)
	return nil
}

func updateAndMergeIndex(chartPath string, repoURL string, localIndexPath string) error {
	args := []string{
		"repo", "index", chartPath,
		"--merge", localIndexPath,
		"--url", repoURL,
	}
	res, err := shellCmd("helm", args)
	if err != nil {
		return err
	}
	if res != "" {
		log.Debugf(res)
	}
	return nil
}

func shellCmd(name string, args []string) (string, error) {
	log.Debugf("Running command %s %s", name, args)
	cmd := exec.Command(name, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	slurp, _ := ioutil.ReadAll(stdout)
	slurperr, _ := ioutil.ReadAll(stderr)

	if err := cmd.Wait(); err != nil {
		log.Warnf("%s", string(slurperr))
		return "", err
	}
	return string(slurp), nil
}

func cleanup(filename string) {
	log.Debugf("Deleting %s", filename)
	os.Remove(filename)
}
