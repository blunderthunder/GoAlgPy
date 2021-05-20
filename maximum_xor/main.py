"return maximum xor given arr with queries"

def maxXor(arr, queries):
    ans = []
    trie = {}
    k = len(bin(max(arr+queries))) - 2 
    for number in ['{:b}'.format(x).zfill(k) for x in arr]:
        node = trie
        for char in number:
            node = node.setdefault(char, {})
    for n in queries:
        node = trie
        s = ''
        for char in'{:b}'.format(n).zfill(k) :
            tmp = str(int(char) ^ 1) 
            tmp = tmp if tmp in node else char 
            s += tmp 
            node = node[tmp]
        yield int(s, 2) ^ n


if __name__ == '__main__':
    print(list(maxXor([3,7,15,10], [3,7])))