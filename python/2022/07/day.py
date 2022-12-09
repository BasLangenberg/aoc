#!/usr/bin/env python3
file = "input-tst"
file = "input"

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

dirs = list()
for key in fs.keys():
    dirs.append(key)

dirs.sort()
dirs.reverse()

for dir in dirs:
    for x in dirs:
        if x == dir:
            continue
        if (x.startswith(dir)) and (len(x) > len(dir)):
            if fs[x]["total"] <= 100000:
                fs[dir]["total"] += fs[x]["total"]



total = 0
for dir in fs:
    #    print(fs[dir]["total"])
    if fs[dir]["total"] <= 100000:
        total += fs[dir]["total"]

#for key in fs:
#    print(key, ": ", fs[key])

print(total)

# Too low: 1077499
