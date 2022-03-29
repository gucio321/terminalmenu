package main

import (
	"os"
	"strconv"
)

type MenuPage struct {
	pageTitle string
	items     []*MenuItem
	last      *MenuItem
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
	m.last = item("Back", func() {
		if m.menuCache.previousPage == "" {
			os.Exit(0)
		}

		m.menuCache.currentPage = m.menuCache.previousPage
	})
	return m.menuCache
}

func (m *MenuPage) str() (entries []string) {
	entries = []string{m.pageTitle}
	for i, item := range m.items {
		entries = append(entries, "\t"+strconv.Itoa(i+1)+") "+item.name)
	}

	entries = append(entries, "\t"+strconv.Itoa(0)+") "+m.last.name)

	return entries
}
