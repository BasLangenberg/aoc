import resource, sys
resource.setrlimit(resource.RLIMIT_STACK, (2**29,-1))
sys.setrecursionlimit(10**6)

def setup_graph(inputFile):
    # Get all nodes
    nodes = set()
    for x in open(inputFile):
        line = str.split(x, "-")
        for x in line:
            nodes.add(str.strip(x))

    g={}

    for x in nodes:
        g[x] = list()

    for x in open(inputFile):
        line = str.split(str.strip(x), "-")
        g[line[0]].append(line[1])
        g[line[1]].append(line[0])

    return g

def find_all_paths(graph, start, end, path=[]):
    path = path + [start]
    if start == end:
        return [path]
    if start not in graph:
        return []
    paths = []
    for node in graph[start]:
        if node.isupper() or node not in path:
            newpaths = find_all_paths(graph, node, end, path)
            for newpath in newpaths:
                paths.append(newpath)
    return paths

def has_more_than_one_double_lower(g):
    s = set()
    count = 0

    for x in g:
        if x.islower() and x in s:
            if count != 1:
                count = 1
            else:
                return True 

        s.add(x)

    return False


def find_all_paths_part2(graph, start, end, path=[]):
    path = path + [start]
    if start == end:
        return [path]
    if start not in graph:
        return []
    paths = []
    for node in graph[start]:
        if (node.isupper() or (node.islower() and not has_more_than_one_double_lower(path))) and node != "start":
            newpaths = find_all_paths_part2(graph, node, end, path)
            for newpath in newpaths:
                paths.append(newpath)
    return paths

graph = setup_graph("input")

print(len(find_all_paths(graph, "start", "end")))
print(len(find_all_paths_part2(graph, "start", "end")))
