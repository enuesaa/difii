package repository

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/c-bata/go-prompt"
	"golang.org/x/term"
)

type FsioInterface interface {
	Confirm(message string) bool
	SelectDir(message string) string
	IsDirOrFileExist(path string) bool
	ListDirs(path string) []string
	ListFilesRecursively(path string) []string
	ReadStream(path string) *os.File
	RemoveFile(path string) error
	CopyFile(srcPath string, dstPath string) error
}

type Fsio struct {
	termState *term.State
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

	fsio.restoreState()
	return answer == "y" || answer == "Y"
}

func (fsio *Fsio) SelectDir(message string) string {
	fsio.saveState()

	options := make([]prompt.Option, 0)
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
		dir := prompt.Input(message, fsio.suggestDirs, options...)
		if fsio.IsDirOrFileExist(dir) {
			fsio.restoreState()
			return dir
		}
		fmt.Printf("Dir %s does not exist. \n", dir)
	}
}

func (fsio *Fsio) suggestDirs(in prompt.Document) []prompt.Suggest {
	suggests := make([]prompt.Suggest, 0)
	suggests = append(suggests, prompt.Suggest{Text: "../"})
	suggests = append(suggests, prompt.Suggest{Text: "./"})

	text := in.Text

	searchDir := filepath.Dir(text)
	for _, dir := range fsio.ListDirs(searchDir) {
		path := filepath.Join(searchDir, dir)
		suggests = append(suggests, prompt.Suggest{Text: path})
	}
	if fsio.IsDirOrFileExist(text) {
		for _, dir := range fsio.ListDirs(text) {
			path := filepath.Join(text, dir)
			suggests = append(suggests, prompt.Suggest{Text: path})
		}
	}

	// filter .git
	suggests = slices.DeleteFunc(suggests, func(suggest prompt.Suggest) bool {
		return strings.HasSuffix(suggest.Text, ".git")
	})

	return prompt.FilterHasPrefix(suggests, text, false)
}

func (fsio *Fsio) IsDirOrFileExist(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func (fsio *Fsio) ListDirs(path string) []string {
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

func (fsio *Fsio) ListFilesRecursively(path string) []string {
	filenames := make([]string, 0)
	filepath.Walk(path, func(fpath string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if file.IsDir() {
			return nil
		}
		filenames = append(filenames, fpath)
		return nil
	})

	return fsio.removeRelativePathFromFilenames(filenames, path+"/")
}

func (fsio *Fsio) removeRelativePathFromFilenames(filenames []string, path string) []string {
	converted := make([]string, 0)
	if strings.HasPrefix(path, "./") {
		path = strings.TrimPrefix(path, "./")
	}
	for _, filename := range filenames {
		converted = append(converted, strings.TrimPrefix(filename, path))
	}
	return converted
}

func (fsio *Fsio) ReadStream(path string) *os.File {
	f, _ := os.Open(path)

	return f
}

func (fsio *Fsio) RemoveFile(path string) error {
	return os.RemoveAll(path)
}

func (fsio *Fsio) CopyFile(srcPath string, dstPath string) error {
	srcF, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcF.Close()

	dstF, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstF.Close()

	_, err = io.Copy(dstF, srcF)
	return err
}
