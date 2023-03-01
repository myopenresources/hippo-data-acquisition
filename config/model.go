package config

type AgentConfig struct {
	Spec string `json:"spec"`
}

type ProcessorConfig struct {
	ProcessorsName string                 `json:"processorsName"`
	Params         map[string]interface{} `json:"params""`
}

type InputConfig struct {
	Tag        map[string]string      `json:"tag""`
	InputName  string                 `json:"inputName""`
	Params     map[string]interface{} `json:"params""`
	Processors []ProcessorConfig      `json:"processors""`
}

type OutputConfig struct {
	OutputName string                 `json:"outputName""`
	Params     map[string]interface{} `json:"params""`
}

type DataAcquisitionConfig struct {
	Tag        map[string]string `json:"tag""`
	Agent      AgentConfig       `json:"agent"`
	Inputs     []InputConfig     `json:"inputs"`
	Processors []ProcessorConfig `json:"processors"`
	Outputs    []OutputConfig    `json:"outputs"`
	LogPath    string            `json:"logPath"`
}
