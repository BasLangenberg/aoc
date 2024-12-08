#!/usr/bin/env python3
import math
import sys


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

    if ok:
        count += int(pages[math.floor(len(pages)/2)])
        continue


def sortNok(l):
    ok = True
    counter = 0
    while True: 
        newl = list()
        if counter > 99:
            return -1

        for pn in range(len(l)-1, -1, -1):
            templ = list()
            templ.append(l[pn])
            if l[pn] not in rules:
                continue
    
            rule = rules[l[pn]]
            for x in l[:pn]:
                if x in rule:
                    ok = False
                    templ.append(x)
                    l.remove(x)

            newl = newl + templ

        if ok:
            return int(pages[math.floor(len(l)/2)])

        if not ok:
            l = newl
            counter += 1


   

p2count = 0
for line in nok:
    nmbr = sortNok(line)
    if nmbr < 0:
        print("You wrote an endless loop moron")
        sys.exit(1)

    print(nmbr)
    p2count += nmbr


print("p1: " + str(count))
print("p2 too high: 80869")
print("p2: " + str(p2count))
