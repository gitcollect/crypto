# Crypto
Some cryptographic attacks, protocols, or other functions.

## Factorization
Factorization algotirhm for composite integer.

 - [x] Brute force attack
 - [x] Rho Pollard
 - [ ] P-1 Pollard
 - [ ] Fermat

```Go
BruteForce(n *big.int) *big.Int
RhoPollard(n *big.Int) *big.Int
```

## LFSR
Linear-feedback shift register implementation. Register and feedback polynomial are represented by binary numbers.
 - [x] LFSR
 - [ ] A5/1
 - [ ] Berlekamp massey algorithm

**overview**

NewLFSR creates a Linead-feedback shift register, where the register and the feedback polynome are represented by bits of numbers.
```Go
NewLFSR(reg, taps *big.Int)(LFSR)
```
Inc cycles the register and returns the output bit.
```Go
(*LFSR) Inc() uint
```
Here is an example on how to create a LFSR.
```Go
// Numbers of taps
size := 16
// If the feedback polynomial is 1 + x¹¹ + x¹³ + x¹⁴ + x¹⁶, it is represented by 0000000000101101 = 0x2d.
// The "one" in the polynomial does not correspond to a tap, it corresponds to the input to the first bit
pol := "2d"
// Any nonzero start.
reg := "ace1"
// Generates *Int with our previous datas. We put '16' because our datas are hexadecimal numbers.
// Then creates the LFSR.
r, _ := new(big.Int).SetString(reg, 16)
p, _ := new(big.Int).SetString(pol, 16)
l := LFSR{Register: r, Taps: p, Size: size}
```

## Logarithm
Solving discret logarithm problem. Given a group **G** of order **p**, a generator **g** , an element **h**, finds **x** where **g^x = h mod p**.

 - [x] Brute force attack
 - [x] Shanks' Baby Steps Giant Steps
 - [ ] Rho Pollard

```Go
BruteForce(g, h, p *big.Int) *big.Int
Shanks(g, h, p *big.Int) *big.Int
```
