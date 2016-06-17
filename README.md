# crypto
Some cryptographic attacks, protocols, or other functions.

## log
Solving discret logarithm problem. Given a group **G** of order **p**,a generator **g** of G, **h** an element of G, finds **x** where **g^x = h mod p**.
 - [x] Brute force attack
 - [x] Shanks' Baby Steps Giant Steps
 - [ ] Rho Pollard
```
BruteForce(g, h, p *Int) *Int
Shanks(g, h, p *Int) *Int 
```
