package model

import (
	"errors"
	"github.com/alexdemen/http-bst/internal/core"
	"testing"
)

func TestTree_Find(t *testing.T) {
	testCases := []struct {
		key int
		err error
	}{
		{key: 100, err: nil},
		{key: 333, err: core.NewKeyNotFoundError(333)},
	}

	tree := NewIntTree()

	for _, val := range []int{100, 50, 30, 300, 200, 150, 250, 400, 350, 500} {
		tree.Add(val)
	}

	for _, testCase := range testCases {
		node, err := tree.Find(testCase.key)
		if !errors.Is(err, testCase.err) && node.Key != testCase.key {
			t.Errorf("error in test case with key %d", testCase.key)
		}
	}
}

func TestTree_Add(t *testing.T) {
	tree := NewIntTree()

	for _, val := range []int{100, 50, 30, 300, 200, 150, 250, 400, 350, 500} {
		tree.Add(val)
	}

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
	tree := NewIntTree()

	for _, val := range []int{100, 50, 30, 300, 200, 150, 250, 400, 350, 500} {
		tree.Add(val)
	}

	tree.Delete(150)

	node, err := tree.Find(150)
	if err == nil || !errors.Is(err, core.NewKeyNotFoundError(150)) {
		t.Fatal("invalid error")
	}

	if node != nil {
		t.Errorf("findNode invalid node")
	}

	tree.Delete(300)

	node, err = tree.Find(300)
	if err == nil || !errors.Is(err, core.NewKeyNotFoundError(300)) {
		t.Fatal("invalid error")
	}

	if node != nil {
		t.Errorf("findNode invalid node")
	}

	tree.Delete(50)

	node, err = tree.Find(50)
	if err == nil || !errors.Is(err, core.NewKeyNotFoundError(50)) {
		t.Fatal("invalid error")
	}

	if node != nil {
		t.Errorf("findNode invalid node")
	}
}

func TestTree_Delete_Root(t *testing.T) {
	tree := NewIntTree()

	for _, val := range []int{100, 50, 30, 300, 200, 150, 250, 400, 350, 500} {
		tree.Add(val)
	}

	tree.Delete(100)

	node, err := tree.Find(100)
	if err == nil || !errors.Is(err, core.NewKeyNotFoundError(100)) {
		t.Fatal("invalid error")
	}

	if node != nil {
		t.Errorf("findNode invalid node")
	}
}

func TestTree_Delete_NodeWithoutChild(t *testing.T) {
	tree := NewIntTree()

	for _, val := range []int{100, 50, 30, 300, 200, 150, 250, 400, 350, 500} {
		tree.Add(val)
	}

	tree.Delete(30)

	node, err := tree.Find(30)
	if err == nil || !errors.Is(err, core.NewKeyNotFoundError(30)) {
		t.Fatal("invalid error")
	}

	if node != nil {
		t.Errorf("findNode invalid node")
	}
}

func TestTree_Delete_NullRoot(t *testing.T) {
	tree := NewIntTree()

	tree.Delete(30)

	node, err := tree.Find(30)
	if err == nil || !errors.Is(err, core.NewKeyNotFoundError(30)) {
		t.Fatal("invalid error")
	}

	if node != nil {
		t.Errorf("findNode invalid node")
	}
}
