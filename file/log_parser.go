package file

import (
	"bufio"
	"io"
	"strings"

	logger "github.com/wawakakakyakya/GolangLogger"
	"github.com/wawakakakyakya/check_logs_by_mail/config"
)

type LogParser struct {
	logger *logger.Logger
	// regexps []*regexp.Regexp
	maxLine int
}

func (l *LogParser) Parse(fp io.Reader, fileName string, words []*config.WordConfig) (int, error) {

	var err error
	readSize := 0

	reader := bufio.NewReader(fp)
	isMatched := false
	for {
		// "\n"だとエラーになる
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			l.logger.DebugF("reached to EOF(%s), read file ended.", fileName)
			break
		} else if err != nil {
			l.logger.ErrorF("read file error: %s", fileName)
			return 0, err
		}

		readSize += len(line)
		l.logger.DebugF("readSize(%s): %d", fileName, readSize)

		for _, wc := range words {
			if wc.TargetRegexp.Match(line) {
				s := wc.TargetRegexp.FindAllString(string(line), -1)
				l.logger.DebugF("findallString: %s", strings.Join(s, ""))
			Loop: // ラベル名は何でも良い
				for _, sw := range wc.StopRegexps {
					if sw.Match(line) {
						l.logger.WarnF("file:%s matched, but stop ward(%s) matched too. skip this line.: %s", fileName, sw.String(), line)
						break Loop
					}
				}
				ls := string(line)
				wc.SMTPData.AddMsg(ls)
				l.logger.WarnF("file:%s, line was matched: %s", fileName, line)
				isMatched = true
			} 
		}
	}
	if !isMatched {
		l.logger.InfoF("no line matched in %s", fileName)
	}
	return readSize, err
}

func NewLogParser(fc *config.FileConfig, logger *logger.Logger) *LogParser {
	childLogger := logger.Child("logParser")
	return &LogParser{maxLine: fc.MaxLine, logger: childLogger}
}
