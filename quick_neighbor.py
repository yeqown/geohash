from typing import List, Dict

odd = '''
b,c,f,g,u,v,y,z
8,9,d,e,s,t,w,x
2,3,6,7,k,m,q,r
0,1,4,5,h,j,n,p
'''

even = '''
p,r,x,z
n,q,w,y
j,m,t,v
h,k,s,u
5,7,e,g
4,6,d,f
1,3,9,c
0,2,8,b
'''


def build_matrix(src: str) -> List[List[str]]:
    mat = []
    for line in src.splitlines():
        line = line.strip('\n')
        line = line.strip(' ')

        if len(line) == 0:
            continue

        mat.append(line.split(','))

    return mat


def generate_mapping(mat: List[List[str]]) -> Dict[str, List[str]]:

    def _getter(x: int, y: int) -> bool:
        if x < 0 or y < 0 or y >= len(mat) or x >= len(mat[y]):
            return "overflow"

        return mat[y][x]

    def _build(x: int, y: int) -> List[str]:
        g = []
        g.append(_getter(x-1, y-1))  # NW
        g.append(_getter(x, y-1))  # N
        g.append(_getter(x+1, y-1))  # NE
        g.append(_getter(x+1, y))  # E
        g.append(_getter(x+1, y+1))  # SE
        g.append(_getter(x, y+1))  # S
        g.append(_getter(x-1, y+1))  # SW
        g.append(_getter(x-1, y))  # W
        return g

    mapping = {}
    for y in range(len(mat)):
        for x in range(len(mat[y])):
            mapping[mat[y][x]] = _build(x, y)

    return mapping


def main():
    mat = build_matrix(odd)
    gen = generate_mapping(mat)
    print(gen)

    mat = build_matrix(even)
    gen = generate_mapping(mat)
    print(gen)


if __name__ == '__main__':
    main()
