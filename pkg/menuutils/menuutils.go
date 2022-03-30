package menuutils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type UtilsCache struct {
	reader *bufio.Reader
}

func Utils() *UtilsCache {
	return &UtilsCache{
		reader: bufio.NewReader(os.Stdin),
	}
}

func getResponse(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")
	return text, nil
}

func GetResponse(prompt string) (string, error) {
	return getResponse(prompt, bufio.NewReader(os.Stdin))
}

func (u *UtilsCache) GetResponse(prompt string) (string, error) {
	return getResponse(prompt, u.reader)
}

func promptEnter(prompt string, reader *bufio.Reader) error {
	fmt.Print(prompt)
	_, err := reader.ReadString('\n')
	return err
}

func PromptEnter(prompt string) error {
	return promptEnter(prompt, bufio.NewReader(os.Stdin))
}

func (u *UtilsCache) PressEnter(prompt string) error {
	return promptEnter(prompt, u.reader)
}
