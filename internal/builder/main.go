package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var (
	tempDir, _   = filepath.Abs("temp")
	downloadsDir = path.Join(tempDir, "downloads")
	buildDir     = path.Join(tempDir, "build")
	tgtDir       = path.Join(tempDir, "tgt")
)

func main() {
	targetOutput, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	os.RemoveAll(tempDir)
	os.RemoveAll(buildDir)
	os.RemoveAll(tgtDir)

	if err := os.MkdirAll(downloadsDir, 0755); err != nil {
		log.Panicln(err)
	}

	if err := os.MkdirAll(buildDir, 0755); err != nil {
		log.Panicln(err)
	}

	if err := os.MkdirAll(tgtDir, 0755); err != nil {
		log.Panicln(err)
	}

	buildAOM()
	buildFFmpeg()

	combine(targetOutput)
}

func buildAOM() {
	zipPath := path.Join(downloadsDir, "aom.tar.gz")
	srcPath := path.Join(buildDir, "aom")
	buildPath := path.Join(buildDir, "aom-build")

	if !exists(zipPath) {
		//https://aomedia.googlesource.com/aom/+/refs/tags/v3.8.1
		download("https://aomedia.googlesource.com/aom/+archive/bb6430482199eaefbeaaa396600935082bc43f66.tar.gz", zipPath)
	}

	untar(zipPath, srcPath)

	if err := os.MkdirAll(buildPath, 0755); err != nil {
		log.Panicln(err)
	}

	{
		log.Println("Running cmake")

		cmd := exec.Command("cmake")
		cmd.Dir = buildPath

		cmd.Args = append(
			cmd.Args,
			"-G",
			"Unix Makefiles",
			fmt.Sprintf("-DCMAKE_INSTALL_PREFIX=%v", tgtDir),
			"-DENABLE_TESTS=OFF",
			// TODO: nasm
			//-DENABLE_NASM=on
			srcPath,
		)
		run("[aom cmake]", cmd)
	}

	{
		log.Println("Running make")

		cmd := exec.Command("make")
		cmd.Dir = buildPath

		cmd.Args = append(
			cmd.Args,
			"-j8",
			"install",
		)
		run("[aom make]", cmd)
	}
}

func buildFFmpeg() {
	zipPath := path.Join(downloadsDir, "ffmpeg.zip")
	buildPath := path.Join(buildDir, "ffmpeg")

	if !exists(zipPath) {
		download("https://codeload.github.com/FFmpeg/FFmpeg/zip/refs/heads/release/6.1", zipPath)
	}

	unzip(zipPath, buildPath)

	{
		log.Println("Running configure")

		cmd := exec.Command(path.Join(buildPath, "configure"))
		cmd.Dir = buildPath

		cmd.Env = append(
			cmd.Env,
			fmt.Sprintf("PKG_CONFIG_PATH=%v/lib/pkgconfig", tgtDir),
		)

		cmd.Args = append(
			cmd.Args,
			"--cc=/usr/bin/clang",
			fmt.Sprintf("--prefix=%v", tgtDir),
			"--pkg-config-flags=--static",
			fmt.Sprintf("--extra-cflags=-I%v/include", tgtDir),
			fmt.Sprintf("--extra-ldflags=-L%v/lib", tgtDir),
			"--disable-autodetect",
			"--disable-programs",
			"--enable-pic",
			"--enable-gpl",
			"--enable-version3",
			"--enable-static",

			// Enable libs
			"--enable-libaom",
		)
		run("[ffmpeg configure]", cmd)
	}

	{
		log.Println("Running make")

		cmd := exec.Command("make")
		cmd.Dir = buildPath

		cmd.Args = append(
			cmd.Args,
			"-j8",
			"install",
		)
		run("[ffmpeg make]", cmd)
	}
}
