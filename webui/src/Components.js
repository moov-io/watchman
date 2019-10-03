// import React from "react";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import MContainer from "@material-ui/core/Container";
import MCircularProgress from "@material-ui/core/CircularProgress";

export const Spinner = MCircularProgress;
export const Container = styled(MContainer)``;
export const Section = styled.section`
  border: 1px solid #eee;
  border-radius: 4px;
  padding: 1em 1em;
  margin-bottom: 1em;
`;

export const SectionTitle = styled.div`
  margin-bottom: 0.5em;
  font-size: 1.2em;
`;

export const Unknown = styled.span`
  color: #666;
`;

export const ResultHeader = styled.div`
  text-transform: uppercase;
  font-size: 0.8em;
  padding-bottom: 0.5em;
  border-bottom: 1px solid #eee;
`;
