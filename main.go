package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "Enter the numbers"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var myFlags arrayFlags

func contains(el string) bool {
	opArr := []string{"add", "sub", "div", "mul"}
	for _ , v := range opArr{
		if el == v {
			return true
		}

	}
    return false
}

func main() {
	flag.Var(&myFlags, "i", "The list of number to be calculated")
	op := flag.String("o", "add", "Option to be performed exmaple: add, sub, div, mul")
	flag.Parse()
	
	status := contains(*op)
    if !status {
		fmt.Printf("%s is not a valid operator \n", *op )
		return
	}

	fmt.Println(myFlags)
	fmt.Println(*op)

	for _, v := range myFlags {
		arr := strings.Split(v, ",")
		o := *op

		var res int = 0
		var m int = 1

		for i, v := range arr {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln(err)
			}
			switch {
				case o == "add":
					res += n
				case o == "mul":
					fmt.Println(m, n)
					m = m * n
					res = m
				case o == "div":
					_ = i
					fmt.Println("Not yet implemented")
				case o == "sub":
					res = - res - n 
			}
		}
		fmt.Println(res)
	}
}
