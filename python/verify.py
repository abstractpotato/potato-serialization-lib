from sys import argv
from cose.messages import CoseMessage
from cose.headers import KID
from cose.keys.keytype import KtyOKP
from cose.keys.curves import Ed25519
from cose.keys.keyparam import (
    KpAlg,
    KpKeyOps,
    KpKty,
    OKPKpCurve,
    OKPKpD,
    OKPKpX
)
from cose.keys.keyops import SignOp, VerifyOp
from cose.keys import CoseKey
from cbor2 import CBORTag, dumps, loads
from nacl.encoding import RawEncoder
from nacl.hash import blake2b
import addr

def verify(signed_message, key=""):
    decoded_message = CoseMessage.decode(bytes.fromhex("d2" + signed_message))

    if key != "":
        cose_key = CoseKey.decode(bytes.fromhex(key))
        vkey = cose_key[OKPKpX]
    else:
        vkey = decoded_message.phdr[KID]
        cose_key = {
            KpKty: KtyOKP,
            OKPKpCurve: Ed25519,
            KpKeyOps: [SignOp, VerifyOp],
            OKPKpX: vkey,  # public key
        }
        cose_key = CoseKey.from_dict(cose_key)

    decoded_message.key = cose_key
    signature_verified = decoded_message.verify_signature()
    message = decoded_message.payload.decode("utf-8")
    akey = decoded_message.phdr["address"]

    addresses_match = (
        akey[1:].hex() == blake2b(vkey, 28, encoder=RawEncoder).hex()
    )
    verified = [
        str(signature_verified & addresses_match).lower(),
        addr.encode("addr", akey),
        message
    ]
    return verified

if __name__ == "__main__":
    if len(argv) == 2: print("\n".join(verify(argv[1])))
    if len(argv) == 3: print("\n".join(verify(argv[1], argv[2])))
