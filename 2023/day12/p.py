def geometric_series_sum(a, r, n):
    return a * (1 - r ** n) / (1 - r)

print(geometric_series_sum(1, 8, 5))
