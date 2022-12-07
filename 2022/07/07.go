package main

import (
	"bufio"
	"fmt"
	"math"
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

	freeSpace := 30000000 - (70000000 - rootDirectory.size)
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

func findSmallestFolderToDelete(rootNode FileNode, least int64) int64 {
	var smallest int64
	smallest = math.MaxInt64
	for _, child := range rootNode.children {
		if child.isDirectory {
			smallestChild := findSmallestFolderToDelete(*child, least)
			if smallestChild < smallest {
				smallest = smallestChild
			}
		}
	}
	if rootNode.size >= least && rootNode.size < smallest {
		smallest = rootNode.size
	}

	return smallest
}
