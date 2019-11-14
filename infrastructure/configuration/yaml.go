package configuration

// configFormatYaml config in yaml
type configFormatYaml struct {
	Port         string `yaml:"port,omitempty"`
	DbConnection string `yaml:"db"`
}
