package model

import (
	"github.com/alexdemen/http-bst/internal/core"
	"github.com/alexdemen/http-bst/internal/core/log"
)

type Tree struct {
	root   *Node
	logger log.Logger
}

func NewIntTree() (tree *Tree) {
	tree = &Tree{}
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
		return nil, core.NewKeyNotFoundError(key)
	}

	return result, nil
}

func findNode(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	var nextNode *Node
	if key < node.Key {
		nextNode = node.LeftNode
	} else if node.Key < key {
		nextNode = node.RightNode
	} else {
		return node
	}

	return findNode(nextNode, key)
}

func (tree *Tree) Delete(key int) {
	if tree.root == nil {
		return
	} else if tree.root.Key == key {
		tree.root = replace(tree.root)
		return
	}

	checkChild(tree.root, key)
}

func checkChild(node *Node, key int) {
	if node.Key > key && node.LeftNode != nil {
		if node.LeftNode.Key == key {
			node.LeftNode = replace(node.LeftNode)
		} else {
			checkChild(node.LeftNode, key)
		}
	} else if node.Key < key && node.RightNode != nil {
		if node.RightNode.Key == key {
			node.RightNode = replace(node.RightNode)
		} else {
			checkChild(node.RightNode, key)
		}
	}
}

func replace(node *Node) *Node {
	if node.RightNode != nil {
		if node.LeftNode != nil {
			add(node.RightNode, node.LeftNode)
		}
		return node.RightNode
	} else if node.LeftNode != nil {
		return node.LeftNode
	}

	return nil
}
