# [Day 2: Red-Nosed Report](https://adventofcode.com/2024/day/2)

## Understanding the problem

- Find how many reports are safe for the following safety criteria:

  - The levels are either all increasing or all decreasing.
  - Any two adjacent levels differ by at least one and at most three.

- Example:

```log
7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
```

- How many reports are safe ?

## Solution

- **Linear search**: traverse each and every element make sure if it is either increasing or decreasing and find the difference
