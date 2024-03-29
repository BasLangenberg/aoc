#!/usr/bin/env python3

input = 'input-tst'
input = 'input'

begin = (0,0)
hpos = (0,0)
prevhpos = (0,0)
tpos = (0,0)

visited = set()

dirs = {
        "U": "ver",
        "D": "ver",
        "L": "hor",
        "R": "hor",
        "": "nod",
        }

lastdir = ""

def is_adjecent(tpos, hpos):
    if ((tpos[0] == hpos[0]) or (tpos[0] + 1 == hpos[0]) or (tpos[0] -1 == hpos[0])) and ((tpos[1] == hpos[1]) or (tpos[1] + 1 == hpos[1]) or (tpos[1] -1 == hpos[1])):
        return True

    return False

with open(input) as f:
    for line in f.readlines():
        direction, moves = line.strip().split(" ")

        for step in range(int(moves)):
            prevhpos = hpos
            if direction == "U":
                hpos = (hpos[0]+1, hpos[1])
            if direction == "D":
                hpos = (hpos[0]-1, hpos[1])
            if direction == "L":
                hpos = (hpos[0], hpos[1]-1)
            if direction == "R":
                hpos = (hpos[0], hpos[1]+1)

            if is_adjecent(tpos, hpos):
                visited.add(tpos)
                print(hpos, tpos)
                continue

            tpos = prevhpos

            visited.add(tpos)

            print(hpos, tpos)

        lastdir = direction
                
            
        
print(len(visited))
#print(visited)

# too high: 6073
