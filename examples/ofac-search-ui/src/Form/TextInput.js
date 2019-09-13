import React from "react";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import MTextField from "@material-ui/core/TextField";

export default ({ id, ...props }) => {
  return (
    <MTextField
      css={`
        && {
          min-width: 100%;
        }
      `}
      name={id}
      {...props}
    />
  );
};
