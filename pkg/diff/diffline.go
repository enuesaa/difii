package diff

type Diffline struct {
	line  int
	text  string
	added bool // TODO enum
}

func NewDiffline(value Value, added bool) *Diffline {
	return &Diffline{
		line:  value.Line(),
		text:  value.Text(),
		added: added,
	}
}

func (l *Diffline) Line() int {
	return l.line
}

func (v *Diffline) Text() string {
	return v.text
}

func (l *Diffline) Added() bool {
	return l.added
}
