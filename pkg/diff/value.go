package diff

// status みたいな感じで hunk を持ちたい
type Value struct {
	has bool
	text string
}
func (v *Value) Has() bool {
	return v.has
}
func (v *Value) Text() string {
	return v.text
}
