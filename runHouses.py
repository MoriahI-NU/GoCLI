#!/usr/bin/env python


#####!/Users/Moria/AppData/Local/Programs/Python/Python310/ python

#####!/usr/bin/env python

import pandas as pd
import time

# number of runs
N = 100 

# start keeping track of time
start = time.process_time()

#read input file, calculate and write summary stats to output file - repeat 100 times
with open('housesOutputPy.txt', 'wt') as outfile:
    for i in range(N):
        houses = pd.read_csv("housesInput.csv")
        outfile.write(houses.describe().to_string(header=True, index=True))
        outfile.write("\n")

#calculate execution time
total = time.process_time() - start

print("Execution Time (sec) for 100 runs = ", total)
print("Average Execution Time (sec) per run = ", total/100)
