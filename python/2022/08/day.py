#!/usr/bin/env python3

#input = "input"
input = "input-tst"

grid = []

with open(input) as file:
    for line in file.readlines():
        ln = list()
        for char in line.strip():
            ln.append(char)

        grid.append(ln)

seen = 0

# WE NEED TO SAVE COORDS

# Horizontal
for rownum in range(1, len(grid)-1):
    maxsize = int(grid[rownum][0])
    treeline = grid[rownum]
    treeline = treeline[1:len(treeline)-1]
    print(treeline)
    for tree in treeline:
        if int(tree) > maxsize:
            print(tree)
            seen += 1
            maxsize = int(tree)

    maxsize = int(grid[rownum][-1])
    treeline.reverse()
    print("max:",maxsize)
    for tree in treeline:
        if int(tree) > maxsize:
            print(tree)
            seen += 1
            maxsize = int(tree)

# Vertical
for colnum in range(1, len(grid[0])-1):
    maxsize = int(grid[colnum][1])
    count = 1
    while count < len(grid):
        if int(grid[colnum][count]) > maxsize:
            seen += 1
            maxsize = int(grid[colnum][count])

    maxsize = int(grid[rownum][-1])
    treeline.reverse()
    print("max:",maxsize)
    for tree in treeline:
        if int(tree) > maxsize:
            print(tree)
            seen += 1
            maxsize = int(tree)
    
print(seen)
