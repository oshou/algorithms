"""
1. 根ノードrootからスタート
2. 末端まで進めるまで進む
3. 末端についたら1つ前に戻ってまだ進んでいない分岐を進む
4. 進んでいない分岐がなかったらもう1つ前に戻ってまだ進んでいない分岐を進む
- 大きなグラフに対するDFSはスタックオーバーフローを起こす可能性あり
"""


def __logging(visited, rest=[]):
    if rest:
        print "visited:%s\n rest:%s\n" % (visited, rest)
    else:
        print "visited:%s" % visited


def dfs(graph, start, end, visited=[]):
    global is_found
    if is_found:
        return visited
    if start == end:
        is_found = True
        visited.append(start)
        return visited
    visited.append(start)
    for label in graph.get(start, []):
        if not
    stack = [start]
    visited = []
    while stack:
        lable = stack.pop(0)
        if label == end:
            visited.append(label)
            __logging(visited, stack)
            return visited
        if label not in visited:


if __name__ == '__main__':
    graph = {1: [2, 6, 8],
             2: [3, 4],
             3: [],
             4: [5],
             5: [1],
             6: [7],
             7: [],
             8: [9, 10],
             9: [7],
             10: [11],
             11: [],
             }
    print dfs(graph, 1, 10)
