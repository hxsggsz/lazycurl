package collection

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type FileNode struct {
	Name     string
	IsDir    bool
	Open     bool
	Children []FileNode
}

type Collection struct {
	Files    []FileNode
	filePath string
}

func NewCollection(filePath string) *Collection {
	return &Collection{filePath: filePath}
}

func (c *Collection) LoadCollectionFiles() {
	log.Println("loading collection files")
	nodes, err := buildTree(c.filePath)
	if err != nil {
		fmt.Printf("Erro ao carregar coleção: %v\n", err)
		return
	}
	c.Files = nodes
}

func buildTree(currentPath string) ([]FileNode, error) {
	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return nil, err
	}

	var nodes []FileNode
	for _, entry := range entries {
		node := FileNode{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
			Open:  false,
		}

		if node.IsDir {
			childPath := filepath.Join(currentPath, entry.Name())
			children, err := buildTree(childPath)
			if err != nil {
				return nil, err
			}
			node.Children = children
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}
