"""
KMP class reference : https://gist.github.com/m00nlight/daa6786cc503fde12a77
"""


class KMP:
    def partial(self, pattern):
        """ Calculate partial match table: String -> [Int]"""
        ret = [0]

        for i in range(1, len(pattern)):
            j = ret[i - 1]
            while j > 0 and pattern[j] != pattern[i]:
                j = ret[j - 1]
            ret.append(j + 1 if pattern[j] == pattern[i] else j)
        return ret

    def search(self, T, P):
        """
        KMP search main algorithm: String -> String -> [Int]
        Return all the matching position of pattern string P in T
        """
        partial, ret, j = self.partial(P), [], 0

        for i in range(len(T)):
            while j > 0 and T[i] != P[j]:
                j = partial[j - 1]
            if T[i] == P[j]: j += 1
            if j == len(P):
                ret.append(i - (j - 1))
                j = partial[j - 1]

        return ret


def compress(word) -> dict:
    """
    compress 는 입력된 문자열을 압축한다
    """
    rtn = {}
    for char in word:
        rtn[char] = rtn.get(char, 0) + 1
    return rtn


def chunk(word, chunk_length, visited) -> list:
    """
    chunk 는 입력된 문자열을 chunk_length 로 끊고 그 중 visited 하지 않은 list 를 반환
    """
    rtn = []
    for i in range(len(word)-chunk_length):
        new_chunk = word[i:i+chunk_length]
        if new_chunk not in visited:
            rtn.append(word[i:i+chunk_length])
    return rtn


def check_dict_has_triple(word, dictionary) -> bool:
    """
    check_dict_has_triple 은 압축한 문자열의 3배수 존재여부를 확인한다.

    'aabaca' 에서 'a' 의 중복을 찾는 경우 'a'가 4개 있으므로 'a' 의 중복확률이 있다.
    """
    compressed = compress(word)
    for i in compressed:
        if dictionary.get(i, 0) < 3 * compressed[i]:
            return False
    return True


def remove_duplicate(source: str) -> str:
    visited = {}
    kmp = KMP()

    for k in range(int(len(source)/3)):
        compressed = compress(source)
        chunk_words = chunk(source, k+1, visited)
        for cw in chunk_words:
            if cw not in visited:
                if check_dict_has_triple(cw, compressed):
                    indices = kmp.search(source, cw)
                    # print(indices, cw)
                    dup_count = 1
                    deleted_ind = 0
                    for ind in range(1, len(indices)):
                        if indices[ind]-indices[ind-1] == len(cw):
                            # print(cw, ind, dup_count)
                            dup_count += 1
                        else:
                            dup_count = 1

                        if dup_count >= 3:
                            # print(source)
                            source = source[:indices[ind]-deleted_ind] + source[indices[ind]+len(cw)-deleted_ind:]
                            # print(source)
                            deleted_ind += len(cw)
                            dup_count = 2
                visited[cw] = True
    return source


if __name__ == '__main__':
    test_cases = [
        '신전신전신전에서 재밌는건 문자열 반복이고, 재밌지 않은건 재미가 없지',
        'AABBBCDDEEFFFFG',
        'AB123123123D',
        '123451234512346',
        "ngeeksngeeeksngeeeks",
        "nnnab",
        "신신신전신신신전신신신전",
        "이게 뭐야메롱 이게 뭐야메롱 이게 뭐야메롱 이게 뭐야메롱 ",
        "아이오아이의 너무너무너무너무 틀어줘",
        "반가움반가움반가움뽀뽀... 윙크사랑윙크... 윙크사랑뽀뽀...",
    ]

    for test in test_cases:
        result = remove_duplicate(test)
        print(test, '   ->  ', result)
