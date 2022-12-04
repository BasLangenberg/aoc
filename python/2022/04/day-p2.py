#!/usr/bin/env python3
pairs = list()

def contains(pair):
    p0 = False
    p1 = False

    for x in pair[0]:
        if x in pair[1]:
            p0 = True
            continue

    for x in pair[1]:
        if x in pair[0]:
            p1 = True
            continue

    if p0 or p1:
        return True

    return False

            

with open("input") as file:
    for line in file:
        pair = list()
        for r in line.split(','):
            p = list()
            ppp = r.split('-')
            for num in range(int(ppp[0]), int(ppp[1])+1):
                p.append(num)

            pair.append(p)

        pairs.append(pair)

count = 0

for pair in pairs:
    if contains(pair):
        count = count + 1


print(count)
        

