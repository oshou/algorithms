def bfs(m, n):
    if (n == 1) return 1
    int ret = 0
    for i in range(m):
        ret += dfs(m-1, n-1)
    return ret
