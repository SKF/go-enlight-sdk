#!/bin/bash -e
# You are not expected to run this script. The certificates are all generated to expire in a 100 years.
# Left for documentation purposes only.

rm -f *.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 36500 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=SE"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=SE/CN=localhost"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -extfile <(printf "subjectAltName=DNS:localhost") -in server-req.pem -days 36500 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=SE/CN=localhost"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -extfile <(printf "subjectAltName=DNS:localhost") -in client-req.pem -days 36500 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem

echo "Client's signed certificate"
openssl x509 -in client-cert.pem -noout -text

rm client-req.pem server-req.pem ca-cert.srl ca-key.pem
