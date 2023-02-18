package input_collection

import "hippo-data-acquisition/commons/queue"

type InputPlugin interface {
	ExeDataAcquisition(dataQueue queue.Queue)
}

var (
	inputCollection = make(map[string]InputPlugin)
)

func Add(name string, input InputPlugin) {
	inputCollection[name] = input
}

func GetInputs() map[string]InputPlugin {
	return inputCollection
}
