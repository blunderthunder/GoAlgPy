def minimumAbsoluteDifference(arr):
    arr.sort()
    initial = arr[0]
    later = 0
    MAD = 999999999999999
    for i in arr[1:]:
        later = i
        CMAD = abs(initial - later)
        if CMAD == 0:
            return 0
        if CMAD < MAD:
            MAD = CMAD

        initial = later 
    
    return MAD


if __name__=='__main__':
    val = minimumAbsoluteDifference([-59, -36, -13, 1, -53, -92, -2, -96, -54, 75])
    print("minimum absolute difference: ", val)