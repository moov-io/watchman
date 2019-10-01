import * as R from "ramda";

const capitalize = R.replace(/^./, R.toUpper);

const makeOption = v => ({ val: v, name: capitalize(v) });
export const makeOptionData = R.map(makeOption);

export const matchToPercent = match => `${(match * 100).toFixed(1)}%`;
export const isNilOrEmpty = R.anyPass([R.isNil, R.isEmpty]);

export const buildQueryString = v => {
  if (!v) return v;
  return R.pipe(
    R.filter(R.complement(isNilOrEmpty)),
    R.mapObjIndexed((val, key) => `${key}=${encodeURIComponent(val)}`),
    R.values,
    R.join("&")
  )(v);
};

export const parseQueryString = v => {
  if (!v) return v;
  return R.pipe(
    R.tail,
    R.split("&"),
    R.map(R.split("=")),
    R.fromPairs,
    R.map(decodeURIComponent)
  )(v);
};
