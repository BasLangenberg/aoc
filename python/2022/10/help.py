from collections import defaultdict


def add_value():
    global x, value_to_add
    x += value_to_add
    value_to_add = 0


def hold(value):
    global should_wait, value_to_add
    value_to_add = value
    should_wait = True


def release():
    global should_wait
    should_wait = False


instructions = iter([x.strip() for x in open("input", "r").readlines()])
signal_strengths = defaultdict(int)
value_to_add = 0
x = 1
should_wait = False

for i in range(1, 300):
    if should_wait:
        signal_strengths[i] = x
        release()
    else:
        add_value()
        signal_strengths[i] = x
        instruction = next(instructions, None)
        if instruction == None:
            break
        elif instruction.startswith("addx"):
            hold(int(instruction.split()[1]))


print(
    "Part 1:",
    signal_strengths[20] * 20
    + signal_strengths[60] * 60
    + signal_strengths[100] * 100
    + signal_strengths[140] * 140
    + signal_strengths[180] * 180
    + signal_strengths[220] * 220,
)
print("Part 2:")
for k, v in signal_strengths.items():
    if (k - 1) % 40 == 0:
        print("\n")
    if (k - 1) % 40 in [v - 1, v, v + 1]:
        print("###", end="")
    else:
        print("   ", end="")
