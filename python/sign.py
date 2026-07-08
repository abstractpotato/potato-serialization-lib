from sys import argv
from cose.headers import KID, Algorithm
from cose.messages import CoseMessage, Sign1Message
from cose.algorithms import EdDSA
from cose.keys.keytype import KtyOKP
from cose.keys.curves import Ed25519
from cose.keys.keyops import SignOp, VerifyOp
from cose.keys.keyparam import (
    KpAlg,
    KpKeyOps,
    KpKty,
    OKPKpCurve,
    OKPKpD,
    OKPKpX
)
from cose.keys import CoseKey
from cbor2 import CBORTag, dumps
from nacl import bindings
import hashlib

def get_keys():
    keys = {}
    with open(".env/skey", "rb") as file:
        keys["skey"] = file.read()
    with open(".env/vkey", "rb") as file:
        keys["vkey"] = file.read()
    with open(".env/akey", "rb") as file:
        keys["akey"] = file.read()
    return keys

def sign(message: str, keys: dict):

    msg = Sign1Message(
        phdr={ Algorithm: EdDSA, "address": keys["akey"] },
        payload=message.encode("utf-8")
    )
    msg.uhdr = {"hashed": False}
    msg.phdr[KID] = keys["vkey"]

    cose_key = {
        KpKty: KtyOKP,
        OKPKpCurve: Ed25519,
        KpKeyOps: [SignOp, VerifyOp],
        OKPKpD: keys["skey"],
        OKPKpX: keys["vkey"]
    }
    cose_key = CoseKey.from_dict(cose_key)
    msg.key = cose_key

    _message = [
        msg.phdr_encoded,
        msg.uhdr_encoded,
        msg.payload,
        sign_with_private_key(keys["skey"][:64], msg._sig_structure)
    ]

    encoded = dumps(
        CBORTag(msg.cbor_tag, _message),
        default=msg._custom_cbor_encoder
    )

    signed_message = encoded.hex()[2:]
    return signed_message

def sign_with_private_key(private_key, sig_structure):
    right = private_key[32:]
    left = private_key[:32]
    r = bindings.crypto_core_ed25519_scalar_reduce(
        hashlib.sha512(right + sig_structure).digest()
    )
    R = bindings.crypto_scalarmult_ed25519_base_noclamp(r)

    pubk = bindings.crypto_scalarmult_ed25519_base_noclamp(left)
    hram = bindings.crypto_core_ed25519_scalar_reduce(
        hashlib.sha512(R + pubk + sig_structure).digest()
    )
    S = bindings.crypto_core_ed25519_scalar_add(
        bindings.crypto_core_ed25519_scalar_mul(hram, left),
        r
    )
    return R + S

if __name__ == "__main__":
    print(sign(argv[1], get_keys()))
