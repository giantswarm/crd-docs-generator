package git

import (
	"fmt"
	"os/exec"

	"github.com/giantswarm/microerror"
)

// CloneRepositoryShallow will clone repository in a given directory.
func CloneRepositoryShallow(user string, repo string, destDir string) error {
	{
		cmd := exec.Command("git", "clone", "--depth", "1", fmt.Sprintf("https://github.com/%s/%s.git", user, repo), destDir)
		err := cmd.Run()
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
