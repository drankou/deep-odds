package utils

import "testing"

func TestRemoveDuplicities(t *testing.T) {
	testSlice := []string{"one", "two", "two", "three"}
	expected := []string{"one", "two", "three"}

	result := RemoveDuplicities(testSlice)
	for i, val := range result {
		if val != expected[i] {
			t.Fatal("result slice is incorrect")
		}
	}
}

func TestRemoveDuplicitiesInt(t *testing.T) {
	testSlice := []int{1, 2, 2, 3}
	expected := []int{1, 2, 3}

	result := RemoveDuplicitiesInt(testSlice)
	for i, val := range result {
		if val != expected[i] {
			t.Fatal("result slice is incorrect")
		}
	}
}

func TestUpdateStringSlice(t *testing.T) {
	testSlice := []string{"one", "two", "two", "three"}
	expected := []string{"two", "two"}

	result := UpdateStringSlice([]int{0, 3}, testSlice)
	for i, val := range result{
		if val != expected[i]{
			t.Fatal("result slice is incorrect")
		}
	}
}

func TestUpdateIntSlice(t *testing.T) {
	testSlice := []int{1, 2, 2, 3}
	expected := []int{2, 2}

	result := UpdateIntSlice([]int{0, 3}, testSlice)
	for i, val := range result {
		if val != expected[i] {
			t.Fatal("result slice is incorrect")
		}
	}
}
