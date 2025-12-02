#!/usr/bin/env python3
input = list()
solution = 0

def get_invalid(num) -> bool:
    if len(num) % 2 == 0:
        mid = int(len(num)/2)
        if num[mid:] == num[:mid]:
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
