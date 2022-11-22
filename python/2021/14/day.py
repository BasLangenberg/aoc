with open("input") as x:
    f = x.readlines()

instructions = {}

polymer = f[0]

for pair in f[2:]:
    i = str.rstrip(pair).split(" -> ")
    instructions[i[0]] = i[1]

for x in range(10):
    newpol = ""
    for y in range(len(polymer)-1):
        seq = polymer[y:y+2]
        if seq in instructions:
            newpol += seq[0] + instructions[seq]
            
        if seq not in instructions:
            newpol += seq[0]

    newpol += polymer[-1]
    polymer = newpol

#print(len(polymer))
uniq = set()

for l in str.strip(f[0]):
    uniq.add(l)

for x in instructions:
    uniq.add(instructions[x])

counts = {}
for let in uniq:
    counts[let] = polymer.count(let)

min = 9999999999999999999999999999999999999999999999
max = 0

for key in counts:
    num = counts[key]
    if num > max:
        max = num

    if num < min:
        min = num

print(max-min)
