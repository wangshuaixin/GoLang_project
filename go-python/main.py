# from ctypes import cdll

# lib = cdll.LoadLibrary('./libadd.so')

# print("Loaded GO generated SO library")
# print("Loaded add function compiled in GO")
# print("Testing add(2,3):")
# result = lib.add(2,3)
# print(result)

import ctypes
import os
print(os.listdir())
add = ctypes.CDLL('./go-python/libadd.so').add
print('so added')

add.argtypes = [ctypes.c_char_p,ctypes.c_char_p]
add.restype = ctypes.c_char_p

left = b"hello"
right = b"world"

print(add(left, right))