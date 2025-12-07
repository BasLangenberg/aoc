#!/usr/bin/env python3
from pprint import pprint

fresh = 0

#with open("test-input") as f:
with open("input") as f:
    ranges, produce = f.read().split("\n\n")

rangeList = list()

print("Processing ranges...")
for seq in ranges.splitlines():
    r = list()
    tup = (int(seq.split("-")[0]), int(seq.split("-")[1]) + 1)

    rangeList.append(tup)

print("Processing products...")
for product in produce.splitlines():
    p = int(product)
    for r in rangeList:
        if p in range(r[0], r[1]):
            fresh += 1
            break

rangeSet = list()

# NO WORKY
print("Counting total number of fresh products...")
for r in rangeList:
    for s in range(len(rangeSet)):
        if r[0] > rangeSet[s][0] and r[1] < rangeSet[s][1]:
            # Already in range
            break
        if r[0] < rangeSet[s][0] and r[1] < rangeSet[s][1]:
            rangeSet[s] = (r[0], rangeSet[s][1])
            break
        if r[0] > rangeSet[s][0] and r[1] > rangeSet[s][1]:
            rangeSet[s] = (rangeSet[s][0], r[1])
            break

    rangeSet.append((r[0], r[1]))

pprint("P1: " + str(fresh))
pprint("P2: len inp " + str(len(rangeList)) + " len set " +str(len(rangeSet)))
