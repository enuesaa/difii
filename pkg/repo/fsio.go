package repo

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/c-bata/go-prompt"
	"golang.org/x/term"
)

type FsioInterface interface {
	Printf(format string, a ...any)
	Confirm(message string) bool
	SelectCompareDir() string
	IsDirOrFileExist(path string) bool
	ListFiles(dir string) []string
	ReadStream(path string) *os.File
}

type Fsio struct {
	termState *term.State
}

func NewFsio() *Fsio {
	return &Fsio{}
}
func (fsio *Fsio) Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

// see https://github.com/c-bata/go-prompt/issues/8
// see https://github.com/c-bata/go-prompt/issues/233
func (fsio *Fsio) saveState() {
	state, _ := term.GetState(int(os.Stdin.Fd()))
	fsio.termState = state
}

func (fsio *Fsio) restoreState() {
	if fsio.termState != nil {
		term.Restore(int(os.Stdin.Fd()), fsio.termState)
	}
	fsio.termState = nil
}

func (fsio *Fsio) Confirm(message string) bool {
	fsio.saveState()

	suggestion := func(in prompt.Document) []prompt.Suggest {
		return make([]prompt.Suggest, 0)
	}
	options := make([]prompt.Option, 0)
	options = append(options, prompt.OptionAddKeyBind(prompt.KeyBind{
		Key: prompt.ControlC,
		Fn: func(*prompt.Buffer) {
			fsio.restoreState()
			os.Exit(0)
		},
	}))
	options = append(options, prompt.OptionSuggestionBGColor(prompt.Black))
	options = append(options, prompt.OptionScrollbarThumbColor(prompt.Black))
	options = append(options, prompt.OptionSuggestionTextColor(prompt.White))
	options = append(options, prompt.OptionSelectedSuggestionBGColor(prompt.Black))
	options = append(options, prompt.OptionSelectedSuggestionTextColor(prompt.Cyan))
	options = append(options, prompt.OptionPrefixTextColor(prompt.Brown))

	answer := prompt.Input(message+" (y/N) ", suggestion, options...)
	answer = strings.TrimSpace(answer)
	answer = strings.ToLower(answer)

	if answer == "y" || answer == "Y" {
		fsio.restoreState()
		return true
	}
	fsio.restoreState()
	return false
}

func (fsio *Fsio) SelectCompareDir() string {
	fsio.saveState()

	options := make([]prompt.Option, 0)
	// TODO: 選んでenterを押すと入力途中なのに完了してしまうのをなんとかしたい
	options = append(options, prompt.OptionAddKeyBind(prompt.KeyBind{
		Key: prompt.ControlC,
		Fn: func(*prompt.Buffer) {
			fsio.restoreState()
			os.Exit(0)
		},
	}))
	options = append(options, prompt.OptionShowCompletionAtStart())
	options = append(options, prompt.OptionSuggestionBGColor(prompt.Black))
	options = append(options, prompt.OptionScrollbarThumbColor(prompt.Black))
	options = append(options, prompt.OptionSuggestionTextColor(prompt.White))
	options = append(options, prompt.OptionSelectedSuggestionBGColor(prompt.Black))
	options = append(options, prompt.OptionSelectedSuggestionTextColor(prompt.Cyan))
	options = append(options, prompt.OptionMaxSuggestion(15))
	options = append(options, prompt.OptionPrefixTextColor(prompt.Brown))
	options = append(options, prompt.OptionCompletionOnDown())

	for {
		dir := prompt.Input("Compare dir (--compare): ", fsio.suggestDirs, options...)
		if fsio.IsDirOrFileExist(dir) {
			fsio.restoreState()
			return dir
		}
		fmt.Printf("Dir %s does not exist. \n", dir)
	}
}

func (fsio *Fsio) suggestDirs(in prompt.Document) []prompt.Suggest {
	suggests := make([]prompt.Suggest, 0)
	suggests = append(suggests, prompt.Suggest{Text: "./"})
	suggests = append(suggests, prompt.Suggest{Text: "../"})

	text := in.Text

	searchDir := text
	if !strings.HasSuffix(text, "/") {
		searchDir = filepath.Dir(text)
	}
	basePath := ""
	if strings.Contains(text, "/") {
		basePath = filepath.Dir(text) + "/"
	}

	for _, dir := range fsio.listDirs(searchDir) {
		suggests = append(suggests, prompt.Suggest{Text: basePath + dir})
	}

	if text != "." && !strings.HasSuffix(text, "/") && fsio.IsDirOrFileExist(text) {
		for _, dir := range fsio.listDirs(text) {
			suggests = append(suggests, prompt.Suggest{Text: text + "/" + dir})
		}
	}

	// filter .git
	suggests = slices.DeleteFunc(suggests, func(suggest prompt.Suggest) bool {
		return strings.HasSuffix(suggest.Text, ".git")
	})

	return prompt.FilterHasPrefix(suggests, text, false)
}

//TODO: refactor
func (fsio *Fsio) listDirs(path string) []string {
	dirs := make([]string, 0)
	files, err := os.ReadDir(path)
	if err != nil {
		return dirs
	}
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
		}
	}

	return dirs
}

func (fsio *Fsio) IsDirOrFileExist(path string) bool {
	// see https://gist.github.com/mattes/d13e273314c3b3ade33f
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func (fsio *Fsio) ListFiles(dir string) []string {
	filenames := make([]string, 0)
	filepath.Walk(dir, func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// TODO: move 
		if strings.HasPrefix(path, ".git/") {
			return nil
		}
		if file.IsDir() {
			return nil
		} else {
			filenames = append(filenames, path)
		}
		return nil
	})

	return fsio.removeRelativePathFromFilenames(filenames, dir+"/")
}

func (fsio *Fsio) removeRelativePathFromFilenames(filenames []string, path string) []string {
	converted := make([]string, 0)
	for _, filename := range filenames {
		converted = append(converted, strings.TrimPrefix(filename, path))
	}
	return converted
}

func (fsio *Fsio) ReadStream(path string) *os.File {
	f, _ := os.Open(path)

	return f
}
