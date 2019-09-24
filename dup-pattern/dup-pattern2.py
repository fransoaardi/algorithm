def compress(word) -> dict:
    """
    compress 는 입력된 문자열을 압축한다
    """
    rtn = {}
    for char in word:
        rtn[char] = rtn.get(char, 0) + 1
    return rtn


def dup_possible(word, dictionary, max_dups) -> bool:
    """
    dup_possible 은 압축한 문자열의 중복 가능여부를 확인한다.
    'aabaca' 에서 'a' 의 중복을 찾는 경우 'a'가 4개 있으므로 'a' 의 중복확률이 있다.
    """
    compressed = compress(word)
    for i in compressed:
        if dictionary.get(i, 0) < max_dups+1 * compressed[i]:
            return False
    return True


def remove_duplicate(source: str) -> str:
    k = 1
    max_duplicate = 2    # 허용 duplicate

    compressed = compress(source)
    while k <= len(source)/(max_duplicate+1):  # 길이 k 인 단어를 가지고 비교
        _s = []
        for idx in range(0, len(source), k):
            if idx + k > len(source) or idx + (max_duplicate+1) * k > len(source):
                _s.extend(source[idx:])
                break
            if not dup_possible(source[idx+k], compressed, max_duplicate):
                _s.extend(source[idx:idx+k])
                continue
            if all(source[idx:idx+k] == source[idx+(m-1)*k:idx+m*k] for m in range(2, max_duplicate+2)):
                pass
            else:
                _s.extend(source[idx:idx+k])

        if len(_s) == len(source):
            k += 1
        else:
            compressed = compress(source)
        source = ''.join(_s)

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
