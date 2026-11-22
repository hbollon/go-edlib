# Changelog

## [1.8.0](https://github.com/hbollon/go-edlib/compare/v1.7.0...v1.8.0) (2026-11-22)


### Features

* require Go 1.24 and pin toolchain to 1.27.4 ([3609e0d](https://github.com/hbollon/go-edlib/commit/3609e0d3f2925846b5204908115c739de5b7a25a))


### Bug Fixes

* **jaro:** round transpositions down and apply Winkler boost only above 0.7 ([#29](https://github.com/hbollon/go-edlib/issues/29)) ([d4b6078](https://github.com/hbollon/go-edlib/commit/d4b607803c1dd6a89b0a577617eb88b20876b14e)), closes [#28](https://github.com/hbollon/go-edlib/issues/28)
* **lcs:** count LCSEditDistance in runes for empty-string inputs ([#27](https://github.com/hbollon/go-edlib/issues/27)) ([51e5abe](https://github.com/hbollon/go-edlib/commit/51e5abe82344f263f3d98793317b94ee178bbca9))

## [1.7.0](https://github.com/hbollon/go-edlib/compare/v1.6.0...v1.7.0) (2025-08-19)


### Features

* reduce OSADamerau–Levenshtein space complexity from O(m*n) to O(3 * min(m,n)) ([#17](https://github.com/hbollon/go-edlib/issues/17)) ([37da2e1](https://github.com/hbollon/go-edlib/commit/37da2e108e4e0edc01b3b63f74240925aa359aa6))


### Bug Fixes

* **v1/algo:** invalid q-gram similarity calculation ([#22](https://github.com/hbollon/go-edlib/issues/22)) ([12ac3d7](https://github.com/hbollon/go-edlib/commit/12ac3d78fb78e00261e0ac4aa2951368e5f82941))

## [1.5.0](https://www.github.com/hbollon/go-edlib/compare/v1.4.0...v1.5.0) (2021-11-21)


### Features

* add k-gram shingle to Jaccard/Cosine sim ([#11](https://www.github.com/hbollon/go-edlib/issues/11)) ([24d61a6](https://www.github.com/hbollon/go-edlib/commit/24d61a62ee09eef2a930322cfbb0d4f13f13c5ca))
