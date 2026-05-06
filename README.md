# tracepipe

`tracepipe` is a Go project in observability. Its focus is to reduce structured logs into trace spans and service latency summaries.

## Use Case

The project exists to keep a narrow engineering decision visible and testable. For this repo, that decision is how span volume and signal loss should influence a review result.

## Tracepipe Review Notes

`edge` and `baseline` are the cases worth reading first. They show the optimistic and cautious ends of the fixture.

## Highlights

- `fixtures/domain_review.csv` adds cases for span volume and latency skew.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/tracepipe-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `signal loss` and `span volume`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Code Layout

The implementation keeps the scoring rule plain: reward signal and confidence, preserve slack, penalize drag, then classify the result into a review lane.

The Go code keeps the review rule close to the tests.

## Run The Check

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Regression Path

The verifier is intentionally local. It should fail if the fixture score math, lane assignment, or language-specific test drifts.

## Future Work

The repository is intentionally scoped to local checks. I would expand it by adding adversarial fixtures before adding features.
