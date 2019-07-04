# Package comparison

[Overview](#overview)  
[Installing](#installing)  
[Index](#index)   

## Overview
comparison allows you to compare two strings, checking if they are similar or not.

## Installing
````
git clone https://github.com/lwahomura/comparison.git
````
Also download bases from https://cloud.mail.ru/public/5DXP/LapZZZniW and add them to ./dictionary/pkg directory 
in your project

## Index

- [func Compare(first, second string) bool](#func-compare)

## func Compare
````
func Compare(first, second string) bool
````  
Compare returns *true* if strings are similar.  
Argument first should be a string in Russian or English.
Argument second should be a string in Russian or English.
### Example
````
package main

import (
	"hub/comparison"
)

func main() {
	a := compare.Compare("testSentence", "tstsentenc") //returns true
	a = comparison.Compare("testSentence", "tssentn") //returns false
}

````