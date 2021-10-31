# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [0.5.0] 2021-10-31

- Moved cmd/typer to gregoryv/typer package

## [0.4.1] 2021-10-23

- Updated dependencies

## [0.4.0] 2021-10-23

- Renamed RandomSentence to RandomStatement
- Renamed cmd/eng to cmd/english
- Optimized random functions, init slice of words

## [0.3.0] 2020-12-31

- Exposed word constants
- Added question words
- Remove types Language and Generator, simplifying the package.
- Improve randomness with seed from crypto/rand
- Colorize cmd/typer startup text
- Add adverbs and prepositions

## [0.2.0] 2020-12-30

- Add cmd/typer, typing speed game
- Add type Generator for generating sentences

## [0.1.0] 2020-12-30

- Add cmd/eng for querying the English language
