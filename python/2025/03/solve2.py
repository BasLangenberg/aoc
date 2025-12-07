#!/usr/bin/env python3

sum = 0

with open("test-input") as f:
#with open("input") as f:
    for line in f.readlines():
        joltages = list()
        line = line.strip()
        for char in line:
            joltages.append(int(char))

        grouplength = len(joltages) % 12
        last = len(joltages)
        groups = list()
        while last > 0:
            groups.append(last)
            last -= grouplength

        mask = list("0" * len(joltages))

        high = 9
        while mask.count("1") < 12:
            if high == 0:
                high = 9
            for i in groups:
                for num in range(grouplength):
                    if mask.count("1") >= 12:
                        break
                    if joltages[i-grouplength+num] == high:
                        mask[i-grouplength+num] = "1"

            high -= 1

        parsed = ""
        for i in range(len(mask)):
            if mask[i] == "1":
                parsed += str(joltages[i])

        print(str(joltages) + " --- " + parsed)
        print(str(mask))
        sum += int(parsed)
        

print("Sum: " + str(sum))
