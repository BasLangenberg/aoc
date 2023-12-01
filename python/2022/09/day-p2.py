#!/usr/bin/env python3

input = 'input-tst'
input = 'input'

knot = (0,0)
rope = ((knot,) * 10)

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
    prevpiece = (10000000000000000000,100000000000000000)
    for line in f.readlines():
        direction, moves = line.strip().split(" ")

        for step in range(int(moves)):
            for piece in rope:
                if is_adjecent(piece, prevpiece):
                    continue

                if direction == "U":
                    piece = (piece[0]+1, piece[1])
                if direction == "D":
                    piece = (piece[0]-1, piece[1])
                if direction == "L":
                    piece = (piece[0], piece[1]-1)
                if direction == "R":
                    piece = (piece[0], piece[1]+1)
    
    
                prevpiece = piece

            visited.add(rope[-1])    

        lastdir = direction
        
print(len(visited))
print(visited)

# too high: 6073
