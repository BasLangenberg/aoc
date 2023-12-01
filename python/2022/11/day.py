#!/usr/bin/env python3

input = 'input'
#input = 'input-tst'

monkeys = []

class Monkey():

    def __init__(self, setup):
        self.inspections = 0
        lines = setup.split("\n")

        self.num = int(lines[0].split(" ")[1][0])

        self.items = []
        for item in lines[1].split(" ")[4:]:
            self.items.append(int(item.replace(",","")))

        self.op = lines[2].split(" ")[5:]

        self.test = int(lines[3].split(" ")[-1])
        self.ok = int(lines[4].split(" ")[-1])
        self.nok = int(lines[5].split(" ")[-1])


    def __repr__(self):
        return f"Money {self.num} | items={self.items}, op={self.op}, test={self.test}, ok={self.ok}, nok={self.nok}"
        
    def round(self):
        for num in range(len(self.items)):
            self.inspections += 1
            newitem = -1

            if self.op[2] == "old":
                rightop = self.items[num]
            else:
                rightop = int(self.op[2])
                
            if self.op[1] == "*":
                newitem = self.items[num] * rightop
            if self.op[1] == "+":
                newitem = self.items[num] + rightop

            newitem = int(newitem / 3)

            if newitem % self.test == 0:
                #print(f"{self.inspections} - {self.num} throw to monkey {self.ok}")
                monkeys[self.ok].items.append(newitem)
            else:
                #                print(f"{self.inspections} - {self.num} throw to monkey {self.nok}")
                monkeys[self.nok].items.append(newitem)

        self.items = []

with open(input) as f:
    for monkey in f.read().strip().split("\n\n"):
        mon = Monkey(monkey)
        monkeys.append(mon)

for _ in range(20):
    for i in monkeys:
        i.round()

inspections = []
for mon in monkeys:
    inspections.append(mon.inspections)

inspections.sort()
inspections.reverse()
print(inspections[0] * inspections[1])
