def sort_pair(arr0, arr1):
    if len(arr0) > len(arr1):
        return arr1, arr0
    else:
        return arr0, arr1
    
def merge(arr0, arr1):
    inversions = 0
    result = []
    # two indices to keep track of where we are in the array
    i0 = 0
    i1 = 0
    # probably doesn't really save much time but neater than calling len() everywhere
    len0 = len(arr0)
    len1 = len(arr1)
    while len0 > i0 and len1 > i1:
        if arr0[i0] <= arr1[i1]:
            result.append(arr0[i0])
            i0 += 1
        else:
            # count the inversion right here: add the length of left array
            inversions += len0 - i0
            result.append(arr1[i1])
            i1 += 1

    if len0 == i0:
        result += arr1[i1:]
    elif len1 == i1:
        result += arr0[i0:]

    return result, inversions

def sort(arr):
    length = len(arr)
    mid = length//2
    if length >= 2:
        sorted_0, counts_0 = sort(arr[:mid])
        sorted_1, counts_1 = sort(arr[mid:])
        result, counts = merge(sorted_0, sorted_1)
        return result, counts + counts_0 + counts_1
    else:
        return arr, 0

def countInversions(a):
    final_array, inversions = sort(a)
    # print(final_array)
    return inversions


if __name__ == '__main__':
    print(count_inversions([5,1,9,3,2,1,1,8,5,11,6,5,2,1,2,7]))