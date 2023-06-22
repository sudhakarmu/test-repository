package com.example;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class MyCalculatorTest {
    @Test
    public void testAdd() {
        MyCalculator calculator = new MyCalculator();
        int result = calculator.add(2, 3);
        assertEquals(5, result);
    }
}
