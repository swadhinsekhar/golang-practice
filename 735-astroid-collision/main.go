package main

import "fmt"

/*
735. Asteroid Collision
We are given an array asteroids of integers representing asteroids in a row.

For each asteroid, the absolute value represents its size, and the sign
represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet,
the smaller one will explode. If both are the same size, both will explode.
Two asteroids moving in the same direction will never meet.

Example 1:

Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.
Example 2:

Input: asteroids = [8,-8]
Output: []
Explanation: The 8 and -8 collide exploding each other.
Example 3:

Input: asteroids = [10,2,-5]
Output: [10]
Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.

*/

func asteroidCollision(asteroids []int) []int {
	var result []int // Initialize the result stack to store asteroids that haven't collided

	for i := 0; i < len(asteroids); i++ {
		asteroid := asteroids[i]

		// If the result stack is empty, the current asteroid doesn't collide,
		// or the current asteroid is moving to the right, add it to the result.
		if len(result) == 0 || result[len(result)-1] < 0 || asteroid > 0 {
			result = append(result, asteroid)
		} else {
			top := result[len(result)-1]

			// Check if the current asteroid collides with the top asteroid on the result stack.
			if -asteroid == top {
				result = result[:len(result)-1] // Both asteroids are destroyed.
			} else if -asteroid > top {
				result = result[:len(result)-1] // The current asteroid destroys the top asteroid.
				i--                             // Recheck with the current asteroid.
			}
		}
	}

	return result
}

func main() {
	fmt.Println("astroid-collision")
	astroids := []int{5, 10, -5}
	fmt.Println("5, 10", asteroidCollision(astroids))
	astroids = []int{8, -8}
	fmt.Println("", asteroidCollision(astroids))
	astroids = []int{10, -2, -5}
	fmt.Println("10", asteroidCollision(astroids))
}
