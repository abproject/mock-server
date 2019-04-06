package rest

type RestConfig struct {
	Global      ControllerConfig   `yaml:"global"`
	Controllers []ControllerConfig `yaml:"endpoints"`
}

func (restConfig *RestConfig) Init() Rest {
	var rest Rest
	rest.Init(*restConfig)
	return rest
}