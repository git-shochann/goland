package lib

type Person struct {
	name string
	age int
}

func Average(s []int) int {
	total := 0
	for _, i:= range s {
	  total += i
  }
	return int(total/len(s))
}