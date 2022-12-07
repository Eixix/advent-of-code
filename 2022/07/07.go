package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	first()
}

type FileNode struct {
	name        string
	parent      *FileNode
	isDirectory bool
	size        int
	children    map[string]*FileNode
}

func first() {
	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rootDirectory := FileNode{"/", nil, true, 0, make(map[string]*FileNode)}
	directory := &rootDirectory

	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Fields(text)
		if words[0] == "$" {
			if words[1] == "cd" {
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
			size, _ := strconv.Atoi(words[0])
			directory.children[words[1]] = &FileNode{words[1], directory, false, size, nil}
		}
	}

	toReturn, size := addSizesToDirs(rootDirectory)

	fmt.Println(toReturn.children["a"].size, size)
}

func addSizesToDirs(rootNode FileNode) (FileNode, int) {
	parentSize := 0
	for _, node := range rootNode.children {
		if node.isDirectory {
			_, node.size = addSizesToDirs(*node)
		}
		parentSize += node.size
	}
	rootNode.size = parentSize
	return rootNode, parentSize
}
