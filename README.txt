- Single Responsability: create an auth microservice
-It 'll be based on OAuth2 authentication
-It 'll use JWT with PKI (Private and public key)
-Private key will be used for signing token
-public key will be used  for verifying token
-for hashing password we'll use bcrypt
-for frontend security, we 'll use CORS
- Each microservice using this authentication microservice  will use public key for verification
- This microservice will be interfaces with HTTP REQUEST ("I 'll not call it REST API to respect Roy Fielding)
- We 'll use the standard library net/http and database/sql package for performance purpose
no ORM, no Web Framework (also for challenges)

OAuth2 principle:
-user will send username/password
-if username/password is correct:
-the server will send two kind of token:
authentication token and refresh token
-else:
- the server will send error
- the user will use the auth token to get access to service
- when auth token expired, user will send refresh_token to server
for getting new pair/token

JWT principle:
- JWT is the tool that we 'll use to generate token
- a JWT token have 3 part:
.hashing method
.payload
.Signature
- the important part is the Signature: it 'll verify that the token is from the server
- On the payload we'll just insert user information
- for signing the token, on JWT we have two method:
- HMAC method using SECRET_KEY
- PKI method using a key cut: private and public
- We'll use private key for signing token and public key for verifying it

Bcrypt principle:
Bcrypt is an hashing method.
I'm not yet expert to it but it is comonly used and I know why
Why?
Because Bcrypt instead of other hashing method can adapt to the current computer power to make difficult
brute force attempt. 
Bcrypt encryption function have a "cost" parameter that is used to make slower the hashing step then
it 'll be more difficult for a bruteforce attacker to create rainbow table if they got the password

CORS principle:
CORS or cross origin ressource sharing
Securised client like modern navigator use SameOrigin policy, this policy principle is
to only autorise connexion between frontend and backend service witch have the same origin:

Having same origin means:
- Having same protocol
- Having same ip
- Having same port

For example: http://localhost is same origin as http://localhost:80
by default http is always using 80 so http://localhost use 80 port

The purpose of SameOrigin policy is to protect user from attack like CSRF forgery or Cross site scripting XSS
But it provoke error when your frontend app and backend app doesn't have same origin
for example, if you use react:
http://localhost:3000
and a server runing on:
http://localhost:8000
When you try to access to the ressource on the server,you 'll get sameorigin error, more specificaly: a " CORS AccessControlAllowOrigin" error

CORS  permit you to fix this error by specify list of origin that can access to the server ressource
for this purpose he will add cors header on the http response that tell the client if the origin can have access to the server
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: *
Access-Control-Allow-Headers: *

If you use "*", it 'll not be good because every origin will be allowed to access your ressource on the server

