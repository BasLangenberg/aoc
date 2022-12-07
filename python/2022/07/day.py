#!/usr/bin/env python3
from collections import Counter

file = "input-tst"
#file = "input"

fs = {
        "/": {}
        }

curdir = "/"
prevdir = ""

# Setup structure
with open(file) as f:
    for line in f.readlines():
        inp = line.strip().split()
        if inp[0] == "$":
            if inp[1] == "cd":
                if inp[2] == "/":
                    continue

                if inp[2] == "..":
                    curdir = prevdir
                    continue

                prevdir = curdir
                curdir += inp[2] + "/"

                fs[curdir] = dict()
                
            continue

        if inp[0] == "dir":
            #         fs[curdir][inp[1]] = dict()
            continue

        fs[curdir][inp[1]] = int(inp[0])


for dir in fs:
    dircount = 0
    for file in fs[dir]:
        dircount += fs[dir][file]

    fs[dir]["total"] = dircount

fs = dict(sorted(fs.items()))


print(fs)

