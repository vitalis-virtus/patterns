package composite

import "fmt"

type File struct {
	name string
}

func (f *File) Search(key string) {
	fmt.Printf("Searching for keyword %s in File %s\n", key, f.name)
}

func NewFile(name string) File {
	return File{
		name: name,
	}
}
