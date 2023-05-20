package config

import (
	"fmt"
	"regexp"

	"github.com/wawakakakyakya/check_logs_by_mail/smtp"
	yml "github.com/wawakakakyakya/configloader/yml"
)

type YamlConfigs struct {
	Files        []*FileConfig `yaml:"files"`
	GlobalConfig *GlobalConfig `yaml:"global_config"`
}

type GlobalConfig struct {
	LogConfig *LogConfig  `yaml:"log"`
	SMTP      *SMTPConfig `yaml:"smtp"`
}

func initConfig(files []*FileConfig) error {
	for _, fc := range files {
		for _, w := range fc.Words {
			targetReg, err := regexp.Compile(w.TargetWord)
			if err != nil {
				return err
			}
			w.SMTPData = smtp.NewSMTPData(w.Recipients, w.Subject)
			w.TargetRegexp = targetReg
			for _, s := range w.StopWords {
				stopReg, err := regexp.Compile(s)
				if err != nil {
					return err
				}
				w.StopRegexps = append(w.StopRegexps, stopReg)
			}

			w.SMTPData = smtp.NewSMTPData(w.Recipients, w.Subject)
		}
	}
	return nil
}

func LoadYamlConfig() (YamlConfigs, error) {
	defaultLogConfig := NewDefaultLogConfig()
	globalConfig := &GlobalConfig{LogConfig: &defaultLogConfig}
	ycArray := YamlConfigs{GlobalConfig: globalConfig}
	err := yml.Load(configPath, &ycArray)

	if err != nil {
		return ycArray, err
	}
	if err := initConfig(ycArray.Files); err != nil {
		fmt.Println("[ERROR]set regexp failed")
		return ycArray, err
	}

	return ycArray, nil
}
