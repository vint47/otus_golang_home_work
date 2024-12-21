package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile           = errors.New("unsupported file")
	ErrOffsetExceedsFileSize     = errors.New("offset exceeds file size")
	ErrOffsetOrLimitLessThanZero = errors.New("offset or limit less than zero")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fmt.Println("Copy from:", fromPath, "to:", toPath)
	fmt.Println("Copy offset:", offset, "limit:", limit)

	if limit < 0 || offset < 0 {
		return ErrOffsetOrLimitLessThanZero
	}

	fileFrom, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	fi, err := fileFrom.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}

	defer func() {
		_ = fileFrom.Close()
	}()

	if limit+offset > fi.Size() {
		return ErrOffsetExceedsFileSize
	}

	fileTo, err := os.Create(toPath)
	if err != nil {
		return ErrUnsupportedFile
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

	_, err = io.Copy(fileTo, reader)
	if err != nil {
		return err
	}

	return nil
}
