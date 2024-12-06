#!/usr/bin/env python3

grid = list(list())
found = set()
word = "XMAS"

def check_word(dir, x, y):
    loc = list()
    print(grid[y][x])
    if grid[y][x] != word[0]:
        return

    loc.append((x,y))

    for i in range(1,4):
        print(word[i])
        if dir == "lu":
            if x-i < 0 or y-i < 0:
                return
            if grid[y-i][x-i] == word[i]:
                loc.append((x-i,y-i))
                continue
            return
        if dir == "u":
            if y-i < 0:
                return
            if grid[y-i][x] == word[i]:
                loc.append((x,y-i))
                continue
            return
        if dir == "ru":
            if x+i >= len(grid[0][0]) or y-i < 0:
                return
            if grid[y-i][x+i] == word[i]:
                loc.append((x,y-i))
                continue
            return
        if dir == "r":
            if x+i >= len(grid[0][0]) or y < 0:
                return
            if grid[y][x+i] == word[i]:
                loc.append((x+i,y))
                continue
            return
        if dir == "rd":
            if x+i >= len(grid[0][0]) or y+i >= len(grid[0]) :
                return
            if grid[y+i][x+i] == word[i]:
                loc.append((x+i,y+i))
                continue
            return
        if dir == "d":
            if y+i >= len(grid[0]) :
                return
            if grid[y+i][x] == word[i]:
                loc.append((x,y+i))
                continue
            return
        if dir == "ld":
            if y+i >= len(grid[0]) or x-i < 0 :
                return
            if grid[y+i][x-i] == word[i]:
                loc.append((x-i,y+i))
                continue
            return
        if dir == "l":
            if x-i < 0 :
                return
            if grid[y][x-i] == word[i]:
                loc.append((x-i,y))
                continue
            return

    print(loc)
    if len(loc) == 4:
        found.add(tuple(loc))

with open("test-input") as f:
    for line in f:
        l = list()
        for char in line.strip():
            l.append(str(char))
        grid.append(l)

for x in range(len(grid)):
    for y in range(len(grid[0])):
        for dir in ["lu","u", "ru", "r", "rd", "d", "ld", "l"]:
            check_word(dir, x, y)


print("part 1: " + str(len(found)))
print(found)
