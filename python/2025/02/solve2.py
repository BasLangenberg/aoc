#!/usr/bin/env python3
input = list()
solution = 0

def get_invalid(num) -> bool:
    ln = 1
    while ln < len(num)//2 + 1:
        if num[:ln] == num[ln:ln+ln]:
            return True
        ln += 1
    
    return False


with open("test-input") as f:
    #with open("input") as f:
    input = f.readline().strip('\n').split(',')

for pair in input:
    a,b = pair.split('-')
    for i in range(int(a), int(b)+1):
        invalid = get_invalid(str(i))
        if invalid:
            solution += i
            print("Invalid: " + str(i))

print(solution)
