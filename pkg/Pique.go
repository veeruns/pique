package pique

//struct block is the minimum required to run the plugin
type Block struct {
	PluginName   string
	FunctionName string
	Arguments    interface{}
}
type BlockState int

const (
	SUCCEEDED BlockState = iota
	PENDING
	FAILED
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
