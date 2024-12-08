package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile           = errors.New("unsupported file")
	ErrOffsetExceedsFileSize     = errors.New("offset exceeds file size")
	ErrOffsetOrLimitLessThanZero = errors.New("offset or limit less than zero")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	//fromPath = "testdata/input.txt"
	//toPath = "out.txt"

	fmt.Println("Copy from:", fromPath, "to:", toPath)
	fmt.Println("Copy offset:", offset, "limit:", limit)

	if limit < 0 || offset < 0 {
		return ErrOffsetOrLimitLessThanZero
	}

	fileFrom, err := os.Open(fromPath)

	if err != nil {
		return err
	}

	fi, err := fileFrom.Stat()
	if err != nil {
		return err
	}

	defer func() {
		_ = fileFrom.Close()
	}()

	if limit+offset > fi.Size() {
		return ErrOffsetExceedsFileSize
	}

	fileTo, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = fileTo.Close()
	}()

	_, err = fileFrom.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	if limit == 0 || limit > fi.Size()-offset {
		limit = fi.Size() - offset
	}

	reader := io.LimitReader(fileFrom, limit)
	progressBar := pb.Full.Start64(limit)
	defer progressBar.Finish()
	reader = progressBar.NewProxyReader(reader)

	//progressBar :=

	//if offset == 0 && limit == 0 {
	_, err = io.Copy(fileTo, reader)
	//} else if offset == 0 {
	//	_, err = io.CopyN(fileTo, fileFrom, limit)
	//}
	if err != nil {
		return err
	}

	return nil
}

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()

	err := Copy(from, to, offset, limit)
	if err != nil {
		log.Fatal(err)
	}
}
