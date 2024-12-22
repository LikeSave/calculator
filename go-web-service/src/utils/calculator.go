package utils

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func Calc(expression string) (float64, error) {
	m := make([]string, 0)
	for _, expmap := range expression {
		m = append(m, string(expmap))
	}

	var nums []float64
	var ops []string

	precedence := func(op string) int {
		if op == "+" || op == "-" {
			return 1
		}
		if op == "*" || op == "/" {
			return 2
		}
		return 0
	}

	applyOp := func() error {
		if len(nums) < 2 || len(ops) == 0 {
			return errors.New("error")
		}
		b, a := nums[len(nums)-1], nums[len(nums)-2]
		nums = nums[:len(nums)-2]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		switch op {
		case "+":
			nums = append(nums, a+b)
		case "-":
			nums = append(nums, a-b)
		case "*":
			nums = append(nums, a*b)
		case "/":
			if b == 0 {
				return errors.New("division by zero")
			}
			nums = append(nums, a/b)
		}
		return nil
	}

	for _, token := range m {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			nums = append(nums, num)
		} else if strings.TrimSpace(token) != "" {
			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(token) {
				if err := applyOp(); err != nil {
					return 0, err
				}
			}
			ops = append(ops, token)
		}
	}

	for len(ops) > 0 {
		if err := applyOp(); err != nil {
			return 0, err
		}
	}

	if len(nums) != 1 {
		return 0, errors.New("invalid expression")
	}
	return nums[0], nil
}