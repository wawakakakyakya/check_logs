package config

func NewConfig() (*YamlConfigs, error) {
	cfgAr, err := LoadYamlConfig()
	if err != nil {
		return nil, err
	}

	return &cfgAr, nil
}
