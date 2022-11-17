import pprint

def create_paper(inp):
    with open(inp) as x:
        instructions = x.read()

    instructions = instructions.split("\n\n")

    coords = instructions[0]
    inst = instructions[1]
    crds = list()

    for x in coords.split("\n"):
        crd = x.split(",") 
        x = int(crd[0])
        y = int(crd[1])

        c = (x, y)

        crds.append(c)

    xy = max_x_y(crds)

    graph = create_graph(xy)
    fill_graphs(graph, crds)
    fold_graph(graph, inst)

    return graph

def fold_graph(graph, inst):
    for i in inst.rstrip().split("\n"):
        instruction = i.split(" ")[2]
        parsed_inst = instruction.split("=")

        if parsed_inst[0] == "y":
            up = graph[:int(parsed_inst[1])]
            down = graph[int(parsed_inst[1])+1:]

            for y in range(len(down)):
                for x in range(len(down[0])):
                    if down[y][x] == "#":
                        up[y][x] = down[len(up)-y][len(up[0])-x]

            pprint.pprint(up)

            graph = up

        if parsed_inst[0] == "x":
            for row in graph:
                left = row[:int(parsed_inst[1])]
                right = row[int(parsed_inst[1])+1:]

                for c in range(len(right)):
                    left[c] = right[len(right)-c]

                row = left

def create_graph(max_xy):
    return [["."] * (int(max_xy[0])) for _ in range(int(max_xy[1]))]

def fill_graphs(graph, coords):
    for c in coords:
        graph[c[1]][c[0]] = "#"


def max_x_y(coords):
    max_x = 0
    max_y = 0

    for pair in coords:
        if pair[0] > max_x:
            max_x = pair[0]

        if pair[1] > max_y:
            max_y = pair[1]

    return (max_x +1, max_y +1)
    
def print_paper(paper):
    pprint.pprint(paper)

inp = "input-tst"
create_paper(inp)
