#!/usr/bin/env python3
import json

input = "input"
input = "input-tst"

pairs = []

with open(input) as f:
    for pair in f.read().strip().split("\n\n"):
        pairs.append((json.loads(pair.split("\n")[0]),json.loads(pair.split("\n")[1])))

def compare_index(left,right):
    print(left,right)
    if type(left) is int and type(right) is int:
        if left < right:
            return 1
        if left == right:
            return -1
    if type(left) is list and type(right) is list:
        for x in range(min(len(left), len(right))):
            if not compare_index(left[x], right[x]):
                return false
    if type(left) is int and type(right) is list:
        return compare_index([left][0], right[0])
    if type(left) is list and type(right) is int:
        return compare_index(left[0], right)

    return 0

def compare(left, right):
    for i in range(min(len(left), len(right))):
        if compare_index(left[i], right[i]) < 1:
            print(f"FALSE: {left}, {right}")
            return False

    if len(left) < len(right):
        return True


count = 1
for c in range(len(pairs)):
    left = pairs[c][0]
    right = pairs[c][1]
    if compare(left,right):
        print(f"pair {c}: OK")
        count += 1
    else:
        print(f"pair {c}: NOK")

print(count)
# Too low: 5617
