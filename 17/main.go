package main

func main() {
	total := 0
	for i := 1; i < 1000; i++ {
		total += CountCharacters(GetNumberAsWords(i))
	}
	total += 11 //For 1000, I'm lazy
	print(total)

	for i := 1; i < 0; i++ {
		println(GetNumberAsWords(i))
	}
}

var units = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tens = []string{
	"",
	"ten",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

func CountCharacters(in string) int {
	count := 0
	for _, char := range in {
		if char != ' ' && char != '-' {
			count++
		}
	}
	return count
}

func GetNumberAsWords(in int) string {

	if in < 100 {
		return GetNumberBelow100AsWords(in)
	}

	result := units[in/100] + "hundred"
	if in%100 != 0 {
		result += "and" + GetNumberBelow100AsWords(in%100)
	}
	return result
}

func GetNumberBelow100AsWords(in int) string {
	if in < 20 {
		return units[in]
	}

	result := tens[in/10]
	if in%10 != 0 {
		result += units[in%10]
	}
	return result
}
