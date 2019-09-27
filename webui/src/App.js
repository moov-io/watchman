import React, { useReducer } from "react";
import * as R from "ramda";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import Form from "Form";
import Results from "Results";
import { Container } from "Components";
import { isNilOrEmpty } from "utils";
import { search } from "api";

const reducer = (state, action) => {
  switch (action.type) {
    case "SEARCH_SUCCESS":
      return R.assoc("results", action.payload)(state);
    case "SEARCH_RESET":
      return initialState;
    default:
      return state;
  }
};
const initialState = {
  results: []
};

function App() {
  const [state, dispatch] = useReducer(reducer, initialState);

  const handleReset = () => {
    dispatch({ type: "SEARCH_RESET" });
  };

  const handleSubmit = values => {
    const onlyLimit = R.pipe(
      R.filter(R.complement(isNilOrEmpty)),
      R.omit(["limit"]),
      R.isEmpty
    )(values);
    if (onlyLimit) return;

    const qs = R.pipe(
      R.filter(R.complement(isNilOrEmpty)),
      R.mapObjIndexed((val, key) => `${key}=${val}`),
      R.values,
      R.join("&")
    )(values);
    // console.log("qs: ", qs);
    search(qs).then(payload => dispatch({ type: "SEARCH_SUCCESS", payload }));
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
      <Results data={state.results} />
    </div>
  );
}

export default App;
