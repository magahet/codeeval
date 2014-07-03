#!/usr/bin/env python

import random
import sys
from fractions import gcd


def is_prime(n):
    '''Deterministic Miller-Rabin. Adapted from implementation on Rosetta Code'''
    
    def _try_composite(a, d, n, s):
        if pow(a, d, n) == 1:
            return False
        for i in range(s):
            if pow(a, 2**i * d, n) == n-1:
                return False
        return True # n  is definitely composite
    
    if n in (0, 1, 2, 3):
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


def pollardRho(N):
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


def brent(N):
        if N % 2 == 0:
                return 2
        y, c, m = random.randint(
            1, N - 1), random.randint(1, N - 1), random.randint(1, N - 1)
        g, r, q = 1, 1, 1
        while g == 1:
                x = y
                for i in range(r):
                        y = ((y * y) % N + c) % N
                k = 0
                while (k < r and g == 1):
                        ys = y
                        for i in range(min(m, r - k)):
                                y = ((y * y) % N + c) % N
                                q = q * (abs(x - y)) % N
                        g = gcd(q, N)
                        k = k + m
                r = r * 2
        if g == N:
                while True:
                        ys = ((ys * ys) % N + c) % N
                        g = gcd(abs(x - ys), N)
                        if g > 1:
                                break
        return g


n = int(sys.argv[1])

while n > 1:
    prime = pollardRho(n)
    if not is_prime(prime):
        continue
    while n % prime == 0:
        color = 'b' if (prime - 3) % 4 == 0 else 'r'
        print prime, color
        n /= prime
    
