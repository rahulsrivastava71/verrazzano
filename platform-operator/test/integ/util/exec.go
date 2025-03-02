// Copyright (C) 2020, 2021, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/onsi/ginkgo/v2"
)

// RunCommand runs an external process, captures the stdout
// and stderr, as well as streaming them to the real output
// streams in real time
func RunCommand(commandLine string) (string, string) {
	ginkgo.GinkgoWriter.Write([]byte("[DEBUG] RunCommand: " + commandLine + "\n"))
	parts := strings.Split(commandLine, " ")
	var cmd *exec.Cmd
	if len(parts) < 1 {
		ginkgo.Fail("No command provided")
	} else if len(parts) == 1 {
		cmd = exec.Command(parts[0], "") //nolint:gosec //#nosec G204
	} else {
		cmd = exec.Command(parts[0], parts[1:]...) //nolint:gosec //#nosec G204
	}
	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		ginkgo.Fail("cmd.Start() failed with " + err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		wg.Done()
	}()

	_, errStderr = io.Copy(stderr, stderrIn)
	wg.Wait()

	cmd.Wait()
	if errStdout != nil || errStderr != nil {
		ginkgo.Fail("failed to capture stdout or stderr")
	}
	outStr, errStr := stdoutBuf.String(), stderrBuf.String()
	return outStr, errStr
}

// Kubectl runs kubectl in an external process, captures the stdout
// and stderr, as well as streaming them to the real output streams in real time
func Kubectl(args string) (string, string) {
	commandLine := fmt.Sprintf("kubectl --kubeconfig %v %v", GetKubeconfig(), args)
	return RunCommand(commandLine)
}
