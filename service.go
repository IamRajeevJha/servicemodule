package mygomodule

import "fmt"

type Service struct{}

func (s Service) Func1(a string, b int) string {
	return fmt.Sprint(a, b)
}
func (s Service) Func2(a string, b string) string {
	return fmt.Sprint(a, b)
}
