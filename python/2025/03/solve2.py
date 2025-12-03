#!/usr/bin/env python3

sum = 0

with open("test-input") as f:
#with open("input") as f:
    for line in f.readlines():
        joltages = list()
        line = line.strip()
        for char in line:
            joltages.append(int(char))

        mask = list("0" * len(joltages))
        high = 9
        while mask.count("1") < 12:
            if high == 0:
                print("ERROR: HIGH IS 0")
                break
            for num in reversed(range(len(joltages))):
                if mask.count("1") >= 12:
                    break
                if joltages[num] == high:
                    mask[num] = "1"

            high -= 1

        parsed = ""
        for i in range(len(mask)):
            if mask[i] == "1":
                parsed += str(joltages[i])

        print(str(joltages) + " --- " + parsed)
        sum += int(parsed)
        

print("Sum: " + str(sum))
