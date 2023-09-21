package diff

type Value struct {
	line int
	has  bool
	text string
}

func NewValue(line int, has bool, text string) *Value {
	return &Value{
		line,
		has,
		text,
	}
}

func (v *Value) Line() int {
	return v.line
}

// to recognize empty string or null.
func (v *Value) Has() bool {
	return v.has
}

func (v *Value) Text() string {
	return v.text
}
