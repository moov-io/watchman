import * as R from "ramda";

export const matchToPercent = match => `${(match * 100).toFixed(1)}%`;
export const isNilOrEmpty = R.anyPass([R.isNil, R.isEmpty]);

export const buildQueryString = R.pipe(
  R.filter(R.complement(isNilOrEmpty)),
  R.mapObjIndexed((val, key) => `${key}=${val}`),
  R.values,
  R.join("&")
);

export const parseQueryString = R.pipe(
  R.tail,
  R.split("&"),
  R.map(R.split("=")),
  R.fromPairs
);
