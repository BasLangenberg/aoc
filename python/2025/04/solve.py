#!/usr/bin/env python3
from pprint import pprint

grid = list()

#with open("test-input") as f:
with open("input") as f:
    for line in f.readlines():
        row = list()
        for char in line.strip():
            row.append(char)

        grid.append(row)

sum = 0

for y in range(len(grid)):
    for x in range(len(grid[y])):
        if grid[y][x] == "@":
            count = 0
            for i in range(-1, 2):
                for j in range(-1, 2):
                    if i == 0 and j == 0:
                        continue

                    if x + j < 0 or x + j >= len(grid[y]):
                        continue
                    if y + i < 0 or y + i >= len(grid):
                        continue

                    if grid[y + i][x + j] == "@":
                        count += 1

            if count < 4:
                sum += 1


pprint(grid)

print(sum)
