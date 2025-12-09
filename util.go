package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func pullStatus(path string) (int, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}
	s := strings.TrimSpace(string(bs))
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	if v != 0 {
		return 1, nil
	}
	return 0, nil
}

func updateDo(path string, status byte) error {
	val := "0"
	if status != 0 {
		val = "1"
	}
	return ioutil.WriteFile(path, []byte(val), 0644)
}
