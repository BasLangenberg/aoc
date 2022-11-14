total = 0
with open("input") as x:
    a = 0
    b = 0
    while line := x.readline():
        a = b
        b = line

        if int(b) > int(a):
            total += 1

print("part 1: %s" % total)

series = []
with open("input") as x:
    while line := x.readline():
        series.append(int(line))


count = 0

a = 0
b = 0
c = 0

prev = 99999999999999999999999999

for x in range(len(series) - (len(series) % 3)):
    a = series[x]
    if a + b + c > prev:
        count += 1

    prev = a + b + c

    c = b
    b = a

print(count)

