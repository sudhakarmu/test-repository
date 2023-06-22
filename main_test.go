package main

import "testing"

func TestGreet(t *testing.T) {
    result := Greet("John")
    expected := "Hello, John!"

    if result != expected {
        t.Errorf("Greet(\"John\") = %s; want %s", result, expected)
    }
}
