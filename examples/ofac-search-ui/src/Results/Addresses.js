import React from "react";
import * as R from "ramda";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import { matchToPercent } from "../utils";
import { Remarks } from "./Remarks";
import * as C from "Components";

export const Addresses = ({ data }) => {
  if (!data) return null;
  return (
    <C.Section>
      <C.SectionTitle>Addresses</C.SectionTitle>
      {data.length > 0 && data.map(s => <Address key={`${s.entityID}-${s.addressID}`} data={s} />)}
    </C.Section>
  );
};

export const Address = ({ data }) => {
  // console.log("address: ", data);
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
            text-transform: capitalize;
          }
        `}
      >
        <div>{R.toLower(data.entityID)}</div>
        <div>{R.toLower(data.address)}</div>
        <div>{R.toLower(data.cityStateProvincePostalCode)}</div>
        <div>{R.toLower(data.country)}</div>
        <div>{data.match && matchToPercent(data.match)}</div>
      </div>
      <Remarks remarks={data.remarks} />
    </div>
  );
};
