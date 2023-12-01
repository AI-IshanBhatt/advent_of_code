
def decode_lines_and_sum(lines):
    sum_ = 0
    for line in lines:
        number = "".join([line[0], line[-1]])
        sum_ += int(number)

    return sum_


def problem_one(file_location: str):
    with open(file_location) as f:
        answer = []
        for line in f:
            digits = [c for c in line if c.isdigit()]
            answer.append(digits)

    return decode_lines_and_sum(answer)


def problem_two(file_location: str):
    words = {"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8",
             "nine": "9"}
    with open(file_location) as f:
        answer = []
        for line in f:
            digits = []
            for i in range(len(line)):
                if line[i].isdigit():
                    digits.append(line[i])
                else:
                    for w in words:
                        if w == line[i:i+len(w)]:
                            digits.append(words[w])
            answer.append(digits)

    return decode_lines_and_sum(answer)


print(problem_one("day_one_input"))
print(problem_two("day_one_input"))