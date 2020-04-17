
import timeit
import cProfile

NUMOPS = 100
iters = 1000

def vectorAdd(df):
    for i in range(0,NUMOPS):
        df.add(1)
    return df

def numpyAdd(npa):
    for i in range(0,NUMOPS):
        npa.add(1)
    return npa


def benchMark(size,iters):
    print("Running Test for Array Size: "+ str(size))
    setup = '''from __main__ import vectorAdd, numpyAdd
import pandas as pd
import numpy as np
size=''' + str(size) + '''
df = pd.DataFrame(np.array(range(size)))
npa = np.array(range(size))
    '''
    #print(timeit.timeit('vectorAdd(df)', number=iters, setup=setup) / iters)
    print(timeit.timeit('numpyAdd(df)', number=iters, setup=setup) / iters)

def main():
    benchMark(10,iters)
    benchMark(10000,iters)
    benchMark(100000,iters)
    benchMark(1000000, 100)
    benchMark(10000000, 10)

if __name__ == "__main__":
    main()