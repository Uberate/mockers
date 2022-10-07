package cfg

import "github.com/uberate/mockers/cmd/common/server/cfg"

type I18nWebConfig struct {
	WebCfg cfg.WebConfig `json:"web" yaml:"web" mapstructure:"web"`

	// If enable NotFoundWith404, return 404 code when specify message not found. Else return the string in format
	// 'namespace-code'.
	NotFoundWith404 bool `json:"not_found_with_404" yaml:"not_found_with_404" mapstructure:"not_found_with_404"`

	// If specify message not found in specify language, will search message in default language scope. If is emtpy,
	// ignore the default language.
	DefaultLanguage string `json:"default_language" yaml:"default_language" mapstructure:"default_language"`

	// If enable change, all the behavior will add a locker to lock the i18n value. It was lower performance.
	EnableChange bool `json:"enable_change" yaml:"enable_change" mapstructure:"enable_change"`
}
