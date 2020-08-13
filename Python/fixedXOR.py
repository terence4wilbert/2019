

def fixed_XOR(hex_a, hex_b):
    xored = int(hex_a, 16) ^ int(hex_b, 16)
    return '{:x}'.format(xored)

print(fixed_XOR('1c0111001f010100061a024b53535009181c','686974207468652062756c6c277320657965'))