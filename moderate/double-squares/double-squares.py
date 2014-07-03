#!/usr/bin/env python

import sys
import random
from fractions import gcd


def get_prime_factor(N):
    '''Pollard Rho Algorithm'''
    if N == 1:
        return 1
    if N % 2 == 0:
        return 2
    x = random.randint(1, N - 1)
    y = x
    c = random.randint(1, N - 1)
    g = 1
    while g == 1:
            x = ((x * x) % N + c) % N
            y = ((y * y) % N + c) % N
            y = ((y * y) % N + c) % N
            g = gcd(abs(x - y), N)
    return g

 
def is_prime(n):
    '''Deterministic Miller-Rabin. Adapted from implementation on Rosetta Code'''
    
    def _try_composite(a, d, n, s):
        if pow(a, d, n) == 1:
            return False
        for i in range(s):
            if pow(a, 2**i * d, n) == n-1:
                return False
        return True # n  is definitely composite
    
    if n in (1, 2, 3):
        return True
    d, s = n - 1, 0
    while not d % 2:
        d, s = d >> 1, s + 1
    # Returns exact according to http://primes.utm.edu/prove/prove2_3.html
    if n < 1373653: 
        return not any(_try_composite(a, d, n, s) for a in (2, 3))
    if n < 25326001: 
        return not any(_try_composite(a, d, n, s) for a in (2, 3, 5))
    if n < 118670087467: 
        if n == 3215031751: 
            return False
        return not any(_try_composite(a, d, n, s) for a in (2, 3, 5, 7))
    raise Exception('N is too large')


def r2(current):
    if current in (0, 1):
        return 1
    a0 = 0
    blue_prime_count = 0
    b = 0
    while current > 1:
        prime = get_prime_factor(current)
        if not is_prime(prime):
            continue
        count = 0
        while current % prime == 0:
            count += 1
            current /= prime
        if prime == 2:
            a0 += count
        elif (prime - 3) % 4 == 0 and count % 2 != 0:
            return 0
        elif (prime - 1) % 4 == 0:
            b = b * (count + 1) if b else count + 1
    if b % 2 == 0:
        return b / 2
    else:
        return (b - ((-1) ** a0)) / 2


with open(sys.argv[1]) as _file:
    n_list = [int(l) for l in _file.readlines()]
    n_list.pop(0)

for i, n in enumerate(n_list):
    print r2(n)