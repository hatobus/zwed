#!/usr/bin/env python3
import codecs


ZWSP={'0': '\u200b', '1': '\u200c', '2':'\u200d', '3': '\ufeff'}

def B10toN(num, base):
    isdev = int(num / base)
    if (isdev):
        return B10toN(isdev, base) + str(num%base)
    return str(num % base)

def BNto10(num, base):
    out = 0

    for i in range(1, len(str(num))+1):
        out += int(num[-i])*(base**(i-1))
    
    return out

def Encode(string):
    quat = []
    for c in list(string):
        h = hex(ord(c))
        d = int(h, 16)
    
        quat.append(B10toN(d, 4))
    
    quatstr = ''.join(quat)
    print(quatstr)
    
    for d, s in ZWSP.items():
        quatstr = quatstr.replace(d, s)
    
    print(quatstr)
    return quatstr

def Decode(zwstr):
    for d, s in ZWSP.items():
        zwstr = zwstr.replace(s, d) 

    print(zwstr)
    
    zwlist = []
    start = 0

    for i in range(int(len(zwstr)/4)):
        end = start+4
        zwlist.append(zwstr[start:end])
        start = end

    print(zwlist)

    b10list = [BNto10(c, 4) for c in zwlist]
    hexlist = [hex(i) for i in b10list]

    hexstrx0 = ''.join(hexlist)
    hexstr = hexstrx0.replace("0x", "")
    bytstr = codecs.decode(hexstr, "hex")

    return bytstr.decode("UTF-8")

if __name__ == "__main__":
    zwstr = Encode("ABC")
    decoded = Decode(zwstr)
    print(decoded)
