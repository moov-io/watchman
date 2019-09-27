import React from "react";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import MSelect from "@material-ui/core/Select";
import InputLabel from "@material-ui/core/InputLabel";
import MenuItem from "@material-ui/core/MenuItem";
import FormControl from "@material-ui/core/FormControl";

export default ({ label, id, options, ...props }) => (
  <FormControl
    css={`
      && {
        min-width: 100%;
      }
    `}
  >
    <InputLabel htmlFor={id}>{label}</InputLabel>
    <MSelect inputProps={{ name: id, id }} {...props}>
      {options.map(o => (
        <MenuItem key={o.name} value={o.val}>
          {o.name}
        </MenuItem>
      ))}
    </MSelect>
  </FormControl>
);
