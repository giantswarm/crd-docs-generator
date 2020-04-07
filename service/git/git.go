package git

import (
	"fmt"
	"os/exec"

	errorpkg "github.com/giantswarm/crd-docs-generator/error"
	"github.com/giantswarm/microerror"
)

// CloneRepositoryShallow will clone repository in a given directory.
func CloneRepositoryShallow(user string, repo string, tag string, destDir string) error {
	{
		cmd := exec.Command("git", "clone", "-b", tag, "--depth", "1", fmt.Sprintf("https://github.com/%s/%s.git", user, repo), destDir)
		err := cmd.Run()
		if err != nil {
			return microerror.Maskf(errorpkg.ExecutionError, "Could not `git clone` source repository.\nTried to execute: %s\n%s", cmd.String(), err.Error())
		}
	}

	return nil
}
