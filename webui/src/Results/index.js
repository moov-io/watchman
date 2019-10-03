import React from "react";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import * as C from "../Components";
import { SDNS } from "./SDN";
import { AltNames } from "./AltNames";
import { Addresses } from "./Addresses";
import { DeniedPersons } from "./DeniedPersons";
import { isNilOrEmpty } from "utils";

export default ({ data }) => {
  const { loading, error, results } = data;
  if (loading)
    return (
      <C.Container>
        <div
          css={`
            display: flex;
            justify-content: center;
          `}
        >
          <C.Spinner />
        </div>
      </C.Container>
    );
  if (error) return <C.Container>ERROR: {error.message}</C.Container>;
  if (isNilOrEmpty(results)) return null;
  return (
    <C.Container>
      <SDNS data={results.SDNs} />
      <AltNames data={results.altNames} />
      <Addresses data={results.addresses} />
      <DeniedPersons data={results.deniedPersons} />
    </C.Container>
  );
};
