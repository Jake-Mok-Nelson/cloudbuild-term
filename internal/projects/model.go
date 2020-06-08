package projects

type ProjectItem struct {
	Name string
	ID   string
}

type Projects struct {
	Items []ProjectItem
}

func (projects *Projects) AddItem(item ProjectItem) []ProjectItem {
	projects.Items = append(projects.Items, item)
	return projects.Items
}
