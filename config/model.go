package config

type Tag struct {
	Name string `json:"name"`
}

type Agent struct {
	Interval string `json:"interval"`
}

type Processor struct {
	ProcessorsName string                 `json:"processorsName"`
	Params         map[string]interface{} `json:"params""`
}

type Input struct {
	Tag        Tag                    `json:"tag""`
	InputName  string                 `json:"inputName""`
	Params     map[string]interface{} `json:"params""`
	Processors []Processor            `json:"processors""`
}

type Output struct {
	OutputName string                 `json:"outputName""`
	Params     map[string]interface{} `json:"params""`
}

type Config struct {
	Tag     Tag      `json:"tag""`
	Agent   Agent    `json:"agent"`
	Inputs  []Input  `json:"inputs"`
	Outputs []Output `json:"outputs"`
}
