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

func (m *Menu) Page(title string) *MenuPage {
	if _, ok := m.pages[title]; ok {
		panic("the page with given name already exists!")
	}

	p := page(title, m)
	m.pages[title] = p
	if m.currentPage == "" {
		m.currentPage = title
	}

	return p
}

func (m *Menu) String() string {
	return m.title + "\n\t" + m.pages[m.currentPage].str()
}
