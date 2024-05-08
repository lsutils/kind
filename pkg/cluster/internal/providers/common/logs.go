package common

import (
	"os"
	"path/filepath"

	"github.com/lsutils/kind/pkg/cluster/nodes"
	"github.com/lsutils/kind/pkg/errors"
	"github.com/lsutils/kind/pkg/exec"
)

// CollectLogs provides the common functionality
// to get various debug info from the node
func CollectLogs(n nodes.Node, dir string) error {
	execToPathFn := func(cmd exec.Cmd, path string) func() error {
		return func() error {
			f, err := FileOnHost(filepath.Join(dir, path))
			if err != nil {
				return err
			}
			defer f.Close()
			return cmd.SetStdout(f).SetStderr(f).Run()
		}
	}

	return errors.AggregateConcurrent([]func() error{
		// record info about the node container
		execToPathFn(
			n.Command("cat", "/kind/version"),
			"kubernetes-version.txt",
		),
		execToPathFn(
			n.Command("journalctl", "--no-pager"),
			"journal.log",
		),
		execToPathFn(
			n.Command("journalctl", "--no-pager", "-u", "kubelet.service"),
			"kubelet.log",
		),
		execToPathFn(
			n.Command("journalctl", "--no-pager", "-u", "containerd.service"),
			"containerd.log",
		),
		execToPathFn(
			n.Command("crictl", "images"),
			"images.log",
		),
	})
}

// FileOnHost is a helper to create a file at path
// even if the parent directory doesn't exist
// in which case it will be created with ModePerm
func FileOnHost(path string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return nil, err
	}
	return os.Create(path)
}
