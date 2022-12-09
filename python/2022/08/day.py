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

seen = set()

# Horizontal
for rownum in range(1, len(grid)-1):
    maxsize = int(grid[rownum][0])
    treeline = grid[rownum]
    print("left")
    for treenum in range(1, len(treeline)-1):
        if int(treeline[treenum]) > maxsize:
            print(rownum, treenum, grid[rownum][treenum])
            seen.add((rownum, treenum))
            maxsize = int(treeline[treenum])

    maxsize = int(grid[rownum][-1])
    treenum = len(treeline)-1
    print("right")
    while treenum > 0:
        if int(treeline[treenum]) > maxsize:
            print(rownum, treenum, grid[rownum][treenum])
            seen.add((rownum, treenum))
            maxsize = int(treeline[treenum])
        treenum -= 1

#print(len(seen))
#print(seen)

# Vertical
for colnum in range(1, len(grid[0])-1):
    print(colnum)
    maxsize = int(grid[0][colnum])
    print("MS:", maxsize)
    count = 1
    while count < len(grid)-2:
        print("up")
        if int(grid[count][colnum]) > maxsize:
            print(count, colnum, grid[count][colnum])
            seen.add((count, colnum))
            maxsize = int(grid[count][colnum])

        count += 1

    maxsize = int(grid[-1][colnum])
    count = len(grid) -1
    while count > 0:
        print("bot")
        if int(grid[count][colnum]) > maxsize:
            print(count, colnum, grid[count][colnum])
            seen.add((count, colnum))
            maxsize = int(grid[count][colnum])

        count -= 1
    
# Add edges
count = len(seen)
count += len(grid[0]) * 2
count += (len(grid)-2)  *2


print("----------")

print(len(seen))
print(count)

# 1316 == too low
