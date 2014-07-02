#!/usr/bin/env python

import sys


def gen_primes(maxn):
    blue_primes = []
    red_primes = []
    i = 0
    p = 0

    while p <= int(maxn ** 0.5):
        p = 4 * i + 3
        blue_primes.append(p)
        p = 4 * (i + 1) + 1
        red_primes.append(p)
        i += 1
    return blue_primes, red_primes


def r2(current):
    for prime in blue_primes:
        blue_prime_count = 0
        while current and current % prime == 0:
            current /= prime
            blue_prime_count += 1
        if blue_prime_count % 2 != 0:
            return 0
        if current < prime:
            break
        
    a0 = 0
    while current and current % 2 == 0:
        current /= 2
        a0 += 1

    b = 1
    for prime in red_primes:
        tr = 0
        while current and current % prime == 0:
            current /= prime
            tr += 1
        b *= tr + 1
        if current < prime:
            break

    if b == 1:
        return 0
    elif b % 2 == 0:
        return b / 2
    else:
        return (b - (-1 ** a0)) / 2


with open(sys.argv[1]) as _file:
    n_list = [int(l) for l in _file.readlines()]
    n_list.pop(0)
    
blue_primes, red_primes = gen_primes(max(n_list))
for n in n_list:
    print r2(n)