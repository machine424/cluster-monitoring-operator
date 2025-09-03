package monitoring_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMonitoring(t *testing.T) {
	RegisterFailHandler(Fail)
	suiteConfig, reporterConfig := GinkgoConfiguration()
	// Hardcoded until we find an easy way to pass it via "go test"
	suiteConfig.Timeout = 4 * time.Hour
	// TODO: for debugging, remove
	// reporterConfig.NoColor = true
	// TODO: for debugging, remove
	suiteConfig.FlakeAttempts = 2
	reporterConfig.VeryVerbose = true
	RunSpecs(t, "Monitoring Suite", suiteConfig, reporterConfig)
}
