
def isBalanced(s: str) -> bool:
    stack = []
    pairs = {"{": "}", "[": "]", "(" : ")"}
    for i in s:
        if not stack:
            stack.append(i)
        elif i == pairs.get(stack[-1]):
            stack.pop()
        else:
            stack.append(i)
    return "YES" if not stack else "NO"


if __name__ == '__main__':

    s = '}][}}(}][))]'
    val = isBalanced(s)
    print('Is Brackets really balanced ?', val)