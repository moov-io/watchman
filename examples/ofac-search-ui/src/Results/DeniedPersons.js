import React from "react";
import * as R from "ramda";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import { matchToPercent } from "../utils";
import { Remarks } from "./Remarks";
import * as C from "../Components";

const l = R.toLower;

const Header = () => (
  <div
    css={`
      display: grid;
      grid-template-columns: 2fr 2fr 1fr 1fr 4em;
      margin-top: 1em;
    `}
  >
    <C.ResultHeader>Name</C.ResultHeader>
    <C.ResultHeader>Action</C.ResultHeader>
    <C.ResultHeader>Effective</C.ResultHeader>
    <C.ResultHeader>Expiration</C.ResultHeader>
    <C.ResultHeader>Match</C.ResultHeader>
  </div>
);

export const DeniedPersons = ({ data }) => {
  if (!data) return null;
  return (
    <C.Section>
      <C.SectionTitle>Denied Persons</C.SectionTitle>
      <Header />
      {data.length > 0 &&
        data.map(s => <DeniedPerson key={`${s.name}-${s.streetAddress}`} data={s} />)}
    </C.Section>
  );
};

export const DeniedPerson = ({ data }) => {
  // console.log("person: ", data);
  return (
    <div
      css={`
        padding-bottom: 1em;
        border-bottom: 1px solid #eee;
        &:last-of-type {
          border-bottom: 0px;
        }
        & > div {
          text-transform: capitalize;
        }
      `}
    >
      <div
        css={`
          display: grid;
          grid-template-columns: 2fr 2fr 1fr 1fr 4em;
          padding-top: 1em;
          & > div {
            margin-right: 1em;
          }
        `}
      >
        <div>{l(data.name)}</div>
        <div>{l(data.action)}</div>
        <div>{l(data.effectiveDate)}</div>
        <div>{l(data.expirationDate)}</div>
        <div>{matchToPercent(data.match)}</div>
      </div>
      <div
        css={`
          display: grid;
          grid-template-columns: 1fr;
          padding-top: 1em;
          & > div {
            margin-right: 1em;
          }
        `}
      />
      <Remarks
        grid="1fr"
        remarks={`${l(data.streetAddress)} ${l(data.city)}, ${l(data.state)} ${l(data.country)} ${l(
          data.postalCode
        )}`}
      />
    </div>
  );
};
