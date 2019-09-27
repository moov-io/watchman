import React from "react";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars

export const Remarks = ({ remarks, grid = "4em 1fr" }) =>
  remarks ? (
    <div
      css={`
        display: grid;
        grid-template-columns: ${grid};
        color: #666;
        & > div {
          margin-right: 1em;
        }
      `}
    >
      <div />
      <div
        css={`
          font-size: 0.9em;
        `}
      >
        {remarks}
      </div>
    </div>
  ) : null;
