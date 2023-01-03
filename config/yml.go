package config

import (
	yml "github.com/wawakakakyakya/configloader/yml"
)

type YamlConfigs struct {
	LogConfig LogConfig     `yaml:"log"`
	Files     []*FileConfig `yaml:"files"`
	SMTP      SMTPConfig    `yaml:"smtp"`
}

func LoadYamlConfig() (YamlConfigs, error) {
	defaultLogConfig := NewDefaultLogConfig()
	ycArray := YamlConfigs{LogConfig: defaultLogConfig}
	err := yml.Load(configPath, &ycArray)
	if err != nil {
		return ycArray, err
	}

	return ycArray, nil
}
