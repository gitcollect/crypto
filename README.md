# crypto
Some cryptographic attacks, protocols, or other functions.

## log
Solving Discret Logarithm. Given a group G=\<g\>, |G|=p and h in G, finds a 'x' where g^x = h mod p.
 - **BruteForce(g, h, p *Int) *Int** : Brute-force attack.
 - **Shanks(g, h, p *Int) *Int** : Shanks' baby steps giant steps.
