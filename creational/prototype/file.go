package prototype

import "fmt"

type File struct {
	Name string
}

func (f *File) Print() {
	fmt.Println("      file: ", f.Name)
}

func (f *File) Clone() Node {
	return &File{Name: f.Name + "_clone"}
}
