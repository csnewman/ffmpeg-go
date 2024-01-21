package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var (
	tempDir, _   = filepath.Abs("temp")
	downloadsDir = path.Join(tempDir, "downloads")
	buildDir     = path.Join(tempDir, "build")
	tgtDir       = path.Join(tempDir, "tgt")
	incDir       = path.Join(tgtDir, "include")
	libDir       = path.Join(tgtDir, "lib")

	libs = []string{
		"libaom",
		"libass",
		"libavcodec",
		"libavdevice",
		"libavfilter",
		"libavformat",
		"libavutil",
		"libbrotlicommon",
		"libbrotlidec",
		"libbrotlienc",
		"libbz2",
		"libfreetype",
		"libfribidi",
		"libharfbuzz",
		"libpng16",
		"libpostproc",
		"libswresample",
		"libswscale",
		"libunibreak",
		"libz",
	}
)

type OS string

var (
	MacOS OS = "macos"
	Linux OS = "linux"
)

type Builder struct {
	os OS
}

func main() {
	targetOutput, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	//os.RemoveAll(tempDir)
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

	if err := os.MkdirAll(libDir, 0755); err != nil {
		log.Panicln(err)
	}

	os := Linux

	if runtime.GOOS == "darwin" {
		os = MacOS
	}

	b := Builder{
		os: os,
	}

	if os == MacOS {
		buildIconv()
		libs = append(libs, "libiconv")
	}

	buildZlib()
	buildBZ2()
	buildBrotli()
	buildPng()
	buildFribidi()
	buildUnibreak()
	buildFreetype()

	if os == Linux {
		buildExpat()
		buildFontconfig()

		libs = append(libs, "libexpat", "libfontconfig")
	}

	buildHarfbuzz()
	b.buildASS()
	buildAOM()
	b.buildFFmpeg()

	if b.os == Linux {
		combineLinux(targetOutput)
	} else if b.os == MacOS {
		combineMac(targetOutput)
	}
}

func buildExpat() {
	zipPath := path.Join(downloadsDir, "expat.tar.gz")
	srcPath := path.Join(buildDir, "expat")

	if !exists(zipPath) {
		download("https://github.com/libexpat/libexpat/releases/download/R_2_5_0/expat-2.5.0.tar.gz", zipPath)
	}

	untar(zipPath, srcPath, "expat-2.5.0/")

	{
		log.Println("Running configure")

		cmd := cmd(
			path.Join(srcPath, "configure"),
			srcPath,
			fmt.Sprintf("--prefix=%v", tgtDir),
			"--disable-debug",
			"--enable-static",
			"--disable-shared",
			fmt.Sprintf("CFLAGS=-I%v", incDir),
			fmt.Sprintf("CPPFLAGS=-I%v", incDir),
			fmt.Sprintf("LDFLAGS=-L%v", libDir),
		)

		run("[expat configure]", cmd)
	}

	{
		log.Println("Running make")

		cmd := cmd(
			"make",
			srcPath,
			"-j8",
			"install",
		)

		run("[expat make]", cmd)
	}
}

func buildFontconfig() {
	zipPath := path.Join(downloadsDir, "fontconfig.tar.gz")
	srcPath := path.Join(buildDir, "fontconfig")

	if !exists(zipPath) {
		download("https://www.freedesktop.org/software/fontconfig/release/fontconfig-2.15.0.tar.gz", zipPath)
	}

	untar(zipPath, srcPath, "fontconfig-2.15.0/")

	{
		log.Println("Running configure")

		cmd := cmd(
			path.Join(srcPath, "configure"),
			srcPath,
			fmt.Sprintf("--prefix=%v", tgtDir),
			"--disable-debug",
			"--enable-static",
			"--disable-shared",
			fmt.Sprintf("CFLAGS=-I%v", incDir),
			fmt.Sprintf("CPPFLAGS=-I%v", incDir),
			fmt.Sprintf("LDFLAGS=-L%v", libDir),
		)

		run("[fontconfig configure]", cmd)
	}

	{
		log.Println("Running make")

		cmd := cmd(
			"make",
			srcPath,
			"-j8",
			"install",
		)

		run("[fontconfig make]", cmd)
	}
}

func buildFribidi() {
	zipPath := path.Join(downloadsDir, "fribidi.tar.xz")
	srcPath := path.Join(buildDir, "fribidi")

	if !exists(zipPath) {
		download("https://github.com/fribidi/fribidi/releases/download/v1.0.13/fribidi-1.0.13.tar.xz", zipPath)
	}

	untar(zipPath, srcPath, "fribidi-1.0.13/")

	{
		log.Println("Running configure")

		cmd := cmd(
			path.Join(srcPath, "configure"),
			srcPath,
			fmt.Sprintf("--prefix=%v", tgtDir),
			"--disable-dependency-tracking",
			"--disable-debug",
			"--disable-silent-rules",
			"--enable-static",
			fmt.Sprintf("CFLAGS=-I%v", incDir),
			fmt.Sprintf("CPPFLAGS=-I%v", incDir),
			fmt.Sprintf("LDFLAGS=-L%v", libDir),
		)

		run("[fribidi configure]", cmd)
	}

	{
		log.Println("Running make")

		cmd := cmd(
			"make",
			srcPath,
			"-j8",
			"install",
		)

		run("[fribidi make]", cmd)
	}
}

func buildIconv() {
	zipPath := path.Join(downloadsDir, "iconv.tar.gz")
	srcPath := path.Join(buildDir, "iconv")

	if !exists(zipPath) {
		download("https://ftp.gnu.org/pub/gnu/libiconv/libiconv-1.17.tar.gz", zipPath)
	}

	untar(zipPath, srcPath, "libiconv-1.17/")

	{
		log.Println("Running configure")

		cmd := cmd(
			path.Join(srcPath, "configure"),
			srcPath,
			fmt.Sprintf("--prefix=%v", tgtDir),
			"--disable-dependency-tracking",
			"--disable-debug",
			"--enable-extra-encodings",
			"--enable-static",
			fmt.Sprintf("CFLAGS=-I%v", incDir),
			fmt.Sprintf("CPPFLAGS=-I%v", incDir),
			fmt.Sprintf("LDFLAGS=-L%v", libDir),
		)

		run("[iconv configure]", cmd)
	}

	{
		log.Println("Running make")

		cmd := cmd(
			"make",
			srcPath,
			"-j8",
			"-f", "Makefile.devel",
		)
		run("[iconv make]", cmd)
	}

	{
		log.Println("Running install")

		cmd := cmd(
			"make",
			srcPath,
			"install",
		)

		run("[iconv install]", cmd)
	}
}

func buildZlib() {
	zipPath := path.Join(downloadsDir, "zlib.zip")
	srcPath := path.Join(buildDir, "zlib")

	if !exists(zipPath) {
		download("https://github.com/madler/zlib/releases/download/v1.3/zlib13.zip", zipPath)
	}

	unzip(zipPath, srcPath)

	{
		log.Println("Running configure ??")

		cmd := cmd(
			path.Join(srcPath, "configure"),
			srcPath,
			fmt.Sprintf("--prefix=%v", tgtDir),
			"--static",
		)

		run("[zlib configure]", cmd)
	}

	{
		log.Println("Running make")

		cmd := cmd(
			"make",
			srcPath,
			"-j8",
			"install",
		)

		run("[zlib make]", cmd)
	}
}

func buildBZ2() {
	zipPath := path.Join(downloadsDir, "bz2.zip")
	srcPath := path.Join(buildDir, "bz2")

	if !exists(zipPath) {
		download("https://gitlab.com/bzip2/bzip2/-/archive/bzip2-1.0.8/bzip2-bzip2-1.0.8.zip", zipPath)
	}

	unzip(zipPath, srcPath)

	{
		log.Println("Running make")

		cmd := cmd(
			"make",
			srcPath,
			"-j8",
			"install",
			fmt.Sprintf("PREFIX=%v", tgtDir),
			fmt.Sprintf("CFLAGS=-I%v", incDir),
			fmt.Sprintf("CPPFLAGS=-I%v", incDir),
			fmt.Sprintf("LDFLAGS=-L%v", libDir),
		)

		run("[bz2 make]", cmd)
	}
}

func buildBrotli() {
	zipPath := path.Join(downloadsDir, "brotli.zip")
	srcPath := path.Join(buildDir, "brotli")

	if !exists(zipPath) {
		download("https://codeload.github.com/google/brotli/zip/refs/heads/v1.1", zipPath)
	}

	unzip(zipPath, srcPath)

	{
		log.Println("Running cmake")

		cmd := cmd(
			"cmake",
			srcPath,
			"-G",
			"Unix Makefiles",
			fmt.Sprintf("-DCMAKE_INSTALL_PREFIX=%v", tgtDir),
			"-DCMAKE_BUILD_TYPE=Release",
			"-DCMAKE_FIND_FRAMEWORK=last",
			"-DCMAKE_VERBOSE_MAKEFILE=ON",
			"-Wno-dev",
			"-DBUILD_SHARED_LIBS=OFF",
			".",
		)

		run("[brotli cmake]", cmd)
	}

	{
		log.Println("Running make")

		cmd := cmd(
			"make",
			srcPath,
			"-j8",
			"install",
		)

		run("[brotli make]", cmd)
	}
}

func buildPng() {
	zipPath := path.Join(downloadsDir, "png.zip")
	srcPath := path.Join(buildDir, "png")

	if !exists(zipPath) {
		download("https://codeload.github.com/pnggroup/libpng/zip/refs/heads/libpng16", zipPath)
	}

	unzip(zipPath, srcPath)

	{
		log.Println("Running configure")

		cmd := cmd(
			path.Join(srcPath, "configure"),
			srcPath,
			fmt.Sprintf("--prefix=%v", tgtDir),
			"--disable-dependency-tracking",
			"--disable-silent-rules",
			"--disable-shared",
			"--enable-static",
			fmt.Sprintf("CFLAGS=-I%v", incDir),
			fmt.Sprintf("CPPFLAGS=-I%v", incDir),
			fmt.Sprintf("LDFLAGS=-L%v", libDir),
		)

		run("[png configure]", cmd)
	}

	{
		log.Println("Running make")

		cmd := cmd(
			"make",
			srcPath,
			"-j8",
			"install",
		)

		run("[png make]", cmd)
	}
}

func buildUnibreak() {
	zipPath := path.Join(downloadsDir, "unibreak.tar.gz")
	srcPath := path.Join(buildDir, "unibreak")

	if !exists(zipPath) {
		download("https://github.com/adah1972/libunibreak/releases/download/libunibreak_5_1/libunibreak-5.1.tar.gz", zipPath)
	}

	untar(zipPath, srcPath, "libunibreak-5.1/")

	{
		log.Println("Running configure")

		cmd := cmd(
			path.Join(srcPath, "configure"),
			srcPath,
			fmt.Sprintf("--prefix=%v", tgtDir),
			"--disable-shared",
			"--enable-static",
			fmt.Sprintf("CFLAGS=-I%v", incDir),
			fmt.Sprintf("CPPFLAGS=-I%v", incDir),
			fmt.Sprintf("LDFLAGS=-L%v", libDir),
		)

		run("[unibreak configure]", cmd)
	}

	{
		log.Println("Running make")

		cmd := cmd(
			"make",
			srcPath,
			"-j8",
			"install",
		)

		run("[unibreak make]", cmd)
	}
}

func buildFreetype() {
	zipPath := path.Join(downloadsDir, "freetype.tar.xz")
	srcPath := path.Join(buildDir, "freetype")

	if !exists(zipPath) {
		download("https://deac-ams.dl.sourceforge.net/project/freetype/freetype2/2.13.2/freetype-2.13.2.tar.xz", zipPath)
	}

	untar(zipPath, srcPath, "freetype-2.13.2/")

	{
		log.Println("Running configure")

		cmd := cmd(
			path.Join(srcPath, "configure"),
			srcPath,
			fmt.Sprintf("--prefix=%v", tgtDir),
			"--without-harfbuzz",
			// brotlii breaks harfbuzz building
			"--without-brotli",
			"--disable-shared",
			"--enable-static",
			fmt.Sprintf("CFLAGS=-I%v", incDir),
			fmt.Sprintf("CPPFLAGS=-I%v", incDir),
			fmt.Sprintf("LDFLAGS=-L%v", libDir),
		)

		run("[freetype configure]", cmd)
	}

	{
		log.Println("Running make")

		cmd := cmd(
			"make",
			srcPath,
			"-j8",
			"install",
		)

		run("[freetype make]", cmd)
	}
}

func buildHarfbuzz() {
	zipPath := path.Join(downloadsDir, "harfbuzz.tar.xz")
	srcPath := path.Join(buildDir, "harfbuzz")

	if !exists(zipPath) {
		download("https://github.com/harfbuzz/harfbuzz/releases/download/8.3.0/harfbuzz-8.3.0.tar.xz", zipPath)
	}

	untar(zipPath, srcPath, "harfbuzz-8.3.0/")

	{
		log.Println("Running setup")

		cmd := cmd(
			"meson",
			srcPath,
			"setup",
			"build",
			fmt.Sprintf("--prefix=%v", tgtDir),
			fmt.Sprintf("--libdir=%v", libDir),
			"--buildtype=release",
			"--default-library=static",
			"-Dcairo=disabled",
			"-Dcoretext=enabled",
			"-Dfreetype=enabled",
			//"-Dglib=enabled",
			//"-Dgobject=enabled",
			//"-Dgraphite=enabled",
			//"-Dicu=enabled",
			//"-Dintrospection=enabled",
			"-Dtests=disabled",
		)

		run("[harfbuzz setup]", cmd)
	}

	{
		log.Println("Running compile")

		cmd := cmd(
			"meson",
			srcPath,
			"compile",
			"-C",
			"build",
			"--verbose",
		)

		run("[harfbuzz compile]", cmd)
	}

	{
		log.Println("Running install")

		cmd := cmd(
			"meson",
			srcPath,
			"install",
			"-C",
			"build",
		)

		run("[harfbuzz install]", cmd)
	}
}

func (b *Builder) buildASS() {
	zipPath := path.Join(downloadsDir, "ass.tar.gz")
	srcPath := path.Join(buildDir, "ass")

	if !exists(zipPath) {
		download("https://github.com/libass/libass/releases/download/0.17.1/libass-0.17.1.tar.gz", zipPath)
	}

	untar(zipPath, srcPath, "libass-0.17.1/")

	{
		log.Println("Running configure")

		cmd := cmd(
			path.Join(srcPath, "configure"),
			srcPath,
			fmt.Sprintf("--prefix=%v", tgtDir),
			"--disable-shared",
		)

		// libass uses coretext on macOS, fontconfig on Linux
		if b.os == MacOS {
			cmd.Args = append(
				cmd.Args,
				"--disable-fontconfig",
			)
		}

		cmd.Args = append(cmd.Args, fmt.Sprintf("CFLAGS=-I%v", incDir))

		run("[ass configure]", cmd)
	}

	{
		log.Println("Running make")

		cmd := cmd(
			"make",
			srcPath,
			"-j8",
			"install",
		)

		run("[ass make]", cmd)
	}
}

func buildAOM() {
	zipPath := path.Join(downloadsDir, "aom.tar.gz")
	srcPath := path.Join(buildDir, "aom")
	buildPath := path.Join(buildDir, "aom-build")

	if !exists(zipPath) {
		//https://aomedia.googlesource.com/aom/+/refs/tags/v3.8.1
		download("https://aomedia.googlesource.com/aom/+archive/bb6430482199eaefbeaaa396600935082bc43f66.tar.gz", zipPath)
	}

	untar(zipPath, srcPath, "")

	if err := os.MkdirAll(buildPath, 0755); err != nil {
		log.Panicln(err)
	}

	{
		log.Println("Running cmake")

		cmd := cmd(
			"cmake",
			buildPath,
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

		cmd := cmd(
			"make",
			buildPath,
			"-j8",
			"install",
		)

		run("[aom make]", cmd)
	}
}

func (b *Builder) buildFFmpeg() {
	zipPath := path.Join(downloadsDir, "ffmpeg.zip")
	buildPath := path.Join(buildDir, "ffmpeg")

	if !exists(zipPath) {
		download("https://codeload.github.com/FFmpeg/FFmpeg/zip/refs/heads/release/6.1", zipPath)
	}

	unzip(zipPath, buildPath)

	{
		log.Println("Running configure")

		cmd := cmd(
			path.Join(buildPath, "configure"),
			buildPath,
			"--cc=/usr/bin/clang",
			fmt.Sprintf("--prefix=%v", tgtDir),
			"--pkg-config-flags=--static",
			fmt.Sprintf("--extra-cflags=-I%v", incDir),
			fmt.Sprintf("--extra-ldflags=-L%v/lib", tgtDir),
			"--disable-autodetect",
			"--disable-programs",
			"--enable-pic",
			"--enable-gpl",
			"--enable-version3",
			"--enable-static",

			// Enable libs
			"--enable-libaom",
			"--enable-libass",
			"--enable-libfreetype",
			"--enable-libfribidi",
			"--enable-libharfbuzz",
		)

		if b.os == Linux {
			cmd.Args = append(
				cmd.Args,
				"--enable-libfontconfig",
			)
		}

		run("[ffmpeg configure]", cmd)
	}

	{
		log.Println("Running make")

		cmd := cmd(
			"make",
			buildPath,
			"-j8",
			"install",
		)

		run("[ffmpeg make]", cmd)
	}
}
