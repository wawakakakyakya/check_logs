package config

type SMTPConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"userName"`
	PassWord string `yaml:"passWord"`
	Timeout  int    `yaml:"timeout"`
	From     string `yaml:"from"`
}
