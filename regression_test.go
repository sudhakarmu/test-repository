package main

import "testing"

func TestRegressionGreet(t *testing.T) {
    result := Greet("Alice")
    expected := "Hello, Alice!"

    if result != expected {
        t.Errorf("Greet(\"Alice\") = %s; want %s", result, expected)
    }
}
