package terminalmenu

// MenuItem represents an item in menu.
type MenuItem struct {
	name     string
	callback func()
}

func item(name string, cb func()) *MenuItem {
	return &MenuItem{
		name:     name,
		callback: cb,
	}
}
