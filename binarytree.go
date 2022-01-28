package main

import (
	"fmt"
	"io"
	"os"
)

type Tree interface {
	insert(data int) *Node
	remove(data int) bool
	remove_recursive(data int)
	get_tree_depth() int
	has_element(data int) bool
	print()
	print_tree()
	count_nodes() int
}

type Node struct {
	data        int
	left_child  *Node
	right_child *Node
	parent      *Node
}

type BinaryTree struct {
	root_node *Node
}

func (btree *BinaryTree) insert(data int) *BinaryTree {
	if btree.root_node == nil {
		btree.root_node = &Node{data: data, left_child: nil, right_child: nil, parent: nil}
	} else {
		btree.root_node.insert(data)
	}
	return btree
}

func (node *Node) insert(data int) {
	if node == nil {
		return
	} else if data <= node.data {
		if node.left_child == nil {
			node.left_child = &Node{data: data, left_child: nil, right_child: nil, parent: node}
		} else {
			node.left_child.insert(data)
		}
	} else {
		if node.right_child == nil {
			node.right_child = &Node{data: data, left_child: nil, right_child: nil, parent: node}
		} else {
			node.right_child.insert(data)
		}
	}
}

func (btree *BinaryTree) remove(data int) bool {
	data_node := btree.has_element(data)
	if data_node == nil {
		return false
	}
	var is_right_child bool = false
	parent_node := data_node.parent

	if parent_node != nil && parent_node.right_child == data_node {
		is_right_child = true
	}
	// first scenario is a leaf node
	if data_node.left_child == nil && data_node.right_child == nil {
		data_node.parent = nil
		data_node = nil
	} else if data_node.left_child == nil { // second scenario only right child
		data_node = data_node.right_child
	} else if data_node.right_child == nil { // third scenario is only left child
		data_node = data_node.left_child
	} else { // forth scenario, the node has two children
		var left_sibling_node *Node = data_node.left_child
		data_node = data_node.right_child
		data_node.include_sibling(left_sibling_node)
	}

	if parent_node == nil {
		btree.root_node = data_node
	} else if is_right_child {
		parent_node.right_child = data_node
	} else {
		parent_node.left_child = data_node
	}
	if data_node != nil {
		data_node.parent = parent_node
	}

	//ToDo it would be nice to balance the tree after the deletion

	return true
}

func (node *Node) include_sibling(sibling_node *Node) {
	tmp_node := node
	for tmp_node.left_child != nil {
		tmp_node = tmp_node.left_child
	}
	tmp_node.left_child = sibling_node
	sibling_node.parent = tmp_node
}

func (btree *BinaryTree) has_element(data int) *Node {
	if btree == nil {
		return nil
	} else {
		var tmp_node *Node = btree.root_node
		for tmp_node != nil {
			if tmp_node.data == data {
				return tmp_node
			} else if tmp_node.data > data {
				tmp_node = tmp_node.left_child
			} else {
				tmp_node = tmp_node.right_child
			}
		}
		return nil
	}
}

func (btree *BinaryTree) remove_recursive(data int) {
	if btree != nil {
		btree.root_node = btree.root_node.remove_recursive(data)
	}
}

func (node *Node) remove_recursive(data int) *Node {
	if node == nil {
		return nil
	}
	if data < node.data {
		node.left_child = node.left_child.remove_recursive(data)
	} else if data > node.data {
		node.right_child = node.right_child.remove_recursive(data)
	} else {
		if node.left_child == nil {
			return node.right_child
		} else if node.right_child == nil {
			return node.left_child
		} else {
			node.data = node.right_child.next_min_data()
			node.right_child = node.right_child.remove_recursive(node.data)
		}
	}
	return node
}

func (node *Node) next_min_data() int {
	tmp := node
	for tmp.left_child != nil {
		tmp = node.left_child
	}
	return tmp.data
}

func (btree *BinaryTree) get_tree_depth() int {
	var result = 0
	if btree != nil && btree.root_node != nil {
		result = btree.root_node.max_depth()
	}
	return result

}

func (node *Node) max_depth() int {
	if node == nil {
		return -1
	} else {
		left_child_depth := node.left_child.max_depth()
		right_child_depth := node.right_child.max_depth()

		if left_child_depth > right_child_depth {
			return left_child_depth + 1
		} else {
			return right_child_depth + 1
		}
	}
}

func (btree *BinaryTree) count_nodes() int {
	if btree != nil && btree.root_node != nil {
		return btree.root_node.count_nodes()
	} else {
		return 0
	}
}

func (node *Node) count_nodes() int {
	if node == nil {
		return 0
	}
	return 1 + node.left_child.count_nodes() + node.right_child.count_nodes()

}

func (btree *BinaryTree) print() {
	if btree == nil {
		return
	}
	print(os.Stdout, btree.root_node, 0, 'M')

}

func print(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left_child, ns+2, 'L')
	print(w, node.right_child, ns+2, 'R')
}

func (btree *BinaryTree) print_tree() {
	if btree == nil {
		return
	}
	print_tree(os.Stdout, btree.root_node, btree.count_nodes()*btree.get_tree_depth(), '\n')

}
func print_tree(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%v%c", node.data, ch)
	print_tree(w, node.left_child, ns/2, ' ')
	print_tree(w, node.right_child, ns+ns/2, '\n')
}

func main() {
	fmt.Println("Hello, World!")
	tree := &BinaryTree{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	tree.print()

	tree2 := &BinaryTree{}
	tree2.insert(100).
		insert(50).
		insert(200).
		insert(10).
		insert(60).
		insert(150).
		insert(250)
	tree2.print_tree()
	fmt.Fprintf(os.Stdout, "the tree height is %v\n", tree.get_tree_depth())
	fmt.Fprintf(os.Stdout, "the total number of nodes is %v\n", tree.count_nodes())
}
