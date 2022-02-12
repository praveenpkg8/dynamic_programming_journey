/*
Link:https://leetcode.com/problems/climbing-stairs/

70. Climbing Stairs
Easy

10348

317

Add to List

Share
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

 

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
 

Constraints:

1 <= n <= 45
*/

package main

func climbStairs(n int) int {
    steps := make([]int, n)
    for i := 0; i < n ; i++ {
        steps[i] = -1;
    }
    return topDownClimberStairs(n, steps)
}

func topDownClimberStairs(n int, steps []int) int {
    if n == 0 || n == 1 {
        return 1;
    }
    
    if steps[n - 1] >= 0 {
        return steps[n - 1]
    } else {
        steps[n-1] = topDownClimberStairs(n-1, steps) + topDownClimberStairs(n-2, steps)
        return steps[n-1]
    }
}



