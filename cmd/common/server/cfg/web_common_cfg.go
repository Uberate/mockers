package cfg

type WebConfig struct {
	Port string `yaml:"port" json:"port" mapstructure:"port"`
}
