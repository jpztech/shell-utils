package markdown

import (
	"github.com/charmbracelet/glamour"
)

func Render(input string) (string, error) {
	rendered, err := glamour.Render(input, "dracula")
	if err != nil {
		return "", err
	}
	return rendered, nil
}