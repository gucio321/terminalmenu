package constructor

import (
	"strconv"
)

type MenuPage struct {
	pageTitle    string
	previousPage string
	items        []*MenuItem
	last         *MenuItem
	menuCache    *Menu
}

func page(title string, menuCache *Menu) *MenuPage {
	return &MenuPage{
		pageTitle: title,
		items:     make([]*MenuItem, 0),
		menuCache: menuCache,
	}
}

func (m *MenuPage) Item(name string, callback func()) *MenuPage {
	m.items = append(m.items, item(name, callback))
	return m
}

func (m *MenuPage) Subpage(title string) *MenuPage {
	if _, ok := m.menuCache.pages[title]; ok {
		panic("the page with given name already exists!")
	}

	p := page(title, m.menuCache)
	m.menuCache.pages[title] = p

	m.Item(p.pageTitle, func() {
		m.menuCache.currentPage = p.pageTitle
	})

	p.previousPage = m.pageTitle

	return p
}

func (m *MenuPage) Exit() *Menu {
	m.last = item("Exit", func() {
		m.menuCache.shouldExit = true
	})
	return m.menuCache
}

func (m *MenuPage) Back() *MenuPage {
	m.last = item("Exit", func() {
		m.menuCache.currentPage = m.previousPage
	})
	return m.menuCache.pages[m.previousPage]
}

func (m *MenuPage) str() (entries []string) {
	entries = []string{m.pageTitle}
	for i, item := range m.items {
		entries = append(entries, "\t"+strconv.Itoa(i+1)+") "+item.name)
	}

	entries = append(entries, "\t"+strconv.Itoa(0)+") "+m.last.name)

	return entries
}
