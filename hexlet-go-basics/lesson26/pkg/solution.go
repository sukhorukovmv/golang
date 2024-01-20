package pkg

type Parent struct {
	Name     string
	Children []Child
}

type Child struct {
	Name string
	Age  int
}

// BEGIN (write your solution here)
func CopyParent(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}
	var cp Parent
	cp.Name = p.Name
	if len(p.Children) > 0 {
		cp.Children = make([]Child, len(p.Children))
		copy(cp.Children, p.Children)
	}
	return cp
}

// END
func CopyParent2(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}
	children := make([]Child, len(p.Children))
	copy(children, p.Children)
	return Parent{
		Name:     p.Name,
		Children: children,
	}
}
