# rock     A X
# paper    B Y
# scissors C Z

winners = { "A": "Y", "B": "Z", "C": "X"}
draws = { "A": "X", "B": "Y", "C": "Z"}
losers = { "A": "Z", "B": "X", "C": "Y"}
points = { "X": 1, "Y": 2, "Z": 3}
pairs = []

with open('input') as f:
    for l in f.readlines():
        pairs.append(l.strip().split(" "))

score = 0

for match in pairs:
    print(f"inputs: {match}")
    if match[1] == winners[match[0]]:
        print("win!")
        score = score + 6
        score = score + points[match[1]]
        continue

    if match[1] == draws[match[0]]:
        print("draw!")
        score = score + 3
        score = score + points[match[1]]
        continue

    print("loss")
    score = score + points[match[1]]



print(score)
