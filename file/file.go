package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

// LowLevel
type FileReadWriter struct {
}

func (f *FileReadWriter) Read() {
}

func (f *FileReadWriter) ReadLine() {
}

func (f *FileReadWriter) Inode(path string) (uint64, error) {
	var inode uint64
	fInfo, err := os.Stat(path)
	if err != nil {
		return inode, err
	}
	stat, ok := fInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return inode, errors.New("not a syscall.Stat_t")
	}
	return stat.Ino, nil

}

func (f *FileReadWriter) GetOneGenerationBeforeFile(path string, oldInode uint64) (string, error) {
	oldFiles, err := filepath.Glob(path + "*")
	if err != nil {
		return "", err
	}
	for _, oldFile := range oldFiles {
		inode, err := f.Inode(oldFile)
		if err != nil {
			return "", err
		}
		if inode == oldInode {
			return oldFile, nil
		}
	}
	return "", fmt.Errorf("one generation before for %s was not found", path)
}
