package keyspaces_test

import (
	"os"
	"os/exec"
	"strings"
	"testing"

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
			output, err := applyCmd.CombinedOutput()
			convey.Printf("kubectl apply output: %s", output)
			convey.So(err, convey.ShouldBeNil)

			convey.Convey("Then the related keyspace should exist", func() {
				keyspaces, err := awsclient.ListKeyspaces()
				convey.So(err, convey.ShouldBeNil)

				keyspaceName := "gary_predev_test33"
				found := false
				for _, ks := range keyspaces {
					if *ks.KeyspaceName == keyspaceName {
						found = true
						break
					}
				}
				convey.So(found, convey.ShouldBeTrue)
			})
		})
	})
}
