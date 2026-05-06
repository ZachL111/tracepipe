# Tracepipe Walkthrough

This note is the quickest way to read the extra review model in `tracepipe`.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | span volume | 175 | ship |
| stress | latency skew | 185 | ship |
| edge | signal loss | 230 | ship |
| recovery | incident shape | 212 | ship |
| stale | span volume | 184 | ship |

Start with `edge` and `baseline`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

The useful comparison is `signal loss` against `span volume`, not the raw score alone.
