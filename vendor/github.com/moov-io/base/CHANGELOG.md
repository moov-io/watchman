## v0.4.0 (Released 2019-01-11)

BREAKING CHANGES

- time: default times to UTC rather than Eastern.

## v0.3.1 (Released 2019-01-09)

- error: Add `ParseError` and `ErrorList` types.
- time: Prevent negative times in `NewTime(t time.Time)`

## v0.3.0 (Released 2019-01-07)

ADDITIONS

- Add ParseError and ErrorList. (See: [moov-io/base #23](https://github.com/moov-io/base/issues/23))

## v0.2.1 (Released 2019-01-03)

BUG FIXES

- http: Add OPTIONS to Access-Control-Allow-Methods

## v0.2.0 (Released 2018-12-18)

ADDITIONS

- Add `base.Time` as an embedded `time.Time` with banktime methods. (AddBankingDay, IsWeekend)

## v0.1.0 (Released 2018-12-17)

- Initial release
