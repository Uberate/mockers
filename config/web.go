package config

// Web is the web config,
type Web struct {
	// Port specify which port will be listened.
	Port string `json:"port" yaml:"port"`

	// EnableLivenessProbe will auto add a liveness probe. The liveness probe will start when application start. It's
	// different from EnableReadinessProbe, the liveness probe start directly, without any operator.
	EnableLivenessProbe bool `json:"enable_liveness_probe" yaml:"enable_liveness_probe"`

	// EnableReadinessProbe will auto add a readiness probe. The readiness probe will start when gin-server ready. It's
	// different from EnableLivenessProbe, the readiness probe start after all handler register. And ready flag can be
	// changed in anywhere.
	EnableReadinessProbe bool `json:"enable_readiness_probe" yaml:"enable_readiness_probe"`
}
