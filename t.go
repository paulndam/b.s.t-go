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

// inserts new tree node in the B.S.T
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

// visits all the nodes in order.
func(tree *BinarySearchTree) InOrderTraverseTree(function func(int)){

	tree.lock.RLock()
	defer tree.lock.RUnlock()

	inOrderTraverseTree(tree.rootNode,function)


}
// traverse the left node, root node and the right node of the B.S.T
func inOrderTraverseTree(treeNode *TreeNode, function func(int)){

	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode,function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode,function)
	}
}

// Visit all tree node with pre-order traversing.
func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)){
	tree.lock.Lock()
	defer tree.lock.Unlock()

	preOrderTraverseTree(tree.rootNode,function)
}

func preOrderTraverseTree(treeNode *TreeNode, function func(int)){

	if treeNode != nil {
		preOrderTraverseTree(treeNode.leftNode,function)
		preOrderTraverseTree(treeNode.rightNode,function)
	}
}

// traverse node in post order.
func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)){
	tree.lock.Lock()
	defer tree.lock.Unlock()

	postOrderTraverseTree(tree.rootNode,function)
}

func  postOrderTraverseTree(treeNode *TreeNode, function func(int)){

	if treeNode != nil {
		postOrderTraverseTree(treeNode.leftNode,function)
		postOrderTraverseTree(treeNode.rightNode,function)
		function(treeNode.value)
	}
}

// return lowest element , checking if val of leftnode is nil.
func (tree *BinarySearchTree) MinNode() *int{
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var treeNode *TreeNode
	treeNode = tree.rootNode

	if treeNode == nil{
		// nil instead of 0
		return(*int)(nil)
	}
	for {
		if treeNode.leftNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.leftNode
	}
}


// return max element , checking if val of leftnode is nil.
func (tree *BinarySearchTree) MaxNode() *int{
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var treeNode *TreeNode
	treeNode = tree.rootNode

	if treeNode == nil{
		// nil instead of 0
		return(*int)(nil)
	}
	for {
		if treeNode.rightNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.rightNode
	}
}

// searches a specific node in the B.S.T 
func (tree *BinarySearchTree) SearchNode(key int) bool {

	tree.lock.Lock()
	defer tree.lock.Unlock()

	return searchNode(tree.rootNode,key)
}

func searchNode(treeNode *TreeNode, key int)bool{

	if treeNode == nil {
		return false
	}

	if key < treeNode.key{
		return searchNode(treeNode.leftNode,key)
	}

	if key > treeNode.key{
		return searchNode(treeNode.rightNode,key)
	}

	return true
}

// removes the element with the key that is been passed in.
func (tree *BinarySearchTree) RemoveNode(key int){
	tree.lock.Lock()
	defer tree.lock.Unlock()

	removeNode(tree.rootNode,key)
}

func removeNode(treeNode *TreeNode,key int) *TreeNode{

	if treeNode == nil{
		return nil
	}

	if key < treeNode.key{
		treeNode.leftNode = removeNode(treeNode.leftNode,key)
		return treeNode
	}

	if key > treeNode.key{
		treeNode.rightNode = removeNode(treeNode.rightNode,key)
		return treeNode
	}

	// key == node.key
	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}

	if treeNode.leftNode == nil{
		treeNode = treeNode.rightNode
		return treeNode
	}

	if treeNode.rightNode == nil{
		treeNode = treeNode.leftNode
		return treeNode
	}

	var leftMostStringNode *TreeNode
	leftMostStringNode = treeNode.rightNode

	for {
		// find smallest value on the right side.
		if leftMostStringNode != nil && leftMostStringNode.leftNode != nil {
			leftMostStringNode = leftMostStringNode.leftNode
		}else{
			break
		}
	}
	treeNode.key,treeNode.value = leftMostStringNode.key,leftMostStringNode.value

	treeNode.rightNode = removeNode(treeNode.rightNode,treeNode.key)

	return treeNode
}

func main(){

	var tree *BinarySearchTree = &BinarySearchTree{}

	tree.InsertElement(8,8)
	tree.InsertElement(3,3)
	tree.InsertElement(10,10)
	tree.InsertElement(1,1)
	tree.InsertElement(6,6)


	tree.String()

	fmt.Println(tree.SearchNode(90))
	fmt.Println(tree.MinNode())
	fmt.Println(tree.MaxNode())

}