package main

//var (
//	ErrUnsupportedFile       = errors.New("unsupported file")
//	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
//)
//
//func Copy(fromPath, toPath string, offset, limit int64) error {
//	//fromPath = "testdata/input.txt"
//	//toPath = "out.txt"
//
//	fmt.Println("Copy from:", fromPath, "to:", toPath)
//	fmt.Println("Copy offset:", offset, "limit:", limit)
//
//	fileFrom, err := os.Open(fromPath)
//
//	if err != nil {
//		return err
//	}
//
//	fi, err := fileFrom.Stat()
//	if err != nil {
//		return err
//	}
//	if limit+offset > fi.Size() {
//		return ErrOffsetExceedsFileSize
//	}
//
//	defer func() {
//		_ = fileFrom.Close()
//	}()
//
//	fileTo, err := os.Create(toPath)
//	if err != nil {
//		return err
//	}
//	defer func() {
//		_ = fileTo.Close()
//	}()
//
//	_, err = fileFrom.Seek(offset, io.SeekStart)
//	if err != nil {
//		return err
//	}
//
//	if offset == 0 && limit == 0 {
//		_, err = io.Copy(fileTo, fileFrom)
//	} else if offset == 0 {
//		_, err = io.CopyN(fileTo, fileFrom, limit)
//	}
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
