package main

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/ulikunitz/xz"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

func writeLines(dst string, lines []string) {
	f, err := os.Create(dst)
	if err != nil {
		log.Panicln(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(f)

	for _, line := range lines {
		if _, err := f.WriteString(fmt.Sprintf("%v\n", line)); err != nil {
			log.Panicln(err)
		}
	}
}

func copyFile(dst string, src string) error {

	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcF.Close()

	info, err := srcF.Stat()
	if err != nil {
		return err
	}

	dstF, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}
	defer dstF.Close()

	if _, err := io.Copy(dstF, srcF); err != nil {
		return err
	}
	return nil
}

func run(prefix string, cmd *exec.Cmd) {
	wg := &sync.WaitGroup{}

	wg.Add(2)

	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Panicln(err)
	}

	errPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Panicln(err)
	}

	if err := cmd.Start(); err != nil {
		log.Panicln(err)
	}

	scanner := bufio.NewScanner(outPipe)
	//scanner.Split(bufio.ScanLines)
	go func() {
		for scanner.Scan() {
			log.Println(prefix, scanner.Text())
		}

		wg.Done()
	}()

	scanner2 := bufio.NewScanner(errPipe)
	//scanner.Split(bufio.ScanLines)
	go func() {
		for scanner2.Scan() {
			log.Println(prefix, scanner2.Text())
		}

		wg.Done()
	}()

	wg.Wait()

	if err := cmd.Wait(); err != nil {
		log.Panicln(err)
	}

}

func exists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	log.Panicln(err)
	return false
}

func download(url string, path string) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	f, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(f)

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		path,
	)

	_, err := io.Copy(io.MultiWriter(f, bar), resp.Body)
	if err != nil {
		log.Panicln("Failed to download file", url, err)
	}
}

func untar(src string, dest string, prefix string) {
	os.RemoveAll(dest)

	if err := os.MkdirAll(dest, 0755); err != nil {
		log.Panicln(err)
	}

	gzipStream, err := os.Open(src)
	if err != nil {
		log.Panicln(err)
	}

	defer gzipStream.Close()

	var uncompressedStream io.Reader

	if strings.HasSuffix(src, ".xz") {
		uncompressedStream, err = xz.NewReader(gzipStream)
		if err != nil {
			log.Fatal("ExtractTarXz: NewReader failed")
		}
	} else {
		uncompressedStream, err = gzip.NewReader(gzipStream)
		if err != nil {
			log.Fatal("ExtractTarGz: NewReader failed")
		}
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		header.Name = strings.TrimPrefix(header.Name, prefix)

		if header.Name == "" {
			continue
		}

		path := filepath.Join(dest, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(path, 0755); err != nil {
				log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
			}
			outFile.Close()

		default:
			log.Fatalf(
				"ExtractTarGz: uknown type: %s in %s",
				header.Typeflag,
				header.Name)
		}

	}
}

func unzip(src string, dest string) {
	os.RemoveAll(dest)

	r, err := zip.OpenReader(src)
	if err != nil {
		log.Panicln(err)
	}
	defer r.Close()

	if err := os.MkdirAll(dest, 0755); err != nil {
		log.Panicln(err)
	}

	prefixes := make(map[string]struct{})

	for _, f := range r.File {
		parts := strings.SplitN(f.Name, string(os.PathSeparator), 2)
		prefixes[parts[0]] = struct{}{}
	}

	if len(prefixes) != 1 {
		log.Panicln("Unexpected prefixes", prefixes)
	}

	var prefix string
	for s, _ := range prefixes {
		prefix = fmt.Sprintf("%v%v", s, string(os.PathSeparator))
	}

	for _, f := range r.File {
		f.Name = strings.TrimPrefix(f.Name, prefix)

		if f.Name == "" {
			continue
		}

		if err := extractAndWriteFile(f, dest); err != nil {
			log.Panicln(err)
		}
	}
}

func extractAndWriteFile(f *zip.File, dest string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer func() {
		if err := rc.Close(); err != nil {
			panic(err)
		}
	}()

	path := filepath.Join(dest, f.Name)

	// Check for ZipSlip (Directory traversal)
	if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
		return fmt.Errorf("illegal file path: %s", path)
	}

	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(path, f.Mode()); err != nil {
			return err
		}
	} else {
		if err := os.MkdirAll(filepath.Dir(path), f.Mode()); err != nil {
			return err
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Panicln(err)
			}
		}()

		_, err = io.Copy(f, rc)
		if err != nil {
			return err
		}
	}
	return nil
}

func cmd(name string, dir string, args ...string) *exec.Cmd {
	cmd := exec.Command(name)
	cmd.Dir = dir
	cmd.Env = os.Environ()

	cmd.Env = append(
		cmd.Env,
		fmt.Sprintf("CFLAGS=-I%v", incDir),
		fmt.Sprintf("CPPFLAGS=-I%v", incDir),
		fmt.Sprintf("CXXFLAGS=-I%v", incDir),
		fmt.Sprintf("LDFLAGS=-L%v", libDir),
		fmt.Sprintf("PKG_CONFIG_PATH=%v/pkgconfig", libDir),
	)

	cmd.Args = append(cmd.Args, args...)

	return cmd
}
