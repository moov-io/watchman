import React from "react";
import * as R from "ramda";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import { matchToPercent } from "../utils";
import { Remarks } from "./Remarks";
import * as C from "Components";

export const Header = ({ withMatch = true }) => (
  <div
    css={`
      margin-top: 1em;
      width: 100%;
    `}
  >
    <div
      css={`
        width: 100%;
        display: grid;
        grid-template-columns: ${withMatch ? `4em 2fr 1fr 1fr 4em` : `4em 2fr 1fr 1fr`};
      `}
    >
      <C.ResultHeader>ID</C.ResultHeader>
      <C.ResultHeader>Address</C.ResultHeader>
      <C.ResultHeader>City</C.ResultHeader>
      <C.ResultHeader>Country</C.ResultHeader>
      {withMatch && <C.ResultHeader>Match</C.ResultHeader>}
    </div>
  </div>
);

export const Addresses = ({ data }) => {
  if (!data) return null;
  return (
    <C.Section>
      <C.SectionTitle>Addresses</C.SectionTitle>
      <Header />
      {data.length > 0 && data.map(s => <Address key={`${s.entityID}-${s.addressID}`} data={s} />)}
    </C.Section>
  );
};

export const Address = ({ data, displayId = "entityID" }) => {
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
          grid-template-columns: ${data.match ? `4em 2fr 1fr 1fr 4em` : `4em 2fr 1fr 1fr`};
          padding-top: 1em;
          & > div {
            margin-right: 1em;
            text-transform: capitalize;
          }
        `}
      >
        <div>{R.toLower(data[displayId])}</div>
        <div>{R.toLower(data.address)}</div>
        <div>{R.toLower(data.cityStateProvincePostalCode)}</div>
        <div>{R.toLower(data.country)}</div>
        {data.match && <div>{matchToPercent(data.match)}</div>}
      </div>
      <Remarks remarks={data.remarks} />
    </div>
  );
};
