package main

type Menu struct {
	title       string
	currentPage string
	pages       map[string]*MenuPage
}

func Create(title string) *Menu {
	return &Menu{
		title: title,
		pages: make(map[string]*MenuPage),
	}
}

func (m *Menu) Page(name string, page *MenuPage) *Menu {
	if _, ok := m.pages[name]; ok {
		panic("the page with given name already exists!")
	}

	m.pages[name] = page
	if m.currentPage == "" {
		m.currentPage = name
	}
	return m
}

func (m *Menu) String() string {
	return m.title + "\n\t" + m.pages[m.currentPage].String()
}
