#!/bin/python3

translate = {
    "zero": "0",
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}

rtranslate = {"".join(reversed(key)): translate[key] for key in translate}


def convert_first_num_to_digit(line, tmap):
    x = ""
    for i in range(len(line)):
        # If we encounter a digit before a written out word, no need to convert
        if line[i].isdigit():
            return line

        x += line[i]

        for key in tmap:
            offset = x.find(key)
            if offset >= 0:
                return line[:offset] + tmap[key] + line[offset + len(key) :]

    return line


def sum_lines(lines):
    sum = 0
    for line in lines:
        # strip non-digits from input
        line = [ch for ch in line if ch.isdigit()]
        if len(line) > 0:
            sum += int(line[0] + line[-1])
    return sum


lines = open("input.txt", "r").readlines()

for idx in range(len(lines)):
    lines[idx] = convert_first_num_to_digit(lines[idx], translate)
    lines[idx] = convert_first_num_to_digit("".join(reversed(lines[idx])), rtranslate)
    lines[idx] = "".join(reversed(lines[idx]))

print(sum_lines(lines))

# print("Testing forward conversion")
# assert convert_first_num_to_digit("one1", translate) == "11"
# assert convert_first_num_to_digit("1one", translate) == "1one"
# assert convert_first_num_to_digit("fo41one", translate) == "fo41one"
# assert convert_first_num_to_digit("twone", translate) == "2ne"
# print("Testing reverse conversion")
# assert convert_first_num_to_digit("".join(reversed("98diseven")), rtranslate) == "7id89"
# assert convert_first_num_to_digit("".join(reversed("twone")), rtranslate) == "1wt"
# assert convert_first_num_to_digit("".join(reversed("twone")), rtranslate) == "1wt"
