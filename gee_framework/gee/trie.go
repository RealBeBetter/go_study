package gee

import "strings"

type Node struct {
	Pattern  string  // 待匹配的路由，例如 /p/:lang
	Part     string  // 待匹配路由的一部分，例如 :lang
	Children []*Node // 子节点，例如 :lang 的可选项 [go, java, php] 等
	IsFuzzy  bool    // 是否模糊匹配，part 含有 : 或者 * 时为 true
}

// MatchChild 获取下一层第一个匹配的路由
func (n *Node) MatchChild(part string) *Node {
	for _, child := range n.Children {
		if child.Part == part || child.IsFuzzy {
			return child
		}
	}

	return nil
}

// MatchChildren 匹配下一层中所有匹配的子路由
func (n *Node) MatchChildren(part string) []*Node {
	nodes := make([]*Node, 0)
	for _, child := range n.Children {
		if child.Part == part || child.IsFuzzy {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

func (n *Node) Insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.Pattern = pattern
		return
	}

	part := parts[height]
	child := n.MatchChild(part)
	if child == nil {
		child = &Node{Part: part, IsFuzzy: part[0] == ':' || part[0] == '*'}
		n.Children = append(n.Children, child)
	}

	child.Insert(pattern, parts, height+1)
}

func (n *Node) Search(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(n.Part, "*") {
		if n.Pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.MatchChildren(part)

	for _, child := range children {
		result := child.Search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
