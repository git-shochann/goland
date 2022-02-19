package main

func main() {
	age := []int{18, 22, 56, 39, 42, 62, 27}
	gen := [3]int{}
	dummy := true

	for i := 0; i < len(gen); i++ {
		switch dummy {

		case age[i] > 49:
			gen[2]++
		case age[i] > 29:
			gen[1]++
		default:
			gen[0]++
		}

	}

}
