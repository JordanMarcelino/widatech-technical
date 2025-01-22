from typing import List


def get_number_combinations(l: int, t: int) -> List[List[int]]:
    results = []

    def dfs(start: int, combs: List[int], total: int) -> None:
        if len(combs) == l:
            if total == t:
                results.append(combs.copy())
            return

        for num in range(start, 10):
            if total + num > t:
                break

            combs.append(num)
            dfs(num + 1, combs, total + num)
            combs.pop()

    dfs(1, [], 0)
    return results


def main() -> None:
    tests = [
        (3, 6),
        (3, 8),
        (4, 5),
        (4, 11),
        (5, 12),
    ]

    for l, t in tests:
        print(get_number_combinations(l, t))


if __name__ == "__main__":
    main()
