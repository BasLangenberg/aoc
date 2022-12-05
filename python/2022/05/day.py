import queue

with open('input') as f:
    input = f.read().split("\n\n")

stacks = input[0]
instructions = input[1]

stacks = stacks.replace("    ", "[$]")
stacks = stacks.replace("    ", "[$]")
stacks = stacks.replace(" ", "")
stacks = stacks.replace("[", "")
stacks = stacks.replace("]", "")

stacks = stacks.split("\n")

stacknum = stacks[len(stacks)-1]
stacks = stacks[:len(stacks)-1]

print(stacks)
print(stacknum[len(stacknum)-1])

ship = dict()

for x in range(len(stacknum)):
    ship[x] = queue.LifoQueue()

s = len(stacks) -1

while s > -1:
    pos = 0
    for char in stacks[s]:
        if char != "$":
            ship[pos].put(char)
              

        pos = pos +1

    s = s -1

instructions = instructions.strip()
instructions = instructions.replace("move ","")
instructions = instructions.replace(" from ",",")
instructions = instructions.replace(" to ",",")

instructions = instructions.split("\n")

for ins in instructions:
    ins = ins.split(",")
    frm = int(ins[1]) -1
    to = int(ins[2]) -1
    for _ in range(int(ins[0])):
        itm = ship[frm].get()
        ship[to].put(itm)

top = ""
for x in range(len(stacknum)):
    top = top + ship[x].get()

print(top)
