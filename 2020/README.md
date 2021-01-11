# Advent of code 2020 highlights

## Highlights

Puzzles that I particularly enjoyed.

### Day 7

I solved this problem using `depth first search` with `pruning`.

### Day 10 part 2

Part 2 was an interesting `dynammic programming` problem which I solved using a `recursive` top down approach with `memoisation` so that the `overlapping subproblems` just required a lookup.

### Day 13 part 2

My naive solution to part 2 was too slow to run on my input. I needed to consult the internet for a faster way to solve the moduler equations.

### Day 14 

In this puzzle I got to use two things that I rarely use in my day job; `bit manipulation` and `recursion`.

### Day 17

Part 1 was a 3d simulation of `Conway's game of life`. Part 2 asked us to extend this to the 4th dimension and my solution got a bit messy at this point as it wasn't easy to extend.

### Day 18

I found this expression evaluation puzzle particularly challenging. If I didn't use `test driven development` I definitely wouldn't have been able to create a (very messy) solution.

### Day 19

I enjoyed this puzzle as I got to use `recursion` and `regexes` to solve it.

### Day 20

This was an interesting problem as for part one you could quickly find the edge tiles as they would be the only tiles that had only 2 matched edges, all other tiles would have 3 or 4. I used `recursion` and `back tracking` to assemble the tiles in the correct orientation to form the image. Having found the edge tiles first and having their adjacent tiles calculated massively reduced the `search space`.

### Day 22 

This was an interesting puzzle in which I used `queues`, `stacks` and `recursion`.

### Day 24

Pen and paper saved the day as my doodling help me work out the geometry of hexagons.

## Day structure 

```
input
input.sample
main.go
main_test.go
```
