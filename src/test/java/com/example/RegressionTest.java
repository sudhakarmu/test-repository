package com.example;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class RegressionTest {
    @Test
    public void testRegression() {
        MyCalculator calculator = new MyCalculator();
        int result = calculator.add(6, 7);
        assertEquals(13, result);
    }
}
