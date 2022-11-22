with open("input") as x:
    f = x.readlines()

instructions = {}

polymer = f[0]
counts = {}

for pair in f[2:]:
    i = str.rstrip(pair).split(" -> ")
    instructions[i[0]] = i[1]
    counts[i[0]] = 0

for let in range(len(polymer)-1):
    if polymer[let:let+2] not in counts:
        counts[polymer[let:let+1]] = 0
    if polymer[let:let+2] in counts:
        counts[polymer[let:let+2]] += 1

for _ in range(10):
    newcount = dict()
    for key in counts:
        if key in instructions:
            incr = counts[key]
            print(incr)
            insert = key[0] + instructions[key]
            if insert in counts:
                if insert not in counts:
                    newcount[insert] = 1

                if insert in counts:
                    newcount[insert] = counts[insert] + incr

    print(newcount)

    counts = newcount



min = 9999999999999999999999999999999999999999999999
max = 0

for key in counts:
    num = counts[key]
    if num > max:
        max = num

    if num < min:
        min = num

print(max-min)
