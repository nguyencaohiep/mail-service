package utils

import "strings"

func UniqueString(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func OneOfStrings(mainString string, stringList ...string) bool {
	for _, str := range stringList {
		if strings.Compare(mainString, str) == 0 {
			return true
		}
	}

	return false
}
