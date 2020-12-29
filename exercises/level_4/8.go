package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range m {
		fmt.Println(k)
		for i, iv := range v {
			fmt.Printf("index no %d => value %s\n", i, iv)
		}
	}
}
