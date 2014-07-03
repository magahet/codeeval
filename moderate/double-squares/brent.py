#!/usr/bin/env python

import random
import sys
from fractions import gcd


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


print pollardRho(int(sys.argv[1]))
