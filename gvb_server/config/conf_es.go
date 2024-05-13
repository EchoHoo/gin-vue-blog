package config

type ES struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (es ES) URL() string {
	// return fmt.Sprintf("%s:%d", es.Host, es.Port)
	return es.Host
}
