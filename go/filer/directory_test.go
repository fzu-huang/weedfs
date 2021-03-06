package filer

import (
	"os"
	"strings"
	"testing"
)

func TestDirectory(t *testing.T) {
	{
		dm, _ := NewDirectoryManagerInMap("/tmp/dir.log")
		dm.MakeDirectory("/a/b/c")
		dm.MakeDirectory("/a/b/d")
		dm.MakeDirectory("/a/b/e")
		dm.MakeDirectory("/a/b/e/f")
		dm.MakeDirectory("/a/b/e/f/g")
		dm.MoveUnderDirectory("/a/b/e/f/g", "/a/b")
		dm.MakeDirectory("/a/b/g/h/i")
		dm.DeleteDirectory("/a/b/e/f")
		dm.DeleteDirectory("/a/b/e")
		dirNames, _ := dm.ListDirectories("/a/b/e")
		for _, v := range dirNames {
			println("sub1 dir:", v.Name, "id", v.Id)
		}
		dm.logFile.Close()

		var path []string
		printTree(dm.Root, path)

		dm2, e := NewDirectoryManagerInMap("/tmp/dir.log")
		if e != nil {
			println("load error", e.Error())
		}
		if !compare(dm.Root, dm2.Root) {
			t.Fatal("restored dir not the same!")
		}
		printTree(dm2.Root, path)
	}
	if true {
		os.Remove("/tmp/dir.log")
	}
}

func printTree(node *DirectoryEntryInMap, path []string) {
	println(strings.Join(path, "/") + "/" + node.Name)
	path = append(path, node.Name)
	for _, v := range node.SubDirectories {
		printTree(v, path)
	}
}

func compare(root1 *DirectoryEntryInMap, root2 *DirectoryEntryInMap) bool {
	if len(root1.SubDirectories) != len(root2.SubDirectories) {
		return false
	}
	if root1.Name != root2.Name {
		return false
	}
	if root1.Id != root2.Id {
		return false
	}
	if !(root1.Parent == nil && root2.Parent == nil) {
		if root1.Parent.Id != root2.Parent.Id {
			return false
		}
	}
	for k, v := range root1.SubDirectories {
		if !compare(v, root2.SubDirectories[k]) {
			return false
		}
	}
	return true
}
