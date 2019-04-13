#!/usr/bin/env python3

ZWSP={'0': '\u200b', '1': '\u200c', '2':'\u200d', '3': '\ufeff'}

def B10toN(num, base):
    isdev = int(num / base)
    if (isdev):
        return B10toN(isdev, base) + str(num%base)
    return str(num % base)

def Encode(string):
    quat = []
    for c in list(string):
        h = hex(ord(c))
        d = int(h, 16)
    
        quat.append(B10toN(d, 4))
    
    quatstr = ''.join(quat)
    print(type(quatstr))
    
    for d, s in ZWSP.items():
        quatstr = quatstr.replace(d, s)
    
    print(quatstr)

if __name__ == "__main__":
    Encode("ABC")
