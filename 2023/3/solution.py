#!/bin/python3

import copy


# lines = open("input_small.txt", "r").readlines()
lines = open("input.txt", "r").readlines()


def print_mat(mat):
    for row in mat:
        print("".join(row))


def grab_serial(mat, row, col):
    """
    This function will extract the serial number whose digit
    was found at mat[row][col], OVERWRITE WITH '.' chars, and
    return the integer form of the serial number.
    """
    start_col = col
    for x in range(col):
        if not mat[row][start_col - 1].isdigit():
            break
        start_col = start_col - 1

    serial = []
    for c in range(start_col, len(mat[row])):
        if not mat[row][c].isdigit():
            break
        serial.append(mat[row][c])
        mat[row][c] = "."

    print(f"Found serial {int(''.join(serial))}")
    return int("".join(serial))


def find_serials(mat, row, col):
    serials = []
    # print(f"Found symbol {mat[row][col]} at ({row}, {col})")
    for rchk in range(row - 1, row + 2):
        if rchk < 0 or rchk >= len(mat):
            continue
        for cchk in range(col - 1, col + 2):
            if cchk < 0 or cchk >= len(mat[row]):
                continue
            elem = mat[rchk][cchk]
            if elem.isdigit():
                serials.append(grab_serial(mat, rchk, cchk))

    return serials


def sum_serials(mat):
    ans = 0

    # Iterate over matrix looking for symbols
    for row in range(len(mat)):
        for col in range(len(mat[row])):
            elem = mat[row][col]
            if elem != "." and not elem.isdigit():
                # There's a symbol here...do the calc
                print(f"Found symbol {mat[row][col]} at ({row}, {col})")
                serials = find_serials(mat, row, col)
                print(f"Found adjacent serials {serials}")
                ans += sum(serials)
                print(f"Running sum: {ans}")
                print_mat(mat)

    return ans


mat = []
for line in lines:
    line = line.strip()  # Strip newline

    if len(line) > 0:
        mat.append([x for x in line])

print_mat(mat)
print("PART 1")
print(f"FINAL SUM: {sum_serials(mat)}")


def grab_gear_component(mat, row, col):
    """
    This function will extract the gear component whose digit
    was found at mat[row][col], return the integer form of it.
    """
    start_col = col
    for x in range(col):
        if not mat[row][start_col - 1].isdigit():
            break
        start_col = start_col - 1

    serial = []
    for c in range(start_col, len(mat[row])):
        if not mat[row][c].isdigit():
            break
        serial.append(mat[row][c])
        mat[row][c] = "."

    print(f"Found gear component {int(''.join(serial))}")
    return int("".join(serial))


def find_gear_ratio(mat, row, col):
    serials = []
    # print(f"Found symbol {mat[row][col]} at ({row}, {col})")
    for rchk in range(row - 1, row + 2):
        if rchk < 0 or rchk >= len(mat):
            continue
        for cchk in range(col - 1, col + 2):
            if cchk < 0 or cchk >= len(mat[row]):
                continue
            elem = mat[rchk][cchk]
            if elem.isdigit():
                serials.append(grab_gear_component(copy.copy(mat), rchk, cchk))

    if len(serials) == 2:
        return serials[0] * serials[1]
    return 0


def sum_gear_ratios(mat):
    ans = 0

    # Iterate over matrix looking for symbols
    for row in range(len(mat)):
        for col in range(len(mat[row])):
            elem = mat[row][col]
            if elem == "*":
                # There's a gear here...do the calc
                print(f"Found possible gear at ({row}, {col})")
                ratio = find_gear_ratio(mat, row, col)
                print(f"Found ratio {ratio}")
                ans += ratio
                print(f"Running gear ratio: {ans}")

    return ans


mat = []
for line in lines:
    line = line.strip()  # Strip newline

    if len(line) > 0:
        mat.append([x for x in line])

print_mat(mat)
print("PART 2")
print(f"FINAL GEAR RATIO: {sum_gear_ratios(mat)}")
