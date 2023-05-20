package file

import (
	"os"
	"time"

	logger "github.com/wawakakakyakya/GolangLogger"
	"github.com/wawakakakyakya/configloader/json"
)

type PosFile struct {
	Inode     uint64 `json:"inode"`
	LastLine  int64  `json:"last_line"`
	Path      string `json:"path"`
	FileName  string
	UpdatedAt string `json:"updated_at"`
	logger    *logger.Logger
}

func (p *PosFile) Write(inode uint64, lastLine int64) error {
	p.Inode = inode
	p.LastLine = lastLine
	p.UpdatedAt = time.Now().Format("2006-01-02T15:04:05")
	return json.Write(p.FileName, p)
}

func NewPosFile(fileName string, logger *logger.Logger) (*PosFile, error) {
	posLogger := logger.Child("posFile")
	posFile := &PosFile{FileName: fileName, logger: posLogger}
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		posLogger.DebugF("posfile: %s does not exist, make init posfile", fileName)
		posFile.Write(posFile.Inode, posFile.LastLine)
	}

	if err := json.Load(fileName, posFile); err != nil {
		posLogger.ErrorF("load posfile: %s failed", fileName)
		return nil, err
	}
	return posFile, nil
}
