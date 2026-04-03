package collection

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// FileNode represents a file or directory in the collection tree.
// IsDir indicates whether the node is a directory.
// Open indicates whether a directory node is expanded in the UI.
// Path holds the absolute filesystem path of the node.
// Children contains the child nodes for directory entries.
type FileNode struct {
	Name     string
	IsDir    bool
	Open     bool
	Path     string
	Children []FileNode
}

// Collection manages a directory-based collection of HTTP request files.
// Files holds the current tree of FileNodes loaded from disk.
// filePath is the root directory path for the collection.
type Collection struct {
	Files    []FileNode
	filePath string
}

// NewCollection creates a new Collection rooted at the given directory path.
func NewCollection(filePath string) *Collection {
	return &Collection{filePath: filePath}
}

// AddFolders creates a nested directory structure within the collection root.
// foldersPath uses forward slashes to denote nested directories (e.g., "a/b/c").
// After creating the directories, it reloads the collection from disk.
func (c *Collection) AddFolders(foldersPath string) {
	path := filepath.Join(c.filePath, foldersPath)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Println("something went wrong while trying to create folder:", err)
	}
	c.LoadCollectionFiles()
}

// GetOpenPaths returns a set of directory paths that are currently expanded
// in the file tree. This is used to preserve UI state across reloads.
func (c *Collection) GetOpenPaths() map[string]bool {
	paths := make(map[string]bool)
	var collect func([]FileNode)
	collect = func(nodes []FileNode) {
		for _, n := range nodes {
			if n.IsDir && n.Open {
				paths[n.Path] = true
				collect(n.Children)
			}
		}
	}
	collect(c.Files)
	return paths
}

// RestoreOpenPaths sets Open=true for directory nodes whose paths exist
// in the provided set. This restores the expanded state after a reload.
func (c *Collection) RestoreOpenPaths(openPaths map[string]bool) {
	var restore func([]FileNode)
	restore = func(nodes []FileNode) {
		for i := range nodes {
			if nodes[i].IsDir {
				if openPaths[nodes[i].Path] {
					nodes[i].Open = true
				}
				restore(nodes[i].Children)
			}
		}
	}
	restore(c.Files)
}

// LoadCollectionFiles reads the collection root directory from disk and
// rebuilds the Files tree. All directory nodes start with Open=false.
func (c *Collection) LoadCollectionFiles() {
	nodes, err := buildTree(c.filePath)
	if err != nil {
		fmt.Printf("Erro ao carregar coleção: %v\n", err)
		return
	}

	log.Println(nodes)
	c.Files = nodes
}

// buildTree recursively constructs a FileNode tree from the filesystem
// starting at currentPath. All directories are created with Open=false.
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
			Path:  filepath.Join(currentPath, entry.Name()),
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
