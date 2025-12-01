#!/usr/bin/env python3

point = 50
counter = 0

def turn(dir, count):
    global point
    global counter
    count = int(count)
    while count > 0:
        if dir == "R": # Go up
            point += 1
            if point == 100:
                point = 0

        if dir == "L": # Go down
            point -= 1
            if point == -1:
                point = 99

        # Comment for part 1
        if point == 0:
            counter += 1

        count -= 1

        # Uncomment for part 1
        #    if point == 0:
        #        counter += 1


#with open("test-input") as f:
with open("input") as f:
    for l in f.readlines():
        l = l.strip('\n')
        dir = l[0]
        count = l[1:]
        turn(dir, count)

print(counter)
