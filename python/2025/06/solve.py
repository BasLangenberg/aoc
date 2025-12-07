#!/usr/bin/env python3

sum = 0
grid = list()
#with open("test-input") as f:
with open("input") as f:
    for line in f.readlines():
        grid.append(line.strip("\n").split())

lastline = len(grid)-1
for col in range(len(grid[0])):
    nums = list()
    for cell in range(lastline):
        nums.append(int(grid[cell][col]))

    print(nums)
    if grid[lastline][col] == "*":
        ans = nums[0]
        for i in range(1, len(nums)):
            ans *= nums[i]
        print(ans)
        sum += ans
        continue
    if grid[lastline][col] == "+":
        ans = nums[0]
        for i in range(1, len(nums)):
            ans += nums[i]
        print(ans)
        sum += ans
        continue

    print("COULD NOT FIND OP")


print(sum)
