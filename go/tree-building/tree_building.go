package tree

import (
	"errors"
	"sort"
)

// Record stores ID for a record and its parent
type Record struct {
	ID, Parent int
}

// Node stores Record like a tree
type Node struct {
	ID       int
	Children []*Node
}

// Build a tree of Node from a list of Record
func Build(records []Record) (*Node, error) {
	maxRecord := len(records)
	if maxRecord == 0 {
		return nil, nil
	}
	// TODO: Is it better to sort a large array here or a smaller arrays later
	// sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	// parent map to check duplicates
	parent := make(map[int]int)
	// node map to keep refs fo nodes allocated
	nodeMap := make(map[int]*Node)
	idEqualsParent := 0
	rootID := 0
	for _, record := range records {
		if record.ID >= maxRecord || record.ID < 0 || record.Parent > record.ID {
			return nil, errors.New("Assumptions Failed")
		}
		if record.Parent == record.ID {
			if idEqualsParent >= 1 {
				return nil, errors.New("Assumptions Failed : 2 nodes with (id=parent)")
			}
			rootID = record.ID
			idEqualsParent++
		}
		if _, ok := parent[record.ID]; ok {
			return nil, errors.New("Assumption Failed : Duplicate ID")
		}
		parent[record.ID] = record.Parent

		if _, ok := nodeMap[record.ID]; !ok {
			nodeMap[record.ID] = &Node{ID: record.ID}
		}
		if _, ok := nodeMap[record.Parent]; !ok {
			nodeMap[record.Parent] = &Node{ID: record.Parent, Children: []*Node{}}
		}
		// don't insert loop for root
		if record.ID != record.Parent {
			nodeMap[record.Parent].Children = append(nodeMap[record.Parent].Children, nodeMap[record.ID])
		}
	}

	for _, node := range nodeMap {
		if len(node.Children) > 1 {
			sort.Slice(node.Children, func(i, j int) bool { return node.Children[i].ID < node.Children[j].ID })
		}
	}

	return nodeMap[rootID], nil
}
