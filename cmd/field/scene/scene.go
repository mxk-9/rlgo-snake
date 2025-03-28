package scene

type Node struct {
	Child  map[string]*Node
	Object any
}

type Scene struct {
	Nodes      *Node
	UpdateFunc func(s *Scene)
	DrawFunc   func(s *Scene)
}

func NewScene() *Scene {
	return &Scene{
		Nodes: &Node{},
	}
}

// func (s *Scene) Update() {
// 	s.UpdateFunc(s)
// }

// func (s *Scene) Draw() {
// 	s.DrawFunc(s)
// }

// Increases level of depth
func (n *Node) NewNode(nName string, object any) {
	child := make(map[string]*Node)
	child[nName] = &Node{}
	n.Child = child
	n.Object = object
}

// Extends node group
func (n *Node) AddNode(nName string, object any) {
	if n.Child == nil {
		n.Child = make(map[string]*Node)
	}

	n.Child[nName] = &Node{
		Object: object,
	}
	return
}
