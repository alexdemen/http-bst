package model

import (
	error2 "github.com/alexdemen/http-bst/internal/core/error"
)

type Tree struct {
	root *Node
}

func NewIntTree(keys []int) (tree *Tree) {
	tree = &Tree{}

	for _, key := range keys {
		tree.Add(key)
	}

	return
}

func (tree *Tree) Add(key int) {
	if tree.root == nil {
		tree.root = &Node{Key: key}
	} else {
		add(tree.root, &Node{Key: key})
	}
}

func add(node *Node, addNode *Node) {
	if addNode.Key < node.Key {
		if node.LeftNode == nil {
			node.LeftNode = addNode
		} else {
			add(node.LeftNode, addNode)
		}
	}

	if node.Key < addNode.Key {
		if node.RightNode == nil {
			node.RightNode = addNode
		} else {
			add(node.RightNode, addNode)
		}
	}
}

func (tree *Tree) Find(key int) (*Node, error) {
	result := findNode(tree.root, key)
	if result == nil {
		return nil, error2.NewKeyNotFoundError(key)
	}

	return result, nil
}

func findNode(node *Node, key int) *Node {
	var nextNode *Node
	if key < node.Key {
		nextNode = node.LeftNode
	} else if node.Key < key {
		nextNode = node.RightNode
	} else {
		return node
	}

	if nextNode == nil {
		return nil
	}

	return findNode(nextNode, key)
}

func (tree *Tree) Delete(key int) {
	node := findNode(tree.root, key)
	if node != nil {
		remove(node)
	}
}

func remove(rmNode *Node) {
	var transferNode *Node
	if rmNode.RightNode != nil {
		transferNode = rmNode.RightNode
		if rmNode.LeftNode != nil {
			add(transferNode, rmNode.LeftNode)
		}
	} else if rmNode.LeftNode != nil {
		transferNode = rmNode.LeftNode
	}

	if transferNode != nil {
		*rmNode = *transferNode
	} else {
		rmNode = nil
	}
}
