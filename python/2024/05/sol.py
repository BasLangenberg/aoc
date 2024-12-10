#!/usr/bin/env python3
import math

with open('input') as f:
    input = f.read()


rrules, rproblem = input.split('\n\n')

rules = dict()

for rule in rrules.split():
    before, after = rule.split('|')
    if before not in rules:
        rules[before] = list()
    rules[before].append(after)

nok = list()
count = 0
for line in rproblem.split():
    ok = True
    pages = line.split(',')
    for pn in range(len(pages)-1, -1, -1):
        if pages[pn] not in rules:
            continue

        rule = rules[pages[pn]]
        for x in pages[:pn]:
            if x in rule:
                ok = False
                nok.append(pages)
                break

        if nok:
            break

    if ok:
        count += int(pages[math.floor(len(pages)/2)])
        continue

def sort(inplist):
    updated = False
    for num in range(len(inplist)):
        if inplist[num] in rules: 
            for pnum in inplist[:num]:
                if pnum in rules[inplist[num]]:
                    inplist.remove(pnum)
                    inplist.insert(num+1, pnum)
                    updated = True

    return updated


def sortNok(inplist):
    while True:
        check = sort(inplist)
        if check:
            break

    print(inplist)
    return int(inplist[math.floor(len(inplist)/2)])

p2count = 0
for line in nok:
    nmbr = sortNok(line)

    p2count += nmbr


print("p1: " + str(count))
print("p2 too high: 80869")
print("p2 too high: 62163")
print("p2 too low: 4735")
print("p2: " + str(p2count))
