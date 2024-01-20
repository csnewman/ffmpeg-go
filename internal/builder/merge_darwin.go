package main

import (
	"log"
	"os/exec"
	"path"
	"path/filepath"
)

func combine(output string) {
	{
		matches, err := filepath.Glob(path.Join(tgtDir, "lib", "*.a"))
		if err != nil {
			log.Panicln(err)
		}

		for _, match := range matches {
			log.Println(match)
		}

		log.Println("Running merge")
		cmd := exec.Command("libtool")
		cmd.Args = append(
			cmd.Args,
			"-static",
			"-o",
			output,
		)
		cmd.Args = append(
			cmd.Args,
			matches...,
		)
		run("[merge]", cmd)
	}

	{

		log.Println("Running strip")
		cmd := exec.Command("strip")
		cmd.Args = append(
			cmd.Args,
			"-S",
			output,
		)
		run("[strip]", cmd)
	}
}
