
package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func say(i int) {
	//cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("sleep %d", i))
	var cmd * exec.Cmd
	if i == 999 {
		cmd = exec.Command("/bin/bash", "-c", `sleep 77`)
	} else {
		cmd = exec.Command("/bin/bash", "-c", `sleep 35`)
	}
	//cmd := exec.Command("/bin/bash", "-c", `df -lh`)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
	fmt.Printf("stdout:\n\n %s", bytes)
	fmt.Printf("index:\n\n %d", i)
}

func main() {
	for i := 0; i < 20; i++ {
		go say(i)
	}
	say(999)
}