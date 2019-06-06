package knot

type Frame struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// NewFrame creates a new frame instance
func NewFrame(frameType string, data interface{}) *Frame {
	return &Frame{frameType, data}
}
