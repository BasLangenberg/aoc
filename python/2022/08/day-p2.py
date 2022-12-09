#!/usr/bin/env python3

input = "input"
#input = "input-tst"

grid = []

with open(input) as file:
    for line in file.readlines():
        ln = list()
        for char in line.strip():
            ln.append(char)

        grid.append(ln)

best = 0

for rownum in range(len(grid)):
    for colnum in range(len(grid[rownum])):
        tsize = grid[rownum][colnum]
        if (colnum == 0) or (colnum == len(grid[0]) -1) or (rownum == 0) or (rownum == len(grid) -1):
            continue
        # Up
        UpLen = 1
        while rownum - UpLen > 0:
            if grid[rownum-UpLen][colnum] >= tsize:
                break

            UpLen += 1

        # Down
        DownLen = 1
        while rownum + DownLen < len(grid)-1:
            if grid[rownum+DownLen][colnum] >= tsize:
                break

            DownLen += 1

        # Left
        LeftLen = 1
        while colnum - LeftLen > 0:
            if grid[rownum][colnum-LeftLen] >= tsize:
                break

            LeftLen += 1

        # Right
        RightLen = 1
        while colnum + RightLen < len(grid[rownum])-1:
            if grid[rownum][colnum+RightLen] >= tsize:
                break

            RightLen += 1

        curscore = abs(UpLen) * abs(DownLen) * abs(LeftLen) * abs(RightLen)
        if curscore > best:
            print(f"tree: {rownum},{colnum}")
            print(f"{UpLen} * {DownLen} * {LeftLen} * {RightLen} == {curscore}")
            best = curscore

print(best)

# too low: 252000
