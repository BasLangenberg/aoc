#!/usr/bin/env python3

input = "input"
input = "input-tst"

pairs = []

def get_bracks(pair):
    bracks = []

    for j in range(len(pair)): 
        pos = []
        stack = []
        for i in range(len(pair[j])):
            if pair[j][i] == "[":
                stack.append(i)
            if pair[j][i] == "]":
                pos.append((stack.pop(), i))

        pos.sort()
        bracks.append(pos)


    return bracks

def compare(pair, brackpos):
    lstring = pair[0]
    rstring = pair[1]

    if len(lstring) < len(rstring):
        count = len(lstring)
        lstringshorter = True
    else:
        count = len(rstring)
        lstringshorter = False

    for i in range(count):
        # Insert missing brackets
        if rstring[i] == "[" and lstring[i] == "[":
            continue
        if rstring[i] == "]" and lstring[i] == "]":
            continue
        if rstring[i] == "[" and lstring[i] != "[":
            lstring = lstring[:i] + "[" + lstring[i:]
            continue
        if rstring[i] == "]" and lstring[i] != "]":
            lstring = lstring[:i] + "]" + lstring[i:]
            continue
        if rstring[i] != "[" and lstring[i] == "[":
            rstring = rstring[:i] + "[" + rstring[i:]
            continue
        if rstring[i] != "]" and lstring[i] == "]":
            rstring = rstring[:i] + "]" + rstring[i:]
            continue

        print(lstring)
        print(rstring)
        if int(lstring[i]) < int(rstring[i]):
            return True

        if int(lstring[i]) > int(rstring[i]):
            return False

    if lstringshorter:
        return True
    else:
        return False

# BUG: This creates 1 0 from 10
with open(input) as f:
    for pair in f.read().strip().split("\n\n"):
        pairs.append((pair.split("\n")[0].replace(",",""),pair.split("\n")[1].replace(",","")))

ind = 0
for i in range(len(pairs)):
    br = get_bracks(pairs)
    if compare(pairs[i], br):
        print(f"{i+1} = True")
        ind += i+1
    else:
        print(f"{i+1} = False")

print(ind)

# Too low: 5617
