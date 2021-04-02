package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(input []Record) (*Node, error) {
	nodes := map[int]*Node{}
	sort.Slice(input, func(i, j int) bool { return input[i].ID < input[j].ID })
	for i, record := range input {
		if i != record.ID || record.Parent > record.ID || record.ID == record.Parent && record.ID != 0 {
			return nil, errors.New("error")
		}
		node := &Node{ID: record.ID}
		nodes[record.ID] = node
		if record.ID > 0 {
			parent := nodes[record.Parent]
			parent.Children = append(parent.Children, node)
		}
	}
	return nodes[0], nil
}
