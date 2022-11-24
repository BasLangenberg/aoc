# Stole a bunch from https://github.com/simonbrahan/aoc2021/tree/master/15
from collections import defaultdict
from heapq import heappop, heappush
import math

def get_neighbours(cell, grid):
    x, y = cell
    directions = [(-1, 0), (0, -1), (0, 1), (1, 0)]

    out = []
    for add_x, add_y in directions:
        neighbour_x = x + add_x
        neighbour_y = y + add_y

        if 0 <= neighbour_y < len(grid) and 0 <= neighbour_x < len(grid[neighbour_y]):
            out.append((neighbour_x, neighbour_y))

    return out


def get_risk(grid):
    destination = (len(grid) - 1, len(grid[0]) - 1)
    cumulative_cell_risks = defaultdict(lambda: math.inf)
    cumulative_cell_risks[(0, 0)] = 0
    cell_visit_queue = []
    heappush(cell_visit_queue, (0, (0, 0)))

    unvisited_cells = set()
    for y, row in enumerate(grid):
        for x, cell in enumerate(row):
            unvisited_cells.add((x, y))

    while destination in unvisited_cells:
        current_risk, current = heappop(cell_visit_queue)

        # heap is tricky to keep clean, so sometimes a visited cell will be added more than once.
        # No bother; just ignore it.
        if current not in unvisited_cells:
            continue

        neighbours = get_neighbours(current, grid)

        for neighbour in neighbours:
            if neighbour not in unvisited_cells:
                continue

            neighbour_risk = min(
                cumulative_cell_risks[neighbour],
                cumulative_cell_risks[current] + grid[neighbour[1]][neighbour[0]]
            )

            cumulative_cell_risks[neighbour] = neighbour_risk
            heappush(cell_visit_queue, (neighbour_risk, neighbour))

        unvisited_cells.remove(current)

    return cumulative_cell_risks[destination]

from math import floor

def translate_risk(risk, x, y, tile_width, tile_height):
    x_translation = floor(x / tile_width)
    y_translation = floor(y / tile_height)

    risk_translation = (risk + x_translation + y_translation -1) % 9 + 1

    return risk_translation


with open('input') as f:
    grid_tile = [[int(risk) for risk in line.strip()] for line in f]

tile_width = len(grid_tile[0])
full_width = tile_width * 5
tile_height = len(grid_tile)
full_height = tile_height* 5

full_grid = [[] for y in range(full_height)]

for y in range(full_height):
    for x in range(full_width):
        full_grid[y].append(translate_risk(grid_tile[y%tile_height][x%tile_width], x, y, tile_width, tile_height))


with open('input') as f:
    grid = [[int(risk) for risk in line.strip()] for line in f]

print(get_risk(grid))
print(get_risk(full_grid))

