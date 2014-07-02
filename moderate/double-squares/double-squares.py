#!/usr/bin/env python

import sys


def gen_primes():
    blue_primes = []
    red_primes = []
    i = 0
    p = 0

    while p <= 46340:
        p = 4 * i + 3
        blue_primes.append(p)
        p = 4 * (i + 1) + 1
        red_primes.append(p)
        i += 1
    return blue_primes, red_primes


def r2(current):
    blue_prime_count = 0
    for prime in blue_primes:
        while current and current % prime == 0:
            current /= prime
            blue_prime_count += 1
        if current < prime ** 2:
            break
    if blue_prime_count % 2 != 0:
        return 0

    r2 = 4
    for prime in red_primes:
        tr = 0
        while current and current % prime == 0:
            current /= prime
            tr += 1
        r2 *= tr + 1
        if current < prime ** 2:
            break

    if r2 % 8 == 0:
        return r2 / 8
    else:
        return (r2 / 8) + 1


blue_primes, red_primes = gen_primes()

with open(sys.argv[1]) as _file:
    _file.readline()
    for line in _file.readlines():
        print r2(int(line))
