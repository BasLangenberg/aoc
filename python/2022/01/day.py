with open("input") as x:
    init = x.read().split("\n\n")

elves = []
for x in range(len(init)):
    elf = []
    print(init[x])
    elves.append(init[x].strip().split("\n"))
    
maxcal = 0

for e in range(len(elves)):
    cur = 0
    for num in elves[e]:
        cur += int(num)

    if cur > maxcal:
        maxcal = cur

print(maxcal) 
    
totalcal = []
for e in range(len(elves)):
    cur = 0
    for num in elves[e]:
        cur += int(num)

    totalcal.append(cur)

totalcal.sort(reverse=True)

print(totalcal[0] + totalcal[1] + totalcal[2])
