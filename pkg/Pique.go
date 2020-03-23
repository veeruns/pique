package pique

//Block is the minimum required to run the plugin
type Block struct {
	PluginName   string
	FunctionName string
	Arguments    interface{}
}

//BlockState should be output of the plugin
type BlockState int

const (
	//SUCCEEDED is when a Block Succeeded
	SUCCEEDED BlockState = iota

	//FAILED is when Block fails, manual intervention is required.
	FAILED

	//PENDING is when a block is waiting for something to complete
	PENDING
)

func (b BlockState) String() string {
	states := [...]string{
		"PENDING",
		"SUCCEEDED",
		"FAILED",
	}

	if b < SUCCEEDED || b > FAILED {
		return "Unknown"
	}

	return states[b]
}
