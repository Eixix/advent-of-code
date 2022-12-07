package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileNode struct {
	name        string
	parent      *FileNode
	isDirectory bool
	size        int64
	children    map[string]*FileNode
}

func main() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rootDirectory := FileNode{"/", nil, true, 0, make(map[string]*FileNode)}
	directory := &rootDirectory

	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Fields(text)
		if words[0] == "$" {
			if words[1] == "cd" && words[2] != "/" {
				if words[2] == ".." {
					directory = directory.parent
				} else {
					dirName := words[2]
					directory = directory.children[dirName]
				}
			} else if words[1] == "ls" {
				// Do nothing
			}
		} else if words[0] == "dir" {
			directory.children[words[1]] = &FileNode{words[1], directory, true, 0, make(map[string]*FileNode)}
		} else {
			size, _ := strconv.ParseInt(words[0], 10, 64)
			directory.children[words[1]] = &FileNode{words[1], directory, false, size, nil}
		}
	}

	// full Tree
	rootDirectory, _ = addSizesToDirs(rootDirectory)

	fmt.Println(addSmallerThan100K(rootDirectory))

	freeSpace := 70000000 - rootDirectory.size
	fmt.Println(findSmallestFolderToDelete(rootDirectory, freeSpace))
}

func addSizesToDirs(rootNode FileNode) (FileNode, int64) {
	var parentSize int64
	parentSize = 0
	for _, node := range rootNode.children {
		if node.isDirectory {
			_, node.size = addSizesToDirs(*node)
		}
		parentSize += node.size
	}
	rootNode.size = parentSize
	return rootNode, parentSize
}

func addSmallerThan100K(rootNode FileNode) int64 {
	var toReturn int64
	toReturn = 0
	for _, node := range rootNode.children {
		if node.isDirectory && node.size <= 100000 {
			toReturn += node.size + addSmallerThan100K(*node)
		} else if node.isDirectory {
			toReturn += addSmallerThan100K(*node)
		}
	}
	return toReturn
}

func findSmallestFolderToDelete(rootNode FileNode, freeSpace int64) FileNode {
	var targetFreeSpace int64
	targetFreeSpace = 30000000
	deltaFreeSpace := targetFreeSpace - freeSpace

	bestDirectory := rootNode

	for _, node := range rootNode.children {
		if node.isDirectory {
			childBestDirectory := findSmallestFolderToDelete(*node, freeSpace)

			if deltaFreeSpace-childBestDirectory.size < deltaFreeSpace-node.size {
				bestDirectory = childBestDirectory
			} else {
				bestDirectory = *node
			}
		}
	}

	return bestDirectory
}
