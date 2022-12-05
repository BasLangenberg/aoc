priorities = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'

with open('input') as f:
    input = f.read().strip().split("\n")

# Part 1
count = 0
for line in input:
    mid = int(len(line) / 2)
    left = line[:mid]
    right = line[mid:]

    dup = ""
    for char in left:
        if char in right:
            dup = char
            
    count = count + priorities.find(dup)+1

print("Part 1:", count)

# Part 2
count = 0
pos = 0

while pos < len(input):
    e1 = input[pos]
    e2 = input[pos+1]
    e3 = input[pos+2]

    es1 = set(e1)
    es2 = set(e2)
    es3 = set(e3)

    for char in es1:
        if char in es2 and char in es3:
            count = count + priorities.find(char)+1

    pos = pos + 3

print("Part 2:", count)
