# Review Journal

The review surface for `tracepipe` is deliberately narrow: one fixture, one scoring rule, and one local check.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its observability focus without claiming live deployment or external usage.

## Cases

- `baseline`: `span volume`, score 175, lane `ship`
- `stress`: `latency skew`, score 185, lane `ship`
- `edge`: `signal loss`, score 230, lane `ship`
- `recovery`: `incident shape`, score 212, lane `ship`
- `stale`: `span volume`, score 184, lane `ship`

## Note

A future change should add new cases before it changes the scoring rule.
