package main

import "fmt"

func main() {
	age := 13

	ageCategory := func(a int) string {
		categories := map[string]func(int) bool{
			"kid":      func(a int) bool { return a > 0 && a < 10 },
			"teenager": func(a int) bool { return a > 10 && a < 18 },
			"Adult":    func(a int) bool { return a >= 18 },
		}

		for category, check := range categories {
			if check(a) {
				return category
			}
		}
		return "Wrong age"
	}

	fmt.Println(ageCategory(age))
}
