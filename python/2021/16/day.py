with open('input-tst') as f:
    line = f.read().strip()

decoded = ""
for char in line:
    decoded += format(int(char, 16), "0>4b")

pointer = 0

#while pointer < len(decoded) - 1:
    # DO STUFF    
