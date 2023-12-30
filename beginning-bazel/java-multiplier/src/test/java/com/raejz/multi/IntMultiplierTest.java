package com.raejz.multi;

import static org.junit.Assert.assertEquals;

import org.junit.Test;

public class IntMultiplierTest {

  @Test
  public void testIntMultiplier() {
    IntMultiplier im = new IntMultiplier(3, 4);
    assertEquals(12, im.getProduct());
  }

  @Test
  public void testIntMultiplier_Failure() {
    IntMultiplier im = new IntMultiplier(4, 5);
    assertEquals(20, im.getProduct());
  }
}