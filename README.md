<h1 align="center">Go-edlib : Edit distance and string comparison library</h1>

<p align="center">
  <a href='https://coveralls.io/github/hbollon/go-edlib?branch=master'>
    <img src='https://coveralls.io/repos/github/hbollon/go-edlib/badge.svg?branch=master' alt='Coverage Status' />
  </a>
  <a href="https://goreportcard.com/report/github.com/hbollon/go-edlib" target="_blank">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/hbollon/go-edlib" />
  </a>
  <a href="https://github.com/hbollon/go-edlib/blob/master/LICENSE.md" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/hbollon/go-edlib" target="_blank">
    <img src="https://pkg.go.dev/badge/github.com/hbollon/go-edlib" alt="PkgGoDev">
  </a>
</p>

> Golang string comparison and edit distance algorithms library featuring : Levenshtein, LCS, Hamming, Damerau levenshtein (OSA and Adjacent transpositions algorithms), Jaro-Winkler, Cosine, etc...

---

## Table of Contents

- [Requirements](#requirements)
- [Introduction](#introduction)
- [Features](#features)
- [Unicode support](#unicode-support)
- [Installation](#installation)
- [Benchmarks](#benchmarks)
- [Documentation](#documentation)
- [Examples](#examples)
- [Author](#author)
- [Contributing](#-contributing)
- [License](#-license)

---

## Requirements

- [Go](https://golang.org/doc/install) (v1.24+)

## Introduction

Golang open-source library which includes most (and soon all) edit-distance and string comparision algorithms with some extra! <br>
Works on Unicode code points, so non-ASCII text is supported out of the box (see [Unicode support](#unicode-support))!<br>
This library is 100% test covered 😁

## Features

- [Levenshtein](https://en.wikipedia.org/wiki/Levenshtein_distance)
- [LCS](https://en.wikipedia.org/wiki/Longest_common_subsequence_problem) (Longest common subsequence) with edit distance, backtrack and diff functions
- [Hamming](https://en.wikipedia.org/wiki/Hamming_distance)
- [Damerau-Levenshtein](https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance), with following variants:
  - OSA (Optimal string alignment)
  - Adjacent transpositions
- [Jaro & Jaro-Winkler](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance) similarity algorithms
- [Cosine Similarity](https://en.wikipedia.org/wiki/Cosine_similarity)
- [Jaccard Index](https://en.wikipedia.org/wiki/Jaccard_index)
- [QGram](https://en.wikipedia.org/wiki/N-gram)
- [Sorensen-Dice](https://en.wikipedia.org/wiki/S%C3%B8rensen%E2%80%93Dice_coefficient)
- Computed similarity percentage functions based on all available edit distance algorithms in this lib
- Fuzzy search functions based on edit distance with unique or multiples strings output
- Unicode support (rune-based comparisons)

## Unicode support

All algorithms operate on Unicode code points: input strings are converted to `[]rune` before being compared, so non-ASCII text (accented letters, CJK, emoji…) is handled correctly and distances and lengths are counted in code points, not bytes.

However, code points are compared as-is: go-edlib doesn't apply any Unicode-aware preprocessing. Depending on your use case, you may need to take care of the following before calling the library:

- **Normalization:** the same visible text can be encoded with different code point sequences. For example, `é` can be a single precomposed code point (`U+00E9`) or an `e` followed by a combining acute accent (`U+0065 U+0301`). Both render identically, but go-edlib sees them as different strings. Normalizing both inputs to the same form (usually NFC), for instance with [`golang.org/x/text/unicode/norm`](https://pkg.go.dev/golang.org/x/text/unicode/norm), solves this:

  ```go
  import "golang.org/x/text/unicode/norm"

  a := "caf\u00e9"  // "café" with a precomposed é
  b := "cafe\u0301" // "café" with e + combining acute accent

  edlib.LevenshteinDistance(a, b)                                   // 2
  edlib.LevenshteinDistance(norm.NFC.String(a), norm.NFC.String(b)) // 0
  ```

- **Grapheme clusters:** a user-perceived character can be made of several code points: letters with combining marks that have no precomposed form, emoji with skin tone modifiers, flags, or sequences such as 👨‍👩‍👧 (5 code points). Each code point counts as one unit, so editing such a character may cost more than 1, and string lengths used to compute similarity indexes are code point counts. Comparing strings grapheme by grapheme isn't supported; if you need to inspect user-perceived characters, use a segmentation library such as [`github.com/rivo/uniseg`](https://github.com/rivo/uniseg).

- **Case:** comparisons are case-sensitive (`"Café"` and `"café"` have a Levenshtein distance of 1). Lowercase or case-fold your inputs (e.g. with [`golang.org/x/text/cases`](https://pkg.go.dev/golang.org/x/text/cases)) if needed.

## Benchmarks

You can check an interactive Google chart with few benchmark cases for all similarity algorithms in this library through **StringsSimilarity** function [here](http://benchgraph.codingberg.com/q5)

However, if you want or need more details, you can also viewing benchmark raw output [here](https://github.com/hbollon/go-edlib/blob/master/tests/outputs/benchmarks.txt), which also includes memory allocations and test cases output (similarity result and errors).

If you are on Linux and want to run them on your setup, you can run `./tests/benchmark.sh` script.

## Installation

Open bash into your project folder and run:

```bash
go get github.com/hbollon/go-edlib
```

And import it into your project:

```go
import (
	"github.com/hbollon/go-edlib"
)
```

### Run tests

If you are on Linux and want to run all unit tests just run `./tests/tests.sh` script.

For Windows users you can run:

```bash
go test ./... # Add desired parameters to this command if you want
```

## Documentation

**You can find all the documentation here :** [Documentation](https://godoc.org/github.com/hbollon/go-edlib)

## Examples

### Calculate string similarity index between two string

You can use `StringSimilarity(str1, str2, algorithm)` function.
**algorithm** parameter must one of the following constants:

```go
// Algorithm identifiers
const (
	Levenshtein Algorithm = iota
	DamerauLevenshtein
	OSADamerauLevenshtein
	Lcs
	Hamming
	Jaro
	JaroWinkler
	Cosine
)
```

Example with levenshtein:

```go
res, err := edlib.StringsSimilarity("string1", "string2", edlib.Levenshtein)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("Similarity: %f", res)
}
```

### Execute fuzzy search based on string similarity algorithm

#### 1. Most matching unique result without threshold

You can use `FuzzySearch(str, strList, algorithm)` function.

```go
strList := []string{"test", "tester", "tests", "testers", "testing", "tsting", "sting"}
res, err := edlib.FuzzySearch("testnig", strList, edlib.Levenshtein)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("Result: %s", res)
}

```

```
Result: testing
```

#### 2. Most matching unique result with threshold

You can use `FuzzySearchThreshold(str, strList, minSimilarity, algorithm)` function.

```go
strList := []string{"test", "tester", "tests", "testers", "testing", "tsting", "sting"}
res, err := edlib.FuzzySearchThreshold("testnig", strList, 0.7, edlib.Levenshtein)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("Result for 'testnig': %s", res)
}

res, err = edlib.FuzzySearchThreshold("hello", strList, 0.7, edlib.Levenshtein)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("Result for 'hello': %s", res)
}

```

```
Result for 'testnig': testing
Result for 'hello':
```

#### 3. Most matching result set without threshold

You can use `FuzzySearchSet(str, strList, resultQuantity, algorithm)` function.

```go
strList := []string{"test", "tester", "tests", "testers", "testing", "tsting", "sting"}
res, err := edlib.FuzzySearchSet("testnig", strList, 3, edlib.Levenshtein)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("Results: %s", strings.Join(res, ", "))
}

```

```
Results: testing, test, tester
```

#### 4. Most matching result set with threshold

You can use `FuzzySearchSetThreshold(str, strList, resultQuantity, minSimilarity, algorithm)` function.

```go
strList := []string{"test", "tester", "tests", "testers", "testing", "tsting", "sting"}
res, err := edlib.FuzzySearchSetThreshold("testnig", strList, 3, 0.5, edlib.Levenshtein)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("Result for 'testnig' with '0.5' threshold: %s", strings.Join(res, " "))
}

res, err = edlib.FuzzySearchSetThreshold("testnig", strList, 3, 0.7, edlib.Levenshtein)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("Result for 'testnig' with '0.7' threshold: %s", strings.Join(res, " "))
}

```

```
Result for 'testnig' with '0.5' threshold: testing test tester
Result for 'testnig' with '0.7' threshold: testing
```

### Get raw edit distance (Levenshtein, LCS, Damerau–Levenshtein, Hamming)

You can use one of the following function to get an edit distance between two strings :

- [LevenshteinDistance](https://pkg.go.dev/github.com/hbollon/go-edlib#LevenshteinDistance)(str1, str2)
- [DamerauLevenshteinDistance](https://pkg.go.dev/github.com/hbollon/go-edlib#DamerauLevenshteinDistance)(str1, str2)
- [OSADamerauLevenshteinDistance](https://pkg.go.dev/github.com/hbollon/go-edlib#OSADamerauLevenshteinDistance)(str1, str2)
- [LCSEditDistance](https://pkg.go.dev/github.com/hbollon/go-edlib#LCSEditDistance)(str1, str2)
- [HammingDistance](https://pkg.go.dev/github.com/hbollon/go-edlib#HammingDistance)(str1, str2)

Example with Levenshtein distance:

```go
res := edlib.LevenshteinDistance("kitten", "sitting")
fmt.Printf("Result: %d", res)
```

```
Result: 3
```

### LCS, LCS Backtrack and LCS Diff

#### 1. Compute LCS(Longuest Common Subsequence) between two strings

You can use `LCS(str1, str2)` function.

```go
lcs := edlib.LCS("ABCD", "ACBAD")
fmt.Printf("Length of their LCS: %d", lcs)
```

```
Length of their LCS: 3
```

#### 2. Backtrack their LCS

You can use `LCSBacktrack(str1, str2)` function.

```go
res, err := edlib.LCSBacktrack("ABCD", "ACBAD")
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("LCS: %s", res)
}
```

```
LCS: ABD
```

#### 3. Backtrack all their LCS

You can use `LCSBacktrackAll(str1, str2)` function.

```go
res, err := edlib.LCSBacktrackAll("ABCD", "ACBAD")
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("LCS: %s", strings.Join(res, ", "))
}
```

```
LCS: ABD, ACD
```

#### 4. Get LCS Diff between two strings

You can use `LCSDiff(str1, str2)` function.

```go
res, err := edlib.LCSDiff("computer", "houseboat")
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("LCS: \n%s\n%s", res[0], res[1])
}
```

```
LCS Diff:
 h c o m p u s e b o a t e r
 + -   - -   + + + + +   - -
```

## Author

👤 **Hugo Bollon**

- Github: [@hbollon](https://github.com/hbollon)
- LinkedIn: [@Hugo Bollon](https://www.linkedin.com/in/hugo-bollon-68a2381a4/)
- Portfolio: [hugobollon.me](https://www.hugobollon.me)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/hbollon/go-edlib/issues).

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2020 [Hugo Bollon](https://github.com/hbollon).<br />
This project is [MIT License](https://github.com/hbollon/go-edlib/blob/master/LICENSE.md) licensed.
