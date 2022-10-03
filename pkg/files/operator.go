package files

import (
	"fmt"
	"path/filepath"
	"sync"
)

// GlobalFileSafeOperator is a global file-named thread-safe locker, it can package the operators.
var GlobalFileSafeOperator FileNameThreadSafeLocker

func init() {
	GlobalFileSafeOperator = FileNameThreadSafeLocker{
		innerLock:  &sync.Mutex{},
		fileMapper: map[string]struct{}{},
	}
}

// Operator defines a file-operator interface, contain the Read and Write. The Operator not limit the thread-safety.
// And operator is unsupportable performance. So if you want to keep the high-performance, please use other tools. All
// the file operator should use at application start age. When application bootstrap job complete, do not use the
// operator.
type Operator interface {
	// Read will read a path file and set value to obj. The obj should a point. The read rule follow implementation
	// settings, different implements have different deserialize behavior.
	//
	// Specify, when obj can't deserialize, the file value should be discarded. All the un-support-serialize field will
	// be skipped.
	Read(path string, obj any) error

	// Write will write the obj to the specify file, if write error, return error directly(the error may create by file
	// system). Like Read, when object can't serialize, no value will be written.
	Write(path string, obj any) error
}

// FileNameThreadSafeLocker will keep the specify file thread-safe. And can register more file operator, all the
// operator will use one lock for one file. In other words, if you create two FileNameThreadSafeLocker, the file may
// read and write at same time. Or write twice at one time.
//
// The files package is already create one instance to quick deal the files.
type FileNameThreadSafeLocker struct {
	innerLock  *sync.Mutex
	fileMapper map[string]struct{}
}

// LockFile will lock the file by path. If function occur error when get file abs, return error directly, and the file
// not be lock. But the abs is not unique, so if input file has more than one abs value, it may be haven logic error.
func (fo *FileNameThreadSafeLocker) LockFile(file string) error {
	fAbsPath, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	fo.innerLock.Lock()
	if _, ok := fo.fileMapper[fAbsPath]; !ok {
		// insert the lock flag
		fo.fileMapper[fAbsPath] = struct{}{}
	}
	fo.innerLock.Unlock()

	return nil
}

// UnlockFile will unlock a file. Like file lock, the abs path may be different, so if real different, the logic maybe
// have error.
func (fo *FileNameThreadSafeLocker) UnlockFile(file string) error {
	fAbsPath, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	fo.innerLock.Lock()
	defer fo.innerLock.Unlock()

	if _, ok := fo.fileMapper[fAbsPath]; !ok {
		return fmt.Errorf("Not lock the file: %s ", file)
	}

	delete(fo.fileMapper, fAbsPath)
	return nil
}

// Package a un-thread-safe operator to a thread-safe operator.
func (fo *FileNameThreadSafeLocker) Package(fileOperator Operator) Operator {
	return &threadSafeFileOperator{
		innerFileLocker: fo,
		innerOperator:   fileOperator,
	}
}

// threadSafeFileOperator --------------------------------------------------

// threadSafeFileOperator is a thread safe file operator, and should create by FileNameThreadSafeLocker.Package().
type threadSafeFileOperator struct {
	innerFileLocker *FileNameThreadSafeLocker
	innerOperator   Operator
}

func (t *threadSafeFileOperator) Read(path string, obj any) error {
	if err := t.innerFileLocker.LockFile(path); err != nil {
		return err
	}
	defer func() { _ = t.innerFileLocker.UnlockFile(path) }()

	return t.innerOperator.Read(path, obj)
}

func (t *threadSafeFileOperator) Write(path string, obj any) error {
	if err := t.innerFileLocker.LockFile(path); err != nil {
		return err
	}
	defer func() { _ = t.innerFileLocker.UnlockFile(path) }()

	return t.innerOperator.Write(path, obj)
}
