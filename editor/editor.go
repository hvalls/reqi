package editor

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const DefaultEditor = "vim"

func EditText(content string) (string, error) {
	filename := "./tpl"
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}

	_, err = f.WriteString(content)
	if err != nil {
		err = os.Remove(filename)
		if err != nil {
			return "", err
		}
		return "", err
	}
	f.Close()

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = DefaultEditor
	}

	executable, err := exec.LookPath(editor)
	if err != nil {
		err = os.Remove(filename)
		if err != nil {
			return "", err
		}
		return "", err
	}

	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	newContent, err := ioutil.ReadFile(filename)
	if err != nil {
		err = os.Remove(filename)
		if err != nil {
			return "", err
		}
		return "", err
	}

	err = os.Remove(filename)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(newContent)), err
}
