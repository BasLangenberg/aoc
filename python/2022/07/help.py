# Credit: https://raw.githubusercontent.com/junefish/adventofcode/main/adventofcode2022/day7/day7problem1.py
# import modules
from collections import defaultdict

# initialise variables
terminal_output = []
filepath = []
sizes = defaultdict(int)
total = 0
max_size = 100000

# read input file
with open('input', 'r') as file:
    for line in file:
        terminal_output.append(line.strip())

# parse input commands
for line in terminal_output:
    # change directories
    if(line.startswith('$ cd')):
        directory = line.split()[-1]
        # go to previous directory
        if(directory == '..'):
            filepath.pop()
        # add directory to filepath
        else:
            filepath.append(directory)
    
    # list contents
    elif(line.startswith('$ ls')):
        continue
    
    # parse ls output for sizes
    else:
        size, _ = line.split()
        if(size.isdigit()):
            size = int(size)
            for i in range(len(filepath)):
                sizes['/'.join(filepath[:i+1])] += size

# calculate sum of directories with size at most 100k
for key, value in sizes.items():
    if(value <= 100_000):
        total += value

# print answer
print("part 1:", total)

# Part 2 - This is written by me

total_disk = 70000000
req_unused = 30000000
to_delete = req_unused - (total_disk - sizes['/'])

min_req = 10000000000000000000
for _, value in sizes.items():
    if (value > to_delete) and (value < min_req):
        min_req = value

print("part2:", min_req)
