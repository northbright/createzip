// Package createzip is a helper to create a new zip file and add files into the zip file.
package createzip

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// ZipFile is wrapper of zip.Writer.
type ZipFile struct {
	w *zip.Writer
}

var (
	DEBUG bool = false // Set to true to enable debug messages from this package.
)

// New creates a new zip file. It should call Close() after use.
func New(w io.Writer) (zf *ZipFile) {
	return &ZipFile{zip.NewWriter(w)}
}

// NewForHTTP creates the downloadable zip for HTTP server dynamically.
func NewForHTTP(w http.ResponseWriter, zipFileName string) (zf *ZipFile) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipFileName))

	return New(w)
}

// Close closes the writer of zip. It should call Close() after use.
func (zf *ZipFile) Close() {
	zf.w.Close()
}

// Add creates a new file in the zip and copy the content from io.Reader.
//
//   Params:
//       fileNameInZip: file name(path) in the zip.
//       r: io.Reader. It will copy the io.Reader to the new create file's writer.
func (zf *ZipFile) Add(fileNameInZip string, r io.Reader) (err error) {
	w, err := zf.w.Create(fileNameInZip)
	if err != nil {
		if DEBUG {
			fmt.Printf("zw.Create(%v) err: %v\n", fileNameInZip, err)
		}
		return err
	}

	if _, err = io.Copy(w, r); err != nil {
		if DEBUG {
			fmt.Printf("io.Copy(w, r) err: %v\n", err)
		}
		return err
	}

	return nil
}

// AddFile creates a new file in the zip and copy the content from the original source file.
//
//   Params:
//       srcFilePath: original file path.
//       fileNameInZip: file name(path) in the zip. If it's empty(""), the file name will be root dir of the zip("./") + file name without dir of original source file.
func (zf *ZipFile) AddFile(srcFilePath, fileNameInZip string) (err error) {
	fi, err := os.Open(srcFilePath)
	if err != nil {
		if DEBUG {
			fmt.Printf("os.Open(%v) err: %v\n", srcFilePath, err)
		}
		return err
	}
	defer fi.Close()

	// If file name in zip is "", use the file name without dir.
	if fileNameInZip == "" {
		_, fileNameInZip = filepath.Split(srcFilePath)
	}

	r := bufio.NewReader(fi)

	return zf.Add(fileNameInZip, r)
}
