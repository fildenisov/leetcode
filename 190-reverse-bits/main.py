from functools import *

def reverseBits( n: int) -> int:
    bts = [c for c in '{0:032b}'.format(n)]
    l = len(bts)
    for i in range(l//2):
        bts[i], bts[l-i-1] = bts[l-i-1], bts[i]

    return int(reduce(lambda x, y: x+y, bts), 2)

if __name__ == "__main__":
    print(reverseBits(43261596))
