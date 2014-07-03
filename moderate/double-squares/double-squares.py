#!/usr/bin/env python

import sys
import random
from fractions import gcd


def get_prime_factor(N):
    '''Pollard'''
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


def is_prime(n, k=10):
    '''Miller-Rabin'''
    if n == 2 or n == 3:
        return True
    if not n & 1:
        return False

    def check(a, s, d, n):
        x = pow(a, d, n)
        if x == 1:
            return True
        for i in xrange(s - 1):
            if x == n - 1:
                return True
            x = pow(x, 2, n)
        return x == n - 1

    s = 0
    d = n - 1

    while d % 2 == 0:
        d >>= 1
        s += 1

    for i in xrange(k):
        a = random.randrange(2, n - 1)
        if not check(a, s, d, n):
            return False
    return True


def r2(current):
    a0 = 0
    blue_prime_count = 0
    b = 1

    while current > 1:
        prime = get_prime_factor(current)
        if not is_prime(prime):
            continue
        count = 0
        while current and current % prime == 0:
            count += 1
            current /= prime

        if prime == 2:
            a0 += count
        elif (prime - 3) % 4 == 0:
            blue_prime_count += count
        else:
            b *= count + 1

    if blue_prime_count % 2 != 0:
        return 0
    if b % 2 == 0:
        return b / 2
    else:
        return (b - (-1 ** a0)) / 2


with open(sys.argv[1]) as _file:
    n_list = [int(l) for l in _file.readlines()]
    n_list.pop(0)

for i, n in enumerate(n_list):
    print r2(n)
