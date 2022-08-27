package memento

// 备忘录
type Memento struct {
	state string
}

func (m *Memento) getSavedState() string {
	return m.state
}
