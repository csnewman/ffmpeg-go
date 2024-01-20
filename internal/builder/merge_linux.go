package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func combine(output string) {
	matches, err := filepath.Glob(path.Join(tgtDir, "lib", "*.a"))
	if err != nil {
		log.Panicln(err)
	}

	var lines []string

	lines = append(lines, fmt.Sprintf("create %v", output))

	for _, match := range matches {
		log.Println(match)
		lines = append(lines, fmt.Sprintf("addlib %v", match))
	}

	lines = append(lines, "save", "end")
	mriAll := strings.Join(lines, "\n")

	log.Println(mriAll)

	{
		log.Println("Running merge")
		cmd := exec.Command("ar")
		cmd.Args = append(
			cmd.Args,
			"-M",
		)

		cmd.Stdin = bytes.NewBufferString(mriAll)

		run("[merge]", cmd)
	}

	{
		log.Println("Running strip")
		cmd := exec.Command("strip")
		cmd.Args = append(
			cmd.Args,
			"-d",
			output,
		)
		run("[strip]", cmd)
	}
}
