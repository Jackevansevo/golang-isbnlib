package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func isValid(isbn string) (valid bool) {
	isbn = clean(isbn)
	switch len(isbn) {
	case 10:
		return isValidISBN10(isbn)
	case 13:
		return isvalidISBN13(isbn)
	default:
		return
	}
}

func isvalidISBN13(isbn string) (valid bool) {
	if len(isbn) != 13 {
		return
	}
	return true
}

func isValidISBN10(isbn string) (valid bool) {
	if len(isbn) != 10 {
		return
	}
	nums, err := stringToIntArray(isbn)
	if err != nil {
		log.Fatal(err)
	}
	i, s, t := 0, 0, 0
	for i = 0; i < len(isbn); i++ {
		t += nums[i]
		s += t
	}
	return (s % 11) == 0
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func stringToIntArray(s string) ([]int, error) {
	nums := make([]int, 0, len(s))
	for _, l := range strings.Split(s, "") {
		i, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, i)
	}
	return nums, nil
}

func clean(isbn string) string {
	// [TODO] Return an err if isbn contains invalid characters
	r := strings.NewReplacer("X", "10", "-", "")
	return r.Replace(isbn)
}
