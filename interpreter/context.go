package main

type context struct {
	contextMap map[string]int
}

func getNewContext() *context {
	newContext := &context{
		contextMap: make(map[string]int, 0),
	}
	return newContext
}

func (c *context) get(element string) int {
	return c.contextMap[element]
}

func (c *context) put(element string, value int) {
	c.contextMap[element] = value
}
