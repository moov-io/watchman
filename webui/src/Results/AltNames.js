import React from "react";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import { matchToPercent } from "../utils";
import { Remarks } from "./Remarks";
import * as C from "Components";

export const Header = ({ withMatch = true }) => (
  <div
    css={`
      margin-top: 1em;
      width: 100%;
      display: grid;
      grid-template-columns: ${withMatch ? `4em 3fr 1fr 4em` : `4em 4fr 1fr`};
    `}
  >
    <C.ResultHeader>ID</C.ResultHeader>
    <C.ResultHeader>Alternate Name</C.ResultHeader>
    <C.ResultHeader>Type</C.ResultHeader>
    {withMatch && <C.ResultHeader>Match</C.ResultHeader>}
  </div>
);

export const AltNames = ({ data }) => {
  if (!data) return null;
  return (
    <C.Section>
      <C.SectionTitle>Alternate Names</C.SectionTitle>
      <Header />
      {data.length > 0 &&
        data.map(s => <AltName key={`${s.entityID}-${s.alternateID}`} data={s} />)}
    </C.Section>
  );
};

export const AltName = ({ data, displayId = "entityID" }) => {
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
          grid-template-columns: ${data.match ? `4em 3fr 1fr 4em` : `4em 4fr 1fr`};
          padding-top: 1em;
          & > div {
            margin-right: 1em;
          }
        `}
      >
        <div>{data[displayId]}</div>
        <div>{data.alternateName}</div>
        <div
          css={`
            text-transform: uppercase;
          `}
        >
          {data.alternateType}
        </div>
        {data.match && <div>{matchToPercent(data.match)}</div>}
      </div>
      <Remarks remarks={data.remarks} />
    </div>
  );
};
