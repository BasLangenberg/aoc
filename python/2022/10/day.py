#!/usr/bin/env python3

input = 'input'
#input = 'input-tst'

x_hist = [1]
crt = []
crtline = ""

def check_crt(crtline):
    if len(crtline) == 40:
         crt.append(crtline)
         return ""

    return crtline

with open(input) as f:
    for line in f.readlines():

        ins = line.split(" ")
        if ins[0] == "addx":
            print("add", x_hist[-1])
            crtline = check_crt(crtline)
            if x_hist[-1] == len(crtline) or x_hist[-1] == len(crtline) -1 or x_hist[-1] == len(crtline) + 1:
                crtline += "#"
            else:
                crtline += "."
            x_hist.append(x_hist[-1])
            crtline = check_crt(crtline)
            if x_hist[-1] == len(crtline) or x_hist[-1] == len(crtline) -1 or x_hist[-1] == len(crtline) + 1:
                crtline += "#"
            else:
                crtline += "."
            print("add", x_hist[-1], "+", int(ins[1]))
            x_hist.append(x_hist[-1] + int(ins[1]))
            continue

        print("noop, add", x_hist[-1])
        crtline = check_crt(crtline)
        if x_hist[-1] == len(crtline) or x_hist[-1] == len(crtline) -1 or x_hist[-1] == len(crtline) + 1:
            crtline += "#"
        else:
            crtline += "."
        x_hist.append(x_hist[-1])

    
strengths = 0
strengths += x_hist[19] * 20
strengths += x_hist[59] * 60
strengths += x_hist[99] * 100
strengths += x_hist[139] * 140
strengths += x_hist[179] * 180
strengths += x_hist[219] * 220

print(strengths)

for line in crt:
    print(line)

print(crtline)
