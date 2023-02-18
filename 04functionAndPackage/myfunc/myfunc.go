package myfunc

import "errors"

func Cal(n1 float64, n2 float64, n3 string) (float64, error) {

	var res float64
	switch n3 {
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	case "*":
		res = n1 * n2
	case "/":
		res = n1 / n2
	default:
		return 0, errors.New("运算符输入错误")
	}
	return res, nil
}
