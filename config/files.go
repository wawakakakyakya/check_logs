package config

type FileConfig struct {
	FileName string       `yaml:"fileName"`
	PosFile  string       `yaml:"posFile"`
	MaxLine  int          `yaml:"maxLine"`
	Words    []WordConfig `yaml:"words"`
}

type WordConfig struct {
	Word          string   `yaml:"word"`
	SkipThreshold int      `yaml:"skipThreshold"`
	Subject       string   `yaml:"subject"`
	Recipients    []string `yaml:"recipients"`
}
