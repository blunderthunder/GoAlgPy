from collections import defaultdict

def freqQuery(queries):
    
    counter = defaultdict(lambda : 0)
    arr = defaultdict(int)
    frequencies = defaultdict(int)
    
    for command,value in queries:
        if command == 1:
            arr[value] += 1
            frequencies[arr[value]] += 1
            frequencies[arr[value] - 1] -= 1
        if command == 2 and arr[value] != 0:
            arr[value] -= 1
            frequencies[arr[value]] += 1
            frequencies[arr[value] + 1] -= 1
        if command == 3:
            ans = 1 if frequencies[value] > 0 else 0
            yield ans

if __name__ == '__main__':
    queries = []
    with open("frequencyQueries/test.file") as fp:
        for i in fp:
            queries.append(tuple(map(lambda x: int(x), i.strip().split())))
    
    print(list(freqQuery(queries)))