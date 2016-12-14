#!/usr/bin/env python2

import re
import matplotlib.pyplot as plt
import math
import numpy as np

expr = re.compile("(\d+)\s+(\d+\.?\d+e*-*\d*)(?:\s+(\d+\.?\d*))?")

with open("timings.txt") as f:
    input_file = f.readlines()

results = []

for line in input_file:
    result = expr.match(line)

    n, g, g_py = result.group(1), result.group(2), result.group(3)
    n, g, g_py = int(n), float(g), float(g_py or 0)

    results.append([n,g,g_py])

def get_column(results, index):
    return [r[index] for r in results]

def get_exponential_column(results, index):
    return [math.log(r[index]) for r in results]

results = results[4:]

reg_go = np.polyfit(get_column(results, 0),get_exponential_column(results, 1), 1)
reg_py = np.polyfit(get_column(results[:7], 0),get_exponential_column(results[:7], 2), 1)


print(reg_go)
print(reg_py)

plt.scatter(get_column(results, 0), get_exponential_column(results, 1))
plt.show()
