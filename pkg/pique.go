package pique

//Nodeblock is the minimum required to run the plugin
type Nodeblock struct {
	PluginName   string
	FunctionName string
	Arguments    string
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
