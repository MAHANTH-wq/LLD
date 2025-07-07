package main

import "fmt"

func main() {

	newContext := getNewContext()
	newContext.put("a", 5)
	newContext.put("b", 4)

	terminalExpressionOne := &terminalExpression{
		element: "a",
	}

	terminalExpressionTwo := &terminalExpression{
		element: "b",
	}

	multiplyNonTerminalExpressionOne := &multiplyNonTerminalExpression{
		leftExpression:  terminalExpressionOne,
		rightExpression: terminalExpressionTwo,
	}

	finalAnswer := multiplyNonTerminalExpressionOne.interpret(newContext)

	fmt.Println(finalAnswer)
}
