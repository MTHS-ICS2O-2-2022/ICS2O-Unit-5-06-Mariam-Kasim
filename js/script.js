// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Mariam Kasim
// Created on: Mar 2023
// This file contains the JS functions for index.html

'use strict'
/**
 * This function multiplies using only addition and loops
 */
function myButtonClicked() {

  // Input
  const number1 = parseInt(document.getElementById('number1').value)
  const number2 = parseInt(document.getElementById('number2').value)

  // Process
  let counter = 0
  let answer = 0
  while (counter < number2) {
    answer += number1
    counter ++
  }

  // Output
  document.getElementById('answer').innerHTML = number1 + ' x ' + number2 + ' = ' + answer
}