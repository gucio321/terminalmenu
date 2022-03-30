package terminalmenu

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gucio321/go-clear"

	"github.com/gucio321/terminalmenu/pkg/menuutils"
)

// Menu represents main menu base. It is used to control the whole
// menu structure. It is a kind of manager for menu pages.
type Menu struct {
	title       string
	currentPage string
	pages       map[string]*MenuPage
	shouldExit  bool

	utils *menuutils.UtilsCache
	clear bool
}

// Create creates a new Menu instance.
func Create(title string, shouldClear bool) *Menu {
	return &Menu{
		title: title,
		pages: make(map[string]*MenuPage),

		utils: menuutils.Utils(),
		clear: shouldClear,
	}
}

// MainPage adds the main page of menu.
// It should be called once per *Menu instance
// Call (*MenuPage).Exit() to get current (*Menu) instance and finish
// setup by calling Run.
func (m *Menu) MainPage(title string) *MenuPage {
	if _, ok := m.pages[title]; ok {
		panic("the page with given name already exists!")
	}

	p := page(title, m)
	m.pages[title] = p

	m.currentPage = title

	return p
}

// String implements fmt.Stringer and allows you to print the menu.
func (m *Menu) String() string {
	result := m.title + "\n"
	currentPage := m.pages[m.currentPage]

	for _, item := range currentPage.str() {
		result += "\t" + item + "\n"
	}

	return result
}

// Run start the main loop of the menu in goroutine.
// It returns chan for error messages.
// receive from it to paus your goroutine until the menu ends (gets exited).
func (m *Menu) Run() chan error {
	result := make(chan error, 1)

	go func() {
		for !m.shouldExit {
			if m.clear {
				if err := clear.Clear(); err != nil {
					result <- err

					return
				}
			}

			fmt.Println(m)

			text, err := m.utils.GetResponse("What to do?: ")
			if err != nil {
				result <- fmt.Errorf("error reading user action: %w", err)

				return
			}

			text = strings.ReplaceAll(text, "\n", "")
			text = strings.ReplaceAll(text, "\r", "")

			num, err := strconv.Atoi(text)
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

		result <- nil
	}()

	return result
}
