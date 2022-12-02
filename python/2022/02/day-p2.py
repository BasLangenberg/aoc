# rock     A X
# paper    B Y
# scissors C Z

strat = {"X": 0, "Y": 1, "Z": 2}
winners = { "A": "B", "B": "C", "C": "A"}
draws = { "A": "A", "B": "B", "C": "C"}
losers = { "A": "C", "B": "A", "C": "B"}
points = { "A": 1, "B": 2, "C": 3}

pairs = []

with open('input') as f:
    for l in f.readlines():
        pairs.append(l.strip().split(" "))

score = 0

for match in pairs:
    # Lose
    if strat[match[1]] == 0:
        score = score + points[losers[match[0]]]
    # Draw
    if strat[match[1]] == 1:
        score = score + 3
        score = score + points[draws[match[0]]]
    # Win
    if strat[match[1]] == 2:
        score = score + 6
        score = score + points[winners[match[0]]]

print(score)
