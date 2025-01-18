package keyspaces_test

import (
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/hfgary/crossplane-keyspaces-test/pkg/awsclient"
	"github.com/hfgary/crossplane-keyspaces-test/pkg/k8sclient"
	"github.com/smartystreets/goconvey/convey"
)

func TestKeyspaceClaimCreation(t *testing.T) {
	convey.Convey("Given a Kubernetes cluster and a Keyspace custom resource", t, func() {
		_, err := k8sclient.NewKubernetesClient()
		convey.So(err, convey.ShouldBeNil)

		keyspaceCR, err := os.ReadFile("../configs/keyspaces_claim.yaml")
		convey.So(err, convey.ShouldBeNil)

		convey.Convey("When the keyspace claim is applied", func() {
			applyCmd := exec.Command("kubectl", "apply", "-f", "-")
			applyCmd.Stdin = strings.NewReader(string(keyspaceCR))
			_, err := applyCmd.CombinedOutput()
			convey.So(err, convey.ShouldBeNil)

			convey.Convey("Then the related keyspace should exist", func() {
				var found bool
				var err error

				timeout := 600 * time.Second
				keyspaceName := "gary_predev_test33"
				startTime := time.Now()

				backoff := 1 * time.Second
				for time.Since(startTime) < timeout {
					keyspaces, err := awsclient.ListKeyspaces()
					if err == nil {
						for _, ks := range keyspaces {
							if *ks.KeyspaceName == keyspaceName {
								found = true
								break
							}
						}
						if found {
							break
						}
					}
					time.Sleep(backoff)
					backoff *= 2 // Exponential backoff
				}

				convey.So(err, convey.ShouldBeNil)
				convey.So(found, convey.ShouldBeTrue)
			})
		})
	})
}
