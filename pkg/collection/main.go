package collection

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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

// DeletePath removes a file or directory from the collection root and
// reloads the collection from disk. The path is relative to the collection root.
func (c *Collection) DeletePath(relPath string) error {
	fullPath := filepath.Join(c.filePath, relPath)
	if err := os.RemoveAll(fullPath); err != nil {
		return fmt.Errorf("failed to delete %s: %w", relPath, err)
	}
	c.LoadCollectionFiles()
	return nil
}

// AddFile creates a new JSON request file at the specified relative path.
// The file contains a default HTTP request structure with the given method and url.
func (c *Collection) AddFile(relPath, method, url string) error {
	if !strings.HasSuffix(relPath, ".json") {
		relPath = relPath + ".json"
	}
	fullPath := filepath.Join(c.filePath, relPath)

	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory for %s: %w", relPath, err)
	}

	requestData := map[string]interface{}{
		"method":  method,
		"url":     url,
		"headers": map[string]string{},
		"body":    "",
	}

	data, err := json.MarshalIndent(requestData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal request data: %w", err)
	}

	if err := os.WriteFile(fullPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", relPath, err)
	}

	c.LoadCollectionFiles()
	return nil
}

// RenameNode renames a file or directory in the collection and reloads.
// oldRelPath is relative to the collection root, newName is just the new name (not a path).
func (c *Collection) RenameNode(oldRelPath, newName string) error {
	oldFullPath := filepath.Join(c.filePath, oldRelPath)
	dir := filepath.Dir(oldFullPath)
	newFullPath := filepath.Join(dir, newName)

	if err := os.Rename(oldFullPath, newFullPath); err != nil {
		return fmt.Errorf("failed to rename %s to %s: %w", oldRelPath, newName, err)
	}

	c.LoadCollectionFiles()
	return nil
}

// GetRootPath returns the root directory path of the collection.
func (c *Collection) GetRootPath() string {
	return c.filePath
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
