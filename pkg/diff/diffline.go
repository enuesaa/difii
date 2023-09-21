package diff

type Differ int

const (
	Added Differ = iota
	Removed
)

type Diffline struct {
	line   int
	text   string
	differ Differ
}

func NewDiffline(line int, text string, differ Differ) *Diffline {
	return &Diffline{
		line:   line,
		text:   text,
		differ: differ,
	}
}

func (l *Diffline) Line() int {
	return l.line
}

func (v *Diffline) Text() string {
	return v.text
}

func (l *Diffline) Added() bool {
	return l.differ == Added
}
