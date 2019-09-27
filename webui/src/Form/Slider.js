import React from 'react';
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import MSlider from "@material-ui/core/Slider";
import Typography from "@material-ui/core/Typography";

export default ({ label, ...props }) => (
  <div
    css={`
      width: 100%;
    `}
  >
    <Typography
      css={`
        && {
        color: rgba(0, 0, 0, 0.54);
      `}
    >
      {label}
    </Typography>
    <MSlider {...props} />
  </div>
);
