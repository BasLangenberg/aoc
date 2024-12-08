#!/usr/bin/env python3

left = list()
right = list()

with open("test-input") as f:
    for l in f.readlines():
        l = l.replace("\n", "")
        x = l.split("   ")
        left.append(int(x[0]))
        right.append(int(x[1]))

left.sort()
right.sort()

sum = 0
for i in range(len(left)):
    sum += abs(int(left[i]) - int(right[i]))

print(f"part 1: %s" % sum)

sum = 0
left = list()
right = list()


with open("input") as f:
    for l in f.readlines():
        l = l.replace("\n", "")
        x = l.split("   ")
        left.append(int(x[0]))
        right.append(int(x[1]))

for i in left:
    count = 0
    for l in right:
        if l == i:
            count += 1

    print(f"{i} * {count} = {i*count}")
    sum += i * count


print(f"part 2: %s" % sum)


for i in zip(left, right):
    print(i)
