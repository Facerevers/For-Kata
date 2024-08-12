/**
* Environment: Go 1.22.2
*
* To install a package, you need to run `go get <package_name>` in the shell.
* This should update the package and dependencies in your `go.mod` and `go.sum` files.
* 
* You can execute tests using the Test command in the "Run" button.
* You can add more commands via ci-config.json file.
*/

package main

import (
	"fmt"
  	"bufio"
  	"os"
  	"strings"
	"strconv"
)
var romeToArab = map[string]int{"I":1, "II":2, "III":3, "IV":4, "V":5, "VI":6, "VII":7, "VIII":8, "IX":9, "X":10}
var arabToRome = map[int]string{1:"I", 2:"II", 3:"III", 4:"IV", 5:"V", 6:"VI", 7:"VII", 8:"VIII", 9:"IX", 10:"X", 20:"XX", 30:"XXX", 40:"XL", 50:"L", 60:"LX", 70:"LXX", 80:"LXXX", 90:"XC", 100:"C"}
func arabianToRome(result int)(text string, err error){
	val, ok := arabToRome[result]
	if ok{
		return val, nil
	}else{
		units := result%10
		tens := result - units
		textresult := arabToRome[tens] + arabToRome[units]
		return textresult, nil
	}
}
func findOper(text string)(string, error){
	switch{
	case strings.Contains(text, "+"):
		return "+", nil
	case strings.Contains(text, "-"):
		return "-", nil
	case strings.Contains(text, "*"):
		return "*", nil
	case strings.Contains(text, "/"):
		return "/", nil
	default:
		return "", fmt.Errorf("Не найден оператор!")
	}
}
func isRoman(numbers string)bool{
	if _, err := romeToArab[numbers]; err{
		return true
	}
	return false
}
func getNumbers(text string, operator string)(numA, numB int, roman bool, err error){
	numbers := strings.Split(text, operator)
	firstNumbType := isRoman(numbers[0])
	secondNumbType := isRoman(numbers[1])
	if firstNumbType != secondNumbType{
		return numA, numB, roman, fmt.Errorf("Формат введённых чисел не совпадает, или числа выходят за границы диапазона допустимых значений!")
	}
	if firstNumbType && secondNumbType{
		roman = true
		numA = romeToArab[numbers[0]]
		numB = romeToArab[numbers[1]]
	}else{
		numA, err = strconv.Atoi(numbers[0])
		if err != nil{return}
		numB, err = strconv.Atoi(numbers[1])
		if err != nil{return}
	}
	if numA < 1 || numA > 10 ||numB < 1 || numB >10{
		return numA, numB, roman, fmt.Errorf("Введённые числа выходят за границы диапазона допустимых значений!")
	}
	return numA, numB, roman, nil
}
func calc(numA, numB int, operator string)(result int, err error){
	switch operator{
	case "+":
		result = numA + numB
	case "-":
		result = numA - numB
	case "*":
		result = numA * numB
	case "/":
		result = numA / numB
	default:
		err=fmt.Errorf(("Допустимый оператор не найден!"))	
	}
	return
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(("Введите выражение, которое необходимо вычислить: "))
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	symbols := strings.Split(text, " ")
	if len(symbols) > 3{
		fmt.Println("Слишком много арифметических действий!")
		return
	}else{
		text = strings.ReplaceAll(text," ", "")
		operator, err := findOper(text)
		if err != nil{panic(err)}
		numA, numB, roman, err := getNumbers(text, operator)
		if err != nil{panic(err)}
		result, err := calc(numA, numB, operator)
		if err != nil{panic(err)}
		if roman{
			if result <= 0{
				panic("Римские цифры не могут быть отрицательными!")
			}
			firstNumb := arabToRome[numA]
			secondNumb := arabToRome[numB]
			romResult, err := arabianToRome(result)
			if err != nil{return}
			fmt.Println(firstNumb, operator, secondNumb, "=", romResult)
		}else{
			fmt.Println(numA, operator, numB, "=", result)
		}
	}
}
