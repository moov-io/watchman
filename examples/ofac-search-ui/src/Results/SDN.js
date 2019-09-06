import React, { useReducer } from "react";
import * as R from "ramda";
import styled, { css } from "styled-components/macro"; // eslint-disable-line no-unused-vars
import { matchToPercent, isNilOrEmpty } from "utils";
import { Remarks } from "./Remarks";
import * as C from "Components";
import { getSDNAlts, getSDNAddresses } from "api";
import { SDNExpandDetails } from "./SDNDetails";

import MExpansionPanel from "@material-ui/core/ExpansionPanel";
import MExpansionPanelSummary from "@material-ui/core/ExpansionPanelSummary";
import MExpansionPanelDetails from "@material-ui/core/ExpansionPanelDetails";
import ExpandMoreIcon from "@material-ui/icons/ExpandMore";

const Header = () => (
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
        grid-template-columns: 4em 1fr 1fr 1fr 4em 36px;
      `}
    >
      <C.ResultHeader>ID</C.ResultHeader>
      <C.ResultHeader>Name</C.ResultHeader>
      <C.ResultHeader>Type</C.ResultHeader>
      <C.ResultHeader>Program</C.ResultHeader>
      <C.ResultHeader>Match</C.ResultHeader>
      <C.ResultHeader />
    </div>
  </div>
);

export const SDNS = ({ data }) => {
  if (!data) return null;
  return (
    <C.Section>
      <C.SectionTitle>Specially Designated Nationals</C.SectionTitle>
      <Header />
      {data && data.length > 0 && data.map(s => <SDN key={s.entityID} data={s} />)}
    </C.Section>
  );
};

const row = css`
  width: 100%;
  display: grid;
  padding-bottom: 1em;
  & > div {
    margin-right: 1em;
  }
`;

const ExpansionPanel = styled(MExpansionPanel)`
  && {
    box-shadow: unset;
    border-bottom: 0px solid #eee;
  }
  &&:last-child,
  &&:first-child {
    border-radius: 0;
  }
`;
const ExpansionPanelSummary = styled(MExpansionPanelSummary)`
  && {
    padding: 0;
  }
`;
const ExpansionPanelDetails = styled(MExpansionPanelDetails)`
  && {
    padding: 0;
  }
`;

const statusList = ["PRE_INIT", "INIT", "SUCCESS", "ERROR"];
const status = R.zipObj(statusList, statusList);

const initialState = {
  loaded: false,
  ALTS: {
    status: status.PRE_INIT,
    data: null
  },
  ADDS: {
    status: status.PRE_INIT,
    data: null
  }
};

const reducer = (state, action) => {
  // console.log("action: ", action);
  switch (action.type) {
    case status.INIT:
      return R.assocPath([action.api, "status"], action.type, state);
    case status.SUCCESS:
      return R.pipe(
        R.assoc("loaded", true),
        R.assocPath([action.api, "status"], action.type),
        R.assocPath([action.api, "data"], action.payload || [])
      )(state);
    // TODO
    //case status.ERROR:
    default:
      return state;
  }
};

export const SDN = ({ data }) => {
  const [details, dispatch] = useReducer(reducer, initialState);

  const handleExpandToggle = () => {
    if (details.loaded) return;

    dispatch({ api: "ALTS", type: status.INIT });
    getSDNAlts(data.entityID).then(alts =>
      dispatch({ api: "ALTS", type: status.SUCCESS, payload: alts })
    );

    dispatch({ api: "ADDS", type: status.INIT });
    getSDNAddresses(data.entityID).then(adds =>
      dispatch({ api: "ADDS", type: status.SUCCESS, payload: adds })
    );
  };

  if (isNilOrEmpty(data)) return null;
  return (
    <div>
      <ExpansionPanel onChange={handleExpandToggle}>
        <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
          <div
            css={`
              width: 100%;
            `}
          >
            <div
              css={`
                width: 100%;
                display: grid;
                grid-template-columns: 4em 1fr 1fr 1fr 4em;
                padding: 0.5em 0;
              `}
            >
              <div>{data.entityID}</div>
              <div>{data.sdnName}</div>
              <div
                css={`
                  text-transform: capitalize;
                `}
              >
                {data.sdnType || <C.Unknown>Unknown Type</C.Unknown>}
              </div>
              <div>{data.program}</div>
              <div>{matchToPercent(data.match)}</div>
            </div>

            {data.sdnType === "individual" && (
              <div
                css={`
                  ${row};
                  grid-template-columns: 4em 1fr;
                `}
              >
                <div />
                <div>{data.title}</div>
              </div>
            )}

            {data.sdnType === "vessel" && (
              <div
                css={`
                  ${row};
                  grid-template-columns: 4em 1fr 1fr 1fr 4em;
                `}
              >
                <div />
                <div>{data.vesselFlag || <C.Unknown>Unknown Flag</C.Unknown>}</div>
                <div>{data.vesselType || <C.Unknown>Unknown Type</C.Unknown>}</div>
                <div>{data.vesselOwner || <C.Unknown>Unknown Owner</C.Unknown>}</div>
              </div>
            )}
            <Remarks remarks={data.remarks} />
          </div>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails>
          <SDNExpandDetails data={details} />
        </ExpansionPanelDetails>
      </ExpansionPanel>
    </div>
  );
};
