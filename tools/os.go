/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2020/6/12 14:39
*/

package tools

import (
	"bytes"
	"os/exec"
)

func ExecLinuxApp(command string) (outBytes []byte, err error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	Logger.Debug("Stdout=" + stdout.String())
	Logger.Error("Stderr=" + stderr.String())
	return
}
