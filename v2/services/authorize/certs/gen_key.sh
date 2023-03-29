#!/bin/bash -e
# Left for documentation purposes only. Generates a shared RSA key used in
# tests involving TLS certificates to avoid wasting time generating keys.

rm -f *.pem

openssl req -newkey rsa:4096 -nodes -keyout rsa-key.pem -subj "/" > /dev/null
