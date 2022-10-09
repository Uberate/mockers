package i18n

type Option func(instance *I18n)

// EnableI18nChange will set the I18n.EnableChange to true. It will lower performance. Default is false.
func EnableI18nChange() Option {
	return func(instance *I18n) {
		if instance != nil {
			instance.EnableChange = true
			//instance.RWLocker = &sync.RWMutex{}
		}
	}
}

func DefaultLanguage(ln string) Option {
	return func(instance *I18n) {
		if instance != nil {
			instance.DefaultLanguage = GetLanguageKey(ln)
		}
	}
}

func NewI18nInstance(options ...Option) *I18n {
	instance := &I18n{
		messages: map[string]map[string]map[LanguageKey]string{},
	}

	for _, item := range options {
		item(instance)
	}

	return instance
}
