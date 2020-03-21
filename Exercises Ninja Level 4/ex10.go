package main

import "fmt"

func main() {
	m := map[string] []string{//mapa com chaves to tipo string e elementos do tipo slice de string
		`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`:[]string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:[]string{ `Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`pegoraro_leticia`]= []string{`Music`,`money`,`courses`,`reading`}
	delete(m,"bond_james")
	delete(m,"no_dr")

	for nome, coisasFavoritas := range m{
		fmt.Printf("As coisas favoritas de %s são: ", nome)
		for _, coisa := range coisasFavoritas{
			fmt.Printf("%s--",coisa)
		}
		fmt.Println()
	}

}
