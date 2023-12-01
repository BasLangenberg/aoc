#!/usr/bin/env python3

import astar

input = 'input'
input = 'input-tst'

grid = []


with open(input) as f:
    for line in f.readlines():
        grid.append([l for l in line.strip()])

def find_in_graph(char):
    for rownum in range(len(grid)):
        for colnum in range(len(grid[rownum])):
            if grid[rownum][colnum] == char:
                return (rownum, colnum)

adjacency_list = {}
for rownum in range(len(grid)):
    for colnum in range(len(grid[rownum])):
        node = (rownum, colnum)
        neighbours = []
        if rownum - 1 > 0:
            neighbours.append((rownum-1,colnum))
        if rownum + 1 < len(grid):
            neighbours.append((rownum+1,colnum))
        if colnum - 1 > 0:
            neighbours.append((rownum,colnum-1))
        if colnum + 1 < len(grid[0]):
            neighbours.append((rownum,colnum+1))

        weighted_neighbours = []
        for n in neighbours:
            weighted_neighbours.append((n, grid[n[0]][n[1]]))
            
        adjacency_list[node] = weighted_neighbours

        #adjacency_list = {
        #'A': [('B', 1), ('C', 3), ('D', 7)],
        #'B': [('D', 5)],
        #'C': [('D', 12)]
        #}
graph1 = astar.Graph(adjacency_list)
S = find_in_graph("S")
E = find_in_graph("E")

graph1.a_star_algorithm(S, E)
