#!/usr/bin/env python3

countp1 = 0
countp2 = 0

def isOk(input):
    ok = False
    sorted = input.copy()
    sorted.sort()
    if input == sorted:
        ok = True
        last = input[0]
        for num in input[1:]:
            if num - last not in (1,2,3):
                ok = False
                break
            last = num

    if ok:
        return 1

    sorted.sort(reverse=True)
    if input == sorted:
        ok = True
        last = input[0]
        for num in input[1:]:
            if last - num not in  (1,2,3):
                ok = False
                break
            last = num

    if ok:
        return 1

    return 0
    

with open("input") as f:
    for line in f:
        origline = list()
        for num in line.split():
            origline.append(int(num))

        lineok = isOk(origline) 

        if lineok:
            countp1 += 1
            countp2 += 1
            continue

        for i in range(len(origline)):
            l = origline.copy()
            del l[i]
            partok = isOk(l) 
            print(l, partok)
            if partok == 1:
                countp2 += 1
                break


print("Part 1: " + str(countp1))
print("Part 2: " + str(countp2))
