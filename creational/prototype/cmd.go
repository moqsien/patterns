package prototype

import "fmt"

/*
原型是一种创建型设计模式， 使你能够复制对象，
甚至是复杂对象， 而又无需使代码依赖它们所属的类。

所有的原型类都必须有一个通用的接口， 使得即使在对象所属的具体类未知的情况下也能复制对象。
原型对象可以生成自身的完整副本， 因为相同类的对象可以相互访问对方的私有成员变量。
*/

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}
