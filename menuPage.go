package main

type MenuPage struct {
	pageTitle string
	items     []*MenuItem
}

func Page(title string) *MenuPage {
	return &MenuPage{
		pageTitle: title,
		items:     make([]*MenuItem, 0),
	}
}

func (m *MenuPage) Item(item *MenuItem) *MenuPage {
	m.items = append(m.items, item)
	return m
}

func (m *MenuPage) String() string {
	return m.pageTitle
}
