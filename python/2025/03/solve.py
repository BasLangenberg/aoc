#!/usr/bin/env python3

sum = 0

#with open("test-input") as f:
with open("input") as f:
    for line in f.readlines():
        joltages = list()
        line = line.strip()
        for char in line:
            joltages.append(int(char))

        first, highest = 0, 0
        for i in range(len(joltages)-1):
            if joltages[i] > highest:
                highest = joltages[i]
                first = i

        sec, highest = 0, 0
        print(joltages[first+1:])
        for i in joltages[first + 1:]:
            if i > highest:
                highest = i

        total = joltages[first] * 10 + highest * 1
        print(str(joltages) + " --- " + str(total))

        sum += total

print("Sum: " + str(sum))
