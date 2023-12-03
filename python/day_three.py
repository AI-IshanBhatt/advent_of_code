from collections import deque

with open("aoc_input_3") as f:
    lines = [l.strip() for l in f.readlines()]

    queue = deque()
    directions = [(-1,-1), (-1,0), (-1,1), (0,-1), (0,1), (1,-1), (1,0), (1,1)]
    digits = []
    symbols_table = []

    for i in range(len(lines)):
        j = 0
        current_digit = ""
        current_digit_location = []
        while j < len(lines[0]):
            if lines[i][j].isdigit():
                current_digit += lines[i][j]
                current_digit_location.append((i, j))
            elif lines[i][j] != ".":
                symbols_table.append((lines[i][j], (i, j)))
                digits.append((current_digit, current_digit_location))
                current_digit = ""
                current_digit_location = []
            else:
                digits.append((current_digit, current_digit_location))
                current_digit = ""
                current_digit_location = []
            j += 1
        if current_digit:
            digits.append((current_digit, current_digit_location))
            current_digit = ""
            current_digit_location = []

    digits = [d for d in digits if d[0]]
    symbols = [s[1] for s in symbols_table]

    # Start bfs for part 1
    sum_ = 0
    for digit, points in digits:
        flag = False
        for p in points:
            for r, c in directions:
                if (p[0] + r, p[1] + c) in symbols:
                    flag = True
        sum_ += int(digit) if flag else 0

    print(sum_)
    sum_ = 0

    # Part 2
    # Start bfs from symbols and try to get to numbers for that
    symbols = [s[1] for s in symbols_table if s[0] == "*"]
    digits_updated = {}
    for k,v in digits:
        for x in v:
            digits_updated[x] = k

    for p in symbols:
        ops = set()
        for r, c in directions:
            if (p[0] + r, p[1] + c) in digits_updated:
                ops.add(digits_updated[p[0] + r, p[1] + c])

        if len(ops) == 2:
            sum_ += (int(ops.pop())*int(ops.pop()))

    print(sum_)



