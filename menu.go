package main

type Menu struct {
	title        string
	previousPage string
	currentPage  string
	pages        map[string]*MenuPage
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

	m.previousPage = m.currentPage
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
