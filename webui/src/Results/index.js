import React from "react";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import * as C from "../Components";
import { SDNS } from "./SDN";
import { AltNames } from "./AltNames";
import { Addresses } from "./Addresses";
import { DeniedPersons } from "./DeniedPersons";
import { isNilOrEmpty } from "utils";

const Results = ({ data }) => {
  if (isNilOrEmpty(data)) return null;
  return (
    <C.Container>
      <SDNS data={data.SDNs} />
      <AltNames data={data.altNames} />
      <Addresses data={data.addresses} />
      <DeniedPersons data={data.deniedPersons} />
    </C.Container>
  );
};

export default Results;
