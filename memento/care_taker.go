package main

type careTaker struct {
	history []*memento
}

func newCareTaker() *careTaker {
	return &careTaker{
		history: make([]*memento, 0),
	}
}

func (c *careTaker) addMemento(m *memento) {
	c.history = append(c.history, m)
}

func (c *careTaker) latestMemento() *memento {
	return c.history[len(c.history)-1]
}
func (c *careTaker) getMemento(index int) *memento {

	if index >= len(c.history) {
		return nil
	}

	return c.history[index]
}
