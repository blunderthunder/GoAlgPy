# Solution of reverse_bits 

## Algorithm

* get binary bits from uint number n
  * maks can be used to get binary bits
  * since we are using uint32, we can loop throught 32 bits by using masks in each position of bits
*  Shift 1 bit at a time `<<` from left which is equivalent to multiplying by 2^k
* after shifting all the bit from left , it will become the reverse of original bits

## solution 

| Language | Runtime | Memory  |
| :--------| :------ | :------ |
| python3  | 45 ms | 13.9 MB |
| go       | 0 ms  | 2.6 MB  |
| rust     | 0 ms   | 2 MB  |

> Go and rust are way better than python3, in terms of performance.
