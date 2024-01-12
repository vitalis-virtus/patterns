package composite

import "fmt"

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) Search(key string) {
	fmt.Printf("Searching for keyword %s in Folder %s\n", key, f.name)

	for _, c := range f.components {
		c.Search(key)
	}
}

func (f *Folder) AddComponent(c Component) {
	f.components = append(f.components, c)
}

func NewFolder(name string) Folder {
	return Folder{
		name: name,
	}
}
