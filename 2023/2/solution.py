#!/bin/python3


def is_pull_impossible(pull):
    for count in pull.split(","):
        num = int(count.split()[0])
        color = count.split()[1]
        if (
            (color == "red" and num > 12)
            or (color == "green" and num > 13)
            or (color == "blue" and num > 14)
        ):
            return True


def is_game_possible(game):
    for pull in game.split(";"):
        if is_pull_impossible(pull):
            return False
    return True


lines = open("input.txt", "r").readlines()

print("PART 1")

id_sum = 0
for line in lines:
    game = int(line.split(":")[0].split()[1])
    if is_game_possible(line.split(":")[1]):
        id_sum += game
        print(f"Game {game} possible {id_sum}")

print(id_sum)

print("PART 2")


def parse_pull(pull):
    parsed = {}
    for count in pull.split(","):
        parsed[count.split()[1]] = int(count.split()[0])
    return parsed


def game_power(game):
    bag = {}
    for pull in game.split(";"):
        parsed = parse_pull(pull)
        for key in parsed:
            if not bag.get(key) or bag.get(key) < parsed[key]:
                bag[key] = parsed[key]

    print(bag)
    power = 1
    for color in bag:
        power *= bag[color]

    print(power)
    return power


sum = 0
for line in lines:
    sum += game_power(line.split(":")[1])

print(sum)
