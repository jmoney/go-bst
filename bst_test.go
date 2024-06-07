package bst_test

import (
	"testing"

	"github.com/jmoney/go-bst"
)

func TestBST_Balance(t *testing.T) {
	t.Run("should balance the tree", func(t *testing.T) {
		root := bst.BST{}
		root.Insert(5)
		root.Insert(15)
		root.Insert(2)
		root.Insert(7)
		root.Insert(12)
		root.Insert(18)

		oldRootValue := root.Root.Data
		root.Balance()
		newRootValue := root.Root.Data

		height := root.Height()
		if height > 3 {
			t.Errorf("Expected tree height to be minimized, got %d", height)
		}

		if oldRootValue == newRootValue {
			t.Errorf("Expected root value to be %d, got %d", oldRootValue, newRootValue)
		}
	})
}

func TestBST_Search(t *testing.T) {
	t.Run("Should search for a value in the tree", func(t *testing.T) {
		root := bst.BST{}
		root.Insert(5)
		root.Insert(15)
		root.Insert(2)
		root.Insert(7)
		root.Insert(12)
		root.Insert(18)

		found := root.Search(7)
		if !found {
			t.Errorf("Expected to find value in tree, got %t", found)
		}

		found = root.Search(8)
		if found {
			t.Errorf("Expected to find value in tree, got %t", found)
		}
	})
}

func TestBST_Height(t *testing.T) {
	t.Run("Should insert and return the height of the tree", func(t *testing.T) {
		root := bst.BST{}
		root.Insert(5)
		root.Insert(15)
		root.Insert(2)
		root.Insert(7)
		root.Insert(12)
		root.Insert(18)

		height := root.Height()
		if height != 4 {
			t.Errorf("Expected tree height to be 4, got %d", height)
		}
	})
}
