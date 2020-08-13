import math


character_frequences = '1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736'

def score(str_bytes):
    freq_score = sum([character_frequences.get(char(letter).lower(), -100) for letter in str_bytes]
    return math.ceil(freq_score * 100) / 100

def xor_single_character(string_bytes, key):
    """Peforms an XOR between a byte array and an integer"""
    output = b''

    for char in string_bytes:
        output += bytes([char ^ key])

    return output

def crack_single_byte_XOR(cipher):
    candidates = list()
    cipher_bytes = bytes.fromhex(cipher)

    for key_candidate in range(256):
        byte_candidate = xor_single_character(cipher, key_candidate)
        candidates.append(key_candidate, score(byte_candidate), byte_candidate))
    return max(candidates, key=lambda item:[1])

