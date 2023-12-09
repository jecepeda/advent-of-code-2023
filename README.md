# advent-of-code-2023

This repository contains the code, tests and solutions for the [Advent of Code](https://adventofcode.com/2023/) website.

## How to run

### Program

To run the code, run

```bash
make build; ./advent-of-code-2023 run
```

And you will have in your terminal an output that looks like this

```
+--------------+-------------+----------------+-------+
| TITLE        |        TIME |         RESULT | NOTES |
+--------------+-------------+----------------+-------+
| Day 1 Part 1 |   303.458µs |          55971 |       |
| Day 1 Part 2 |    1.8925ms |          54719 |       |
| Day 2 Part 1 |   783.917µs |           2278 |       |
| Day 2 Part 2 |   732.625µs |          67953 |       |
| Day 3 Part 1 |   983.917µs |         532445 |       |
| Day 3 Part 2 |   818.958µs |       79842967 |       |
| Day 4 Part 1 |   926.417µs |          22193 |       |
| Day 4 Part 2 |  1.017667ms |        5625994 |       |
| Day 5 Part 1 |   197.917µs |      579439039 |       |
| Day 5 Part 2 |   244.042µs |        7873084 |       |
| Day 6 Part 1 |    31.625µs |         741000 |       |
| Day 6 Part 2 | 43.411167ms |       38220708 |       |
| Day 7 Part 1 |  1.364958ms |      249483956 |       |
| Day 7 Part 2 |   1.51825ms |      252137472 |       |
| Day 8 Part 1 |   570.083µs |          11911 |       |
| Day 8 Part 2 |  2.875958ms | 10151663816849 |       |
| Day 9 Part 1 |   237.792µs |     1637452029 |       |
| Day 9 Part 2 |   192.417µs |            908 |       |
+--------------+-------------+----------------+-------+
```

### Tests

To test the code, just write in your terminal

```bash
make test
```

And the tests will be executed
