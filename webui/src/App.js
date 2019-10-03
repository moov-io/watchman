import React, { useReducer } from "react";
import * as R from "ramda";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import Form from "Form";
import Results from "Results";
import { Container } from "Components";
import { buildQueryString, isNilOrEmpty } from "utils";
import { search } from "api";
import { createBrowserHistory } from "history";

const history = createBrowserHistory();

const reducer = (state, action) => {
  switch (action.type) {
    case "SEARCH_INIT":
      return R.pipe(
        R.assoc("results", null),
        R.assoc("loading", true),
        R.assoc("error", null)
      )(state);
    case "SEARCH_SUCCESS":
      return R.pipe(
        R.assoc("results", action.payload),
        R.assoc("loading", false),
        R.assoc("error", null)
      )(state);
    case "SEARCH_ERROR":
      return R.pipe(
        R.assoc("results", null),
        R.assoc("loading", false),
        R.assoc("error", action.payload)
      )(state);
    case "SEARCH_RESET":
      return initialState;
    default:
      return state;
  }
};
const initialState = {
  error: null,
  loading: false,
  results: null
};

const valuesOnlyContainLimit = R.pipe(
  R.filter(R.complement(isNilOrEmpty)),
  R.omit(["limit"]),
  R.isEmpty
);

function App() {
  const [state, dispatch] = useReducer(reducer, initialState);
  const executeSearch = async qs => {
    dispatch({ type: "SEARCH_INIT" });
    try {
      const payload = await search(qs);
      dispatch({ type: "SEARCH_SUCCESS", payload });
    } catch (err) {
      dispatch({ type: "SEARCH_ERROR", payload: err });
    }
  };

  const handleReset = () => {
    dispatch({ type: "SEARCH_RESET" });
    history.push({ ...history.location, search: "" });
  };

  const handleSubmit = values => {
    if (valuesOnlyContainLimit(values)) return;
    const qs = buildQueryString(values);
    history.push({ ...history.location, search: qs });
    executeSearch(qs);
  };

  return (
    <div
      css={`
        width: 80vw;
        margin: 1em auto;
      `}
    >
      <Container>
        <h1>OFAC Search</h1>
      </Container>
      <Form onSubmit={handleSubmit} onReset={handleReset} />
      <Results data={state} />
    </div>
  );
}

export default App;
