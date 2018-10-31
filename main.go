package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if ln := len(os.Args); ln == 1 {
		fmt.Fprintln(os.Stderr, "[args] srcfile dstpath")
		os.Exit(2)
	} else if ln != 3 {
		fmt.Fprintln(os.Stderr, "invalid args length")
		os.Exit(2)
	}

	srcfile := os.Args[1]
	if fi, err := os.Stat(srcfile); err != nil {
		fmt.Fprintln(os.Stderr, "invalid srcfile")
		os.Exit(2)
	} else if fi.IsDir() {
		fmt.Fprintln(os.Stderr, "srcfile is directory")
		os.Exit(2)
	}

	dstpath := os.Args[2]
	if fi, err := os.Stat(dstpath); err != nil {
		fmt.Fprintln(os.Stderr, "invalid dstpath")
		os.Exit(2)
	} else if !fi.IsDir() {
		fmt.Fprintln(os.Stderr, "dstpath is not directory")
		os.Exit(2)
	}

	dstfiles, err := find(dstpath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := update(srcfile, dstfiles); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func find(path string) ([]string, error) {
	pathfiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var allfiles []string
	for _, pathfile := range pathfiles {
		fullpath := filepath.Join(path, pathfile.Name())
		if pathfile.IsDir() {
			files, err := find(fullpath)
			if err != nil {
				return files, err
			}
			allfiles = append(allfiles, files...)
		} else {
			allfiles = append(allfiles, fullpath)
		}
	}

	return allfiles, nil
}

func update(srcfile string, dstfiles []string) error {
	srcabs, err := filepath.Abs(srcfile)
	if err != nil {
		return err
	}

	srcbase := filepath.Base(srcfile)

	for _, dstfile := range dstfiles {
		dstabs, err := filepath.Abs(dstfile)
		if err != nil {
			return err
		}

		// srcfile does not copy.
		if srcabs == dstabs {
			continue
		}

		if _, dstbase := filepath.Split(dstfile); srcbase == dstbase {
			if err := copy(dstfile, srcfile); err != nil {
				return err
			}
		}
	}

	return nil
}

func copy(dstfile, srcfile string) error {
	src, err := os.Open(srcfile)
	if err != nil {
		return err
	}
	defer src.Close()

	// backup dstfile.
	nows := strings.Split(time.Now().Format("20060102150405.000"), ".")
	newfile := dstfile + "." + nows[0] + nows[1]
	if err := os.Rename(dstfile, newfile); err != nil {
		return err
	}

	dst, err := os.Create(dstfile)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
