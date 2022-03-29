package main

type MenuPage struct {
	pageTitle string
	items     []*MenuItem
	menuCache *Menu
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

func (m *MenuPage) Back() *Menu {
	return m.menuCache
}

func (m *MenuPage) str() string {
	return m.pageTitle
}
