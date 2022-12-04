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

for _ in range(40):
     newcount = dict()
     for key in instructions:
         incr = counts.get(key, 0)
         p1 = key[0] + instructions[key]
         p2 = instructions[key] + key[1]


         if incr != 0:
             #            print(f"{incr} - {p1} - {p2}")
            newcount[p1] = newcount.get(p1,0) + incr
            newcount[p2] = newcount.get(p2,0) + incr

     counts = newcount

totals = {}
for let in counts:
    if let[0] in totals:
        totals[let[0]] = totals[let[0]] + counts[let]
    
    if let[0] not in totals:
        totals[let[0]] = counts[let]

print(totals)

maxlet = max(totals, key=totals.get)
minlet = min(totals, key=totals.get)

# Not sure why we need to do -1
total = totals[maxlet] - totals[minlet] -1
print(f"{totals[maxlet]} ({maxlet}) - {totals[minlet]} ({minlet}) = {total}")

