package main

import (
	"testing"
)

func TestRemove(t *testing.T) {
	tree := &BinaryTree{}
	tree.insert(1).
		insert(2).
		insert(3)

	tables := []struct {
		x int
		b bool
	}{
		{1, true},
		{2, true},
		{4, false},
		{3, true},
		{10, false},
	}
	for _, table := range tables {
		total := tree.remove(table.x)
		if total != table.b {
			t.Errorf("Removing of the data %d was incorrect, got: %t, want: %t", table.x, total, table.b)
		}
	}
}

func TestCountNodes(t *testing.T) {
	tree := &BinaryTree{}
	nodes_cnt := tree.count_nodes()
	if nodes_cnt != 0 {
		t.Errorf("Node count is incorrect, got %d, want %d", nodes_cnt, 0)
	}

	tree.insert(1).
		insert(2).
		insert(3)
	nodes_cnt = tree.count_nodes()
	if nodes_cnt != 3 {
		t.Errorf("Node count is incorrect, got %d, want %d", nodes_cnt, 3)
	}

}

func TestRemoveRecursive(t *testing.T) {
	tree := &BinaryTree{}
	tree.insert(1).
		insert(2).
		insert(3)

	tree.remove_recursive(1)
	tree.remove_recursive(2)
	tree.remove_recursive(3)

	if tree.count_nodes() != 0 {
		t.Errorf("expected nodes count to be 0, but found %d", tree.count_nodes())
	}
}

func TestInsert(t *testing.T) {
	tree := &BinaryTree{}
	tree.insert(2).
		insert(1).
		insert(3)
	if tree.root_node.data != 2 {
		t.Errorf("expected nodes data to be 2, but found %d", tree.root_node.data)
	}
	if tree.root_node.left_child.data != 1 {
		t.Errorf("expected left child node's data to be 1, but found %d", tree.root_node.data)
	}
	if tree.root_node.right_child.data != 3 {
		t.Errorf("expected right child node's data to be 3, but found %d", tree.root_node.data)
	}
}

func TestGetTreeDepth(t *testing.T) {
	tree := &BinaryTree{}
	if tree.get_tree_depth() != 0 {
		t.Errorf("expected tree depth to be 0, but found %d", tree.get_tree_depth())
	}
	tree.insert(2).
		insert(1).
		insert(3)
	if tree.get_tree_depth() != 1 {
		t.Errorf("expected tree depth to be 2, but found %d", tree.get_tree_depth())
	}
}

func TestHasElement(t *testing.T) {
	tree := &BinaryTree{}
	if tree.has_element(1) != nil {
		t.Errorf("expected to find nothing on empty tree, but found nothing instead")
	}

	tree.insert(2).
		insert(1).
		insert(3)

	if tree.has_element(1) == nil {
		t.Errorf("expected to find the node holding data 1, but found nothing instead")
	}
	if tree.has_element(10) != nil {
		t.Errorf("expected not to find a node holding data 10, but found something instead")
	}
}
