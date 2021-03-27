package main

type Plugin struct{}

func (p Plugin) Process(parameterAccessor func(string) (interface{}, bool)) *map[string]interface{} {
	firstOperand, isFirstOperandAvailable := parameterAccessor("Mul1")
	secondOperand, isSecondOperandAvailable := parameterAccessor("Mul2")

	resultMap := make(map[string]interface{})

	if isFirstOperandAvailable && isSecondOperandAvailable {
		intOp1 := firstOperand.(int32)
		intOp2 := secondOperand.(int32)

		resultMap["MultiplicationResult"] = intOp1 * intOp2
	}

	return &resultMap
}

var Implementation Plugin
