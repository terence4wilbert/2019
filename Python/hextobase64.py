import base64
from base64 import binascii

def hex_to_base64(hex_str):
    return base64.b64encode(binascii.unhexlify(hex_str))



print(str(hex_to_base64('49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d'), "utf-8"))