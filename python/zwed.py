import codecs
import binascii


ZWSP={'0': '\u200b', '1': '\u200c', '2':'\u200d', '3': '\ufeff'}

class ZWED():
    def __init__(self):
        pass

    # Exchange base N --> 10
    def B10toN(self, num, base):
        isdev = int(num / base)
        if (isdev):
            return self.B10toN(isdev, base) + str(num%base)
        return str(num % base)
    
    # Excange base 10 --> N
    def BNto10(self, num, base):
        out = 0
    
        for i in range(1, len(str(num))+1):
            out += int(num[-i])*(base**(i-1))
        
        return out
    
    def Encode(self, string):
        string = string.replace("\n", "")
        
        bytstr = binascii.hexlify(string.encode("utf-8"))
        hexstr = bytstr.decode("utf-8")
        dec = int(hexstr, 16)

        quatstr = self.B10toN(dec, 4)

        for d, s in ZWSP.items():
            quatstr = quatstr.replace(d, s)
        
        return quatstr
    
    def Decode(self, zwstr):
        for d, s in ZWSP.items():
            zwstr = zwstr.replace(s, d) 
    
        dec = self.BNto10(zwstr, 4)
        print(dec)
        hexstr = hex(dec)[2:]
        print(hexstr)

        print(binascii.unhexlify(hexstr.encode("utf-8")))

        return str(binascii.unhexlify(hexstr), 'utf-8')

if __name__ == "__main__":
    ZW = ZWED()
    #zwstr = ZW.Encode("あいうえお")
    zwstr = ZW.Encode("abcdefgh")
    print(zwstr)
    decoded = ZW.Decode(zwstr)
    print(decoded)
    #print(decoded.decode('utf-8'))
