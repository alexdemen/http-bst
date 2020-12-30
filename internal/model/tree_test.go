package model

import (
	"errors"
	error2 "github.com/alexdemen/http-bst/internal/core/error"
	"testing"
)

func TestTree_Find(t *testing.T) {
	testCases := []struct {
		key int
		err error
	}{
		{key: 100, err: nil},
		{key: 333, err: error2.NewKeyNotFoundError(333)},
	}

	tree := NewIntTree([]int{100, 50, 30, 300, 200, 150, 250, 400, 350, 500})

	for _, testCase := range testCases {
		node, err := tree.Find(testCase.key)
		if !errors.Is(err, testCase.err) && node.Key != testCase.key {
			t.Errorf("error in test case with key %d", testCase.key)
		}
	}
}

func TestTree_Add(t *testing.T) {
	tree := NewIntTree([]int{100, 50, 30, 300, 200, 150, 250, 400, 350, 500})

	tree.Add(333)

	node, err := tree.Find(333)
	if err != nil {
		t.Fatal("error at findNode node")
	}

	if node.Key != 333 {
		t.Errorf("finded node invalid")
	}
}

func TestTree_Delete(t *testing.T) {
	tree := NewIntTree([]int{100, 50, 30, 300, 200, 150, 250, 400, 350, 500})

	tree.Delete(300)

	node, err := tree.Find(300)
	if err == nil || !errors.Is(err, error2.NewKeyNotFoundError(300)) {
		t.Fatal("invalid error")
	}

	if node != nil {
		t.Errorf("findNode invalid node")
	}
}
