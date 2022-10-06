package cfg

import "github.com/uberate/mockers/cmd/common/server/cfg"

type I18nWebConfig struct {
	WebCfg cfg.WebConfig `json:"web" yaml:"web" mapstructure:"web"`
}
