#!/usr/bin/env python3

import re

with open("input") as f:
    text = f.read()

regex = r'mul\((\d+,\d+)\)'

p = re.compile(regex)

print(p.findall(text))
total = 0
for i in p.findall(text):
    tup = i.split(",")
    total += int(tup[0]) * int(tup[1])

print(total)

regex = r'mul\((\d+,\d+)\)|(do\(\))|(don\'t\(\))'
regex = r"mul\((\d+,\d+)\)|(do\(\))|(don\'t\(\))"

p = re.compile(regex)

total = 0
enable = True
print(p.findall(text))
for i in p.findall(text):
    if i[2] == "don't()":
        enable = False
        continue
    if i[1]== "do()":
        enable = True
        continue
    if enable:
        tup = i[0].split(",")
        total += int(tup[0]) * int(tup[1])

print(total)
