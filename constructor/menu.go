package constructor

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Menu struct {
	title       string
	currentPage string
	pages       map[string]*MenuPage
	shouldExit  bool

	reader *bufio.Reader
}

func Create(title string) *Menu {
	return &Menu{
		title: title,
		pages: make(map[string]*MenuPage),

		reader: bufio.NewReader(os.Stdin),
	}
}

func (m *Menu) MainPage(title string) *MenuPage {
	if _, ok := m.pages[title]; ok {
		panic("the page with given name already exists!")
	}

	p := page(title, m)
	m.pages[title] = p

	m.currentPage = title

	return p
}

func (m *Menu) String() string {
	result := m.title + "\n"
	currentPage := m.pages[m.currentPage]
	for _, item := range currentPage.str() {
		result += "\t" + item + "\n"
	}

	return result
}

func (m *Menu) Run() error {
	for !m.shouldExit {
		fmt.Println(m)
		fmt.Print("What to do?: ")
		text, err := m.reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("error reading user action: %w", err)
		}

		text = strings.ReplaceAll(text, "\n", "")
		text = strings.ReplaceAll(text, "\r", "")

		num, _ := strconv.Atoi(text)
		if err != nil {
			continue
		}

		if num > len(m.pages[m.currentPage].items) {
			continue
		}

		switch num {
		case 0:
			m.pages[m.currentPage].last.callback()
		default:
			m.pages[m.currentPage].items[num-1].callback()
		}
	}

	return nil
}
