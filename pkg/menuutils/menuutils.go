package menuutils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/buger/goterm"
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

func getNumber(prompt string, reader *bufio.Reader) (int, error) {
	str, err := getResponse(prompt, reader)
	if err != nil {
		return 0, err
	}

	result, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("error converting user answer to intager: %w", err)
	}

	return result, nil
}

func GetNumber(prompt string) (int, error) {
	return getNumber(prompt, bufio.NewReader(os.Stdin))
}

func (u *UtilsCache) GetNumber(prompt string) (int, error) {
	return getNumber(prompt, u.reader)
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

func CenterLine(line string) (result string) {
	for i := 0; i < (goterm.Width()-len(line))/2; i++ {
		result += " "
	}

	return result + line
}
