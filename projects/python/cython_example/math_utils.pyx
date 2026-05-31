# cython: language_level=3

def fast_factorial(int n):
    """A fast factorial implementation using Cython C types."""
    if n < 0:
        raise ValueError("Factorial not defined for negative numbers.")

    cdef long long result = 1
    cdef int i

    for i in range(1, n + 1):
        result *= i

    return result
