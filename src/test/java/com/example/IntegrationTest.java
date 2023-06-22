package com.example;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class IntegrationTest {
    @Test
    public void testIntegration() {
        MyCalculator calculator = new MyCalculator();
        int result = calculator.add(4, 5);
        assertEquals(9, result);
    }
}
