from pycardano import PaymentSigningKey, StakeSigningKey, PaymentVerificationKey, StakeVerificationKey
import json

payment_signing_key = PaymentSigningKey.generate()
payment_verification_key = PaymentVerificationKey.from_signing_key(payment_signing_key)

data =

message = {}
