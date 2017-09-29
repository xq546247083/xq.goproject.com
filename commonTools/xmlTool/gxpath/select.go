package gxpath

import (
	"xq.goproject.com/commonTools/xmlTool/gxpath/internal/build"
	"xq.goproject.com/commonTools/xmlTool/gxpath/internal/query"
	"xq.goproject.com/commonTools/xmlTool/gxpath/xpath"
)

// NodeIterator holds all matched Node object.
type NodeIterator struct {
	node  xpath.NodeNavigator
	query query.Query
}

// Current returns current node which matched.
func (t *NodeIterator) Current() xpath.NodeNavigator {
	return t.node
}

// MoveNext moves Navigator to the next match node.
func (t *NodeIterator) MoveNext() bool {
	n := t.query.Select(t)
	if n != nil {
		if !t.node.MoveTo(n) {
			t.node = n.Copy()
		}
		return true
	}
	return false
}

// Select selects a node set using the specified XPath expression.
func Select(root xpath.NodeNavigator, expr string) *NodeIterator {
	qy, err := build.Build(expr)
	if err != nil {
		panic(err)
	}
	t := &NodeIterator{query: qy, node: root}
	return t
}
