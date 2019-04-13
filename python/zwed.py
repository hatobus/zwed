#!/usr/bin/env python3
import codecs


ZWSP={'0': '\u200b', '1': '\u200c', '2':'\u200d', '3': '\ufeff'}

class ZWED():
    def __init__(self):
        pass

    def B10toN(self, num, base):
        isdev = int(num / base)
        if (isdev):
            return self.B10toN(isdev, base) + str(num%base)
        return str(num % base)
    
    def BNto10(self, num, base):
        out = 0
    
        for i in range(1, len(str(num))+1):
            out += int(num[-i])*(base**(i-1))
        
        return out
    
    def Encode(self, string):
        quat = []
        for c in list(string):
            h = hex(ord(c))
            d = int(h, 16)
        
            quat.append(self.B10toN(d, 4))
        
        quatstr = ''.join(quat)
        
        for d, s in ZWSP.items():
            quatstr = quatstr.replace(d, s)
        
        return quatstr
    
    def Decode(self, zwstr):
        for d, s in ZWSP.items():
            zwstr = zwstr.replace(s, d) 
    
        zwlist = []
        start = 0
    
        for i in range(int(len(zwstr)/4)):
            end = start+4
            zwlist.append(zwstr[start:end])
            start = end
    
        b10list = [self.BNto10(c, 4) for c in zwlist]
        hexlist = [hex(i) for i in b10list]
    
        hexstrx0 = ''.join(hexlist)
        hexstr = hexstrx0.replace("0x", "")
        bytstr = codecs.decode(hexstr, "hex")
    
        return bytstr.decode("UTF-8")


if __name__ == "__main__":
    ZW = ZWED()
    zwstr = ZW.Encode("ABC")
    print(zwstr)
    decoded = ZW.Decode(zwstr)
    print(decoded)
