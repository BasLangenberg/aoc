#!/usr/bin/env python3
from pprint import pprint
sum = 0
grid = list()
#with open("test-input") as f:
with open("input") as f:
    for line in f.readlines():
        grid.append(list(line.strip("\n")))

start  = grid[0].index("S")
grid[1][start] = '|'

for i in range(2, len(grid)):
    for j in range(len(grid[i])):
        if grid[i][j] == "." and grid[i-1][j] == "|":
            grid[i][j] = "|"
        if grid[i][j] == "^" and grid[i-1][j] == "|":
            sum += 1
            grid[i][j-1] = "|"
            grid[i][j+1] = "|"

pprint(grid)
pprint(sum)

