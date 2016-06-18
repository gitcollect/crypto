# crypto
Some cryptographic attacks, protocols, or other functions.

## factorization
Factorization algotirhm for composite integer. 

 - [x] Brute force attack
 - [x] Rho Pollard
 - [ ] P-1 Pollard
 - [ ] Fermat

```
BruteForce(n *big.int) *big.Int
RhoPollard(n *big.Int) *big.Int
```

## log
Solving discret logarithm problem. Given a group **G** of order **p**, a generator **g** , an element **h**, finds **x** where **g^x = h mod p**.

 - [x] Brute force attack
 - [x] Shanks' Baby Steps Giant Steps
 - [ ] Rho Pollard

```
BruteForce(g, h, p *big.Int) *big.Int
Shanks(g, h, p *big.Int) *big.Int 
```
