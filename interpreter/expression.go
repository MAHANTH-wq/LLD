package main

type expression interface {
	interpret(*context) int
}

type terminalExpression struct {
	element string
}

func (t *terminalExpression) interpret(context *context) int {
	return context.get(t.element)
}

type multiplyNonTerminalExpression struct {
	leftExpression  expression
	rightExpression expression
}

func (m *multiplyNonTerminalExpression) interpret(context *context) int {
	return m.leftExpression.interpret(context) * m.rightExpression.interpret(context)
}

type addNonTerminalExpression struct {
	leftExpression  expression
	rightExpression expression
}

func (a *addNonTerminalExpression) interpret(context *context) int {

	return a.leftExpression.interpret(context) + a.rightExpression.interpret(context)
}
