package main

import (
	"errors"
	"io"
	"os"
	"syscall"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile           = errors.New("unsupported file")
	ErrOffsetExceedsFileSize     = errors.New("offset exceeds file size")
	ErrOffsetOrLimitLessThanZero = errors.New("offset or limit less than zero")
	ErrSameFile                  = errors.New("same file")
)

func isFileSame(fromPath, toPath string) (bool, error) {
	if fromPath == toPath {
		return true, nil
	}

	fileFrom, err := os.Stat(fromPath)
	if err != nil {
		return false, err
	}

	fileTo, err := os.Stat(toPath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	fileFromData, okFrom := fileFrom.Sys().(*syscall.Stat_t)
	fileToData, okTo := fileTo.Sys().(*syscall.Stat_t)

	if !okFrom || !okTo {
		return false, nil
	}

	isSame := fileFromData.Dev == fileToData.Dev && fileFromData.Ino == fileToData.Ino

	return isSame, nil
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	if limit < 0 || offset < 0 {
		return ErrOffsetOrLimitLessThanZero
	}

	if isSame, err := isFileSame(fromPath, toPath); err != nil || isSame {
		return ErrSameFile
	}

	fileFrom, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	defer func() {
		_ = fileFrom.Close()
	}()

	fi, err := fileFrom.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}

	if !fi.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > fi.Size() {
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
