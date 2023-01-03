package file

import (
	"fmt"
	"os"
	"syscall"
)

type FileReadWriter struct {
}

func (f *FileReadWriter) Read() {
}

func (f *FileReadWriter) ReadLine() {
}

func (f *FileReadWriter) getInode(path string) (int, error) {
	var inode int
	fInfo, err := os.Stat(path)
	if err != nil {
		return inode, err
	}
	stat, ok := fInfo.Sys().(*syscall.Stat_t)
    if !ok {
        fmt.Printf("Not a syscall.Stat_t")
        return
    }
	return fInfo.

}
