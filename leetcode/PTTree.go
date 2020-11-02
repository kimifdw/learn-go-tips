package main

// PTNode 结点
type PTNode struct {
	// 结点数据
	data interface{}
	// 双亲结点
	parent int32
}

// PTTree 树结构
type PTTree struct {
	// 结点数组
	nodes []PTNode
	// 根的位置和结点数
	r, n int32
}

// CTNode 孩子结点
type CTNode struct {
	child int32
	next  *CTNode
}

// CTBox 表头结构
type CTBox struct {
	data       interface{}
	firstChild *CTNode
}

// CTTree 树
type CTTree struct {
	r, n  int32
	nodes []CTBox
}
