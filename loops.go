package main

func main() {

	// loops
	for i := 0; i < 5; i++ {
		println(i)
	}

	// similar to while-loop
	i := 0
	for {
		i++
		println(i)
		if i > 5 {
			break
		}
	}

	// how to loop over collections
	s := []string{"foo", "bar", "buz"}

	for idx, v := range s {
		println(idx, v)
	}

	// loop values in map
	m := make(map[string]string)
	m["first"] = "foo"
	m["second"] = "bar"
	m["third"] = "buz"

	for k, v := range m {
		println(k, v)
	}
}
