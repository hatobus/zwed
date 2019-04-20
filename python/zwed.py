import binascii


ZWSP={'0': '\u200b', '1': '\u200c', '2':'\u200d', '3': '\ufeff'}

class ZWED():
    def __init__(self):
        pass

    # Exchange base 10 --> N
    def B10toN(self, num, base):
        res = []
        while num:
            amari =str(num % base)
            res.insert(0, amari)
            num = num // base
            if num == 0:
                break

        s = "".join(res)
        return s
    
    # Excange base N --> 10
    def BNto10(self, num, base):
        out = 0
    
        for i in range(1, len(str(num))+1):
            out += int(num[-i])*(base**(i-1))
        
        return out
    
    def Encode(self, string):
        #string = string.replace("\n", "")
        
        bytstr = binascii.hexlify(string.encode("utf-8"))
        hexstr = bytstr.decode("utf-8")

        # base 16 --> 10
        dec = int(hexstr, 16)

        quatstr = self.B10toN(dec, 4)

        # Replace zero width spaces
        for d, s in ZWSP.items():
            quatstr = quatstr.replace(d, s)
        
        return quatstr
    
    def Decode(self, zwstr):
        for d, s in ZWSP.items():
            zwstr = zwstr.replace(s, d) 
    
        # base 4 --> 10
        dec = self.BNto10(zwstr, 4)

        hexstr = hex(dec)[2:]

        return str(binascii.unhexlify(hexstr), 'utf-8')

if __name__ == "__main__":
    ZW = ZWED()
    zwstr = ZW.Encode("こんにちは世界")
    #zwstr = ZW.Encode("abcdefgh")
    print(zwstr)
    decoded = ZW.Decode(zwstr)
    print(decoded)
