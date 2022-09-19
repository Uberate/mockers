package index

type DefaultValueFunction func()

type Fields struct {
	FiledAlias         []string
	JsonPath           []string
	IndexableInterface Indexable
	DefaultValue       DefaultValueFunction
}

type FieldIndexConfig struct {
	// EnableHashMapper will create a map to save all value of fields and values. The key must a unique value when this
	// option is enabled.
	EnableHashMapper bool
	EnableEscOrder   bool
	EnableDescOrder  bool
}
