from itertools import combinations
from collections import defaultdict

def is_GP(gp, r):
    """
    to determine the either gp integer is in Geometrics progresssion or not
    """
    # gp = sorted(gp)
    if int(gp[0] / r**0) == int(gp[1] / r**1) == int(gp[2] / r**2):
        return True
    else:
        return False

def countTriplets(arr, r):
    return sum([is_GP(i, r) for i in combinations(arr, 3)])

def countTripletsv1(arr, r):
    v2 = defaultdict(int)
    v3 = defaultdict(int)
    count = 0
    for k in arr:
        count += v3[k]
        v3[k*r] += v2[k]
        v2[k*r] += 1
    return count

if __name__ == '__main__':
    val = countTriplets([12,5,24,48,96, 48], 2)
    countTripletsv1([12,5,24,48,96, 48], 2)
    print(val)