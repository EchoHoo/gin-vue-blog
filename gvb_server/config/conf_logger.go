package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`      //是否打印行号
	LogInConsole bool   `yaml:"log_in_console"` //是否打印路径
}
