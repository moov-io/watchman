import { buildQueryString, parseQueryString } from "utils";

describe("parseQueryString", () => {
  it.each([
    [null, null],
    ["", ""],
    [undefined, undefined],
    ["?limit=1", { limit: "1" }],
    ["?limit=1&name=nicolas%20maduro", { limit: "1", name: "nicolas maduro" }]
  ])("should parse %s => %s", (input, expected) => {
    expect(parseQueryString(input)).toEqual(expected);
  });
});

describe("buildQueryString", () => {
  it.each([
    [null, null],
    ["", ""],
    [undefined, undefined],
    [{ limit: "1" }, "limit=1"],
    [{ limit: "1", name: "nicolas maduro" }, "limit=1&name=nicolas%20maduro"]
  ])("should parse %s => %s", (input, expected) => {
    expect(buildQueryString(input)).toEqual(expected);
  });
});
