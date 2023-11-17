package prototype

import "fmt"

type Folder struct {
	Children []Node
	Name     string
}

func (f *Folder) Print() {
	fmt.Println("   folder: ", f.Name)

	for _, node := range f.Children {
		node.Print()
	}
}

func (f *Folder) Clone() Node {
	cloneFolder := &Folder{Name: f.Name + "_clone"}

	var tempChildren []Node
	for _, node := range f.Children {
		copyChild := node.Clone()

		tempChildren = append(tempChildren, copyChild)
	}

	cloneFolder.Children = tempChildren

	return cloneFolder
}
