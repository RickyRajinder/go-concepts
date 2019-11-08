package main

import "strconv"

//64 bit limited
func addBinary64bit(a, b string) string {
	if len(a) == 0 || a == "0" {
		return b
	} else if len(b) == 0 || b == "0" {
		return a
	}
	first, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		return err.Error()
	}
	second, _ := strconv.ParseInt(b, 2, 64)
	return strconv.FormatInt(first+second, 2)
}


func main() {
	res := addBinary64bit("1010", "1011")
	if res != "10101" {
		println("FAILED")
	}
}
