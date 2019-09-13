import * as R from 'ramda';

export const matchToPercent = match => `${(match * 100).toFixed(1)}%`;
export const isNilOrEmpty = R.anyPass([R.isNil, R.isEmpty]);

