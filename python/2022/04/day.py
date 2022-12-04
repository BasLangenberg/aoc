#!/usr/bin/env python3
pairs = list()

def contains(pair):
    p0 = True
    p1 = True

    for x in pair[0]:
        if x not in pair[1]:
            p0 = False
            continue

    for x in pair[1]:
        if x not in pair[0]:
            p1 = False
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
        

