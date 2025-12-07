#!/usr/bin/env python3
import math
input = list()
solution = 0

def get_invalid(num) -> bool:
    if len(num) < 2:
        return False

    # We do not need to check further than halve way
    pivot = math.ceil(len(num)/2)

    for sequence_length in range(1, pivot +1):
        has_seq = True
        if len(num) % sequence_length != 0:
            continue

        seq = num[0:sequence_length]
        for i in range(1, len(num)//sequence_length):
            start = i * sequence_length
            end = start + sequence_length
            if num[start:end] != seq:
                has_seq = False
                break

        if has_seq:
            return True

    return False


#with open("test-input") as f:
with open("input") as f:
    input = f.readline().strip('\n').split(',')

for pair in input:
    a,b = pair.split('-')
    for i in range(int(a), int(b)+1):
        invalid = get_invalid(str(i))
        if invalid:
            solution += i
            print("Invalid: " + str(i))

print(solution)
