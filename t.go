package main

import (
	"fmt"
	"sync"
)

// Trees.

/*
A tree is non linear D.S, they are used for searching and other cases.

-A binary tree have a node with a maximun of 2 children.

- In a binary search tree node, the left node property values are lesser than that of the right node property value. In other words, any node child that is to the left of binary search tree is less and any node child that is to the right of the binary search tree is greater than.

- The root node is at the zero level of a tree and each child node can be a root node



---------------------------------------.'
BINARY SEARCH TREE.

- A binary search tree have node with properties or attributes.

--> A key integer.
--> A value integer.
--> LeftNode and RightNode which are instances of the TreeNode

- A Binarhy search tree is D.S that allows quick looping, addition and removal of an element.

- It stores the keys in a sorted manner to enable faster look-up.

- Space usage of B.T is O(n)

- Time usage for insert, removal and search is O(log n)


*/

// Representation of B.S.T.

type TreeNode struct{
	key int
	value int
	leftNode *TreeNode
	rightNode *TreeNode

}

type BinarySearchTree struct {
	rootNode *TreeNode
	lock sync.RWMutex
}

// Insert method.
func (tree *BinarySearchTree) InsertElement(key int, val int){
	// the tree is locked first
	tree.lock.Lock()
	// after the fucn is been ran and returned then it will unlock and insert the element.
	defer tree.lock.Unlock()

	var treeNode *TreeNode
	treeNode = &TreeNode{
		key,val,nil,nil,
	}

	// check if there's a root node
	if tree.rootNode == nil {
		tree.rootNode = treeNode
		fmt.Println("---- Root Node is ---->",tree.rootNode)
	}else{
		// if no root node then insert one.
		InsertTreeNode(tree.rootNode,treeNode)
	}


}


func InsertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode){
	// if new tree node to insert is less than the root node. then move it to the left side of the B.S.T
	if newTreeNode.key < rootNode.key {
		// but first, check to see if there's a node in there already.
		if rootNode.leftNode == nil {
			// then set that treenode that is on the left side to be the new tree node.
			rootNode.leftNode = newTreeNode
		}else{
			// other call the insert method and set that newTreeNode to be the new tree node and store it on the left side
			InsertTreeNode(rootNode.leftNode,newTreeNode)
		}
	}else{
		// otherwise if the new node to insert is greater than , the to the opposite of what was done previously
		if rootNode.rightNode == nil {
			rootNode.rightNode = newTreeNode
		}else{
			InsertTreeNode(rootNode.rightNode,newTreeNode)
		}
	}

}

//  string method turn the tree into a string format

func(tree *BinarySearchTree) String(){

	tree.lock.Lock()

	defer tree.lock.Unlock()

	fmt.Println("---------------------------------")
	Stringify(tree.rootNode,0)

	fmt.Println("---------------------------------")
}

// this method prints the tree node recursively.
func Stringify(treeNode *TreeNode, level int){

	if treeNode != nil {
		format := ""

		for i := 0; i < level ; i++ {
			format += " "
		}

		format += "---[ "
		level++

		Stringify(treeNode.leftNode,level)
		fmt.Printf(format+"%d\n",treeNode.key)
		Stringify(treeNode.rightNode,level)


	}
}

func main(){

	var tree *BinarySearchTree = &BinarySearchTree{}

	tree.InsertElement(8,8)
	tree.InsertElement(3,3)
	tree.InsertElement(10,10)
	tree.InsertElement(1,1)
	tree.InsertElement(6,6)

	tree.String()

}