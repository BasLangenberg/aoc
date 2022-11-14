def graph(inputFile):
    # Get all nodes
    nodes = set()
    for x in open(inputFile):
        line = str.split(x, "-")
        for x in line:
            nodes.add(str.strip(x))

    print(nodes)


graph("input-tst")
