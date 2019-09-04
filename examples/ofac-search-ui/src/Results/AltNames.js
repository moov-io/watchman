import React from "react";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import { matchToPercent } from "../utils";
import { Remarks } from "./Remarks";
import * as C from "Components";

export const AltNames = ({ data }) => {
  if (!data) return null;
  return (
    <C.Section>
      <C.SectionTitle>Alternate Names</C.SectionTitle>
      {data.length > 0 &&
        data.map(s => <AltName key={`${s.entityID}-${s.alternateID}`} data={s} />)}
    </C.Section>
  );
};

export const AltName = ({ data }) => {
  return (
    <div
      css={`
        padding-bottom: 1em;
        border-bottom: 1px solid #eee;
        &:last-of-type {
          border-bottom: 0px;
        }
        & > div {
          border: 0px solid red;
        }
      `}
    >
      <div
        css={`
          display: grid;
          grid-template-columns: 4em 2fr 1fr 1fr 4em;
          padding-top: 1em;
          & > div {
            margin-right: 1em;
          }
        `}
      >
        <div>{data.entityID}</div>
        <div>{data.alternateName}</div>
        <div />
        <div
          css={`
            text-transform: uppercase;
          `}
        >
          {data.alternateType}
        </div>
        <div>{data.match && matchToPercent(data.match)}</div>
      </div>
      <Remarks remarks={data.remarks} />
    </div>
  );
};
