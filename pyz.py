import hashlib
import time
import base64
import pyotp
from pyotp.utils import build_uri

# The challenge here is to Generate TOTP 10 Digit Password for their Basic Authentication
# The needed is Secret Shared, and it should be encrypted with HMAC-SHA-512 instead of HMAC-SHA-1 (Default)
secret_key = 'HENNGECHALLENGE003'
msg = 'xsephtionx@gmail.com'
key = msg+secret_key # Convert this first to b32 -> byte
passw = pyotp.TOTP(base64.b32encode(bytearray(key, 'ascii')), digits=10, digest=hashlib.sha512) # This the one that is needed.
# passw = pyotp.HOTP(base64.b32encode(bytearray(key, 'ascii')), digits=10, digest=hashlib.sha512) # Tester
passwz = passw.now() # Generate password now
print(passwz) # just copy paste the password. I'm using Postman to send requests 
time.sleep(5)
print(passw.verify(passwz))