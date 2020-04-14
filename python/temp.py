
import timeit
import cProfile

NUMOPS = 100000
iters = 10

def vectorAdd(df):
    for i in range(0,NUMOPS):
        df.add(1)
    return df

def benchMark(size):
    print("Running Test for Size: "+ str(size))
    setup = '''from __main__ import vectorAdd
import pandas as pd
import numpy as np
size=''' + str(size) + '''
df = pd.DataFrame(np.array(range(size)))
    '''
    print(timeit.timeit('vectorAdd(df)', number=iters, setup=setup) / iters)

def main():
    #benchMark(10)
    #benchMark(10000)
    benchMark(100000)
    #benchMark(1000000)
    #benchMark(10000000)

if __name__ == "__main__":
    main()