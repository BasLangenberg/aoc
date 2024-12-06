#!/usr/bin/env python3
with open("input") as f:
    count = 0
    for line in f:
        origline = list()
        for num in line.split():
            origline.append(int(num))

        # Check if only ascends or decends
        ok = False
        sorted = origline.copy()
        sorted.sort()
        if origline == sorted:
            ok = True
            last = origline[0]
            for num in origline[1:]:
                if num - last not in (1,2,3):
                    ok = False
                    break
                last = num

        if ok:
            count += 1
            continue

        sorted.sort(reverse=True)
        if origline == sorted:
            ok = True
            last = origline[0]
            for num in origline[1:]:
                if last - num not in  (1,2,3):
                    ok = False
                    break
                last = num

        if ok:
            count += 1

print("Part 1: " + str(count))

with open("input") as f:
    count = 0
    for line in f:
        ok = True

        origline = list()
        for num in line.split():
            origline.append(int(num))

        asc = 0
        desc = 0
        same = 0
        
        last = origline[0]

        for num in origline[1:]:
            if last > num:
                desc += 1
            if last < num:
                asc += 1
            if last == num:
                same += 1

            last = num

        if same > 1 or (desc > 1 and asc  > 1) :
            ok = False
            continue

        exception = 0
        last = origline[0]
        for num in origline[1:]:
            if num == last and same == 1:
                exception += 1
                continue
            if num > last and asc == 1:
                exception += 1
                continue
            if num < last and desc == 1:
                exception += 1
                continue

            if abs(num - last) not in (1,2,3):
                exception += 1

            last = num

        if exception > 1:
            ok = False

        if ok:
            count += 1

        print("Line: " + str(origline) + " Asc: " + str(asc) + " Desc: " + str(desc) + " Same: " + str(same) + " Ok: " + str(ok))

# 572 too high
# 606 too high
print("Part 2: " + str(count))
