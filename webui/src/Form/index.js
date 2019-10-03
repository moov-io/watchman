import React, { useCallback, useEffect } from "react";
import * as R from "ramda";
import styled from "styled-components/macro"; // eslint-disable-line no-unused-vars
import MButton from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import * as C from "../Components";
import Select from "./Select";
import TextInput from "./TextInput";
import Slider from "./Slider";
import { countryOptionData, listOptionData } from "../data";
import { parseQueryString } from "utils";
import { useTypeOptions, useProgramOptions } from "./options";

const Button = styled(MButton)`
  margin: 1em;
`;

const ButtonSet = styled.div`
  display: flex;
  justify-content: flex-start;
  && > * {
    margin-right: 1em;
  }
`;

const Cell = styled.div`
  outline: 0px solid #eee;
  display: flex;
  align-items: flex-end;
  margin-bottom: 1em;
`;

const TwoColumns = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-gap: 1em 2em;
`;

const initialValues = {
  address: "",
  name: "",
  altName: "",
  city: "",
  state: "",
  providence: "",
  country: "",
  zip: "",
  limit: 10,
  q: "",
  sdnType: "",
  program: ""
  // disabled ///////////
  // idNumber: "",
  // list: "All",
  // score: 100
};

export default ({ onSubmit, onReset }) => {
  const [values, setValues] = React.useState(initialValues);

  const { values: typeOptionValues } = useTypeOptions();
  const { values: programOptionValues } = useProgramOptions();

  const handleChange = name => e => {
    const value = R.path(["target", "value"], e);
    setValues(values => R.assoc(name, value, values));
  };

  const handleChangeSlider = name => (e, value) => {
    setValues(values => R.assoc(name, value, values));
  };

  const handleSearchClick = () => {
    const activeValues = R.omit(["idNumber", "list", "score"])(values);
    onSubmit(activeValues);
  };

  const handleResetClick = () => {
    setValues(initialValues);
    onReset();
  };

  const submit = useCallback(onSubmit, []);
  useEffect(() => {
    const { search } = window.location;
    if (!search) {
      return;
    }
    setValues(values => {
      const newValues = R.mergeDeepRight(values, parseQueryString(search));
      submit(newValues);
      return newValues;
    });
  }, [submit]);

  return (
    <Container>
      <form
        onSubmit={e => {
          e.preventDefault();
          handleSearchClick();
        }}
      >
        <C.Section>
          <C.SectionTitle>Search</C.SectionTitle>
          <TwoColumns>
            <div>
              <Cell>
                <TextInput
                  label="Name | Alt | Address"
                  id="q"
                  value={values["q"]}
                  onChange={handleChange("q")}
                />
              </Cell>
              <Cell>
                <TextInput
                  label="Name"
                  id="name"
                  value={values["name"]}
                  onChange={handleChange("name")}
                />
              </Cell>
              <Cell>
                <TextInput
                  label="Alt Name"
                  id="altName"
                  value={values["altName"]}
                  onChange={handleChange("altName")}
                />
              </Cell>
              <Cell>
                <Select
                  label="Type"
                  id="sdnType"
                  value={values["sdnType"]}
                  onChange={handleChange("sdnType")}
                  options={typeOptionValues}
                />
              </Cell>
              <Cell>
                <Select
                  label="Program"
                  id="program"
                  value={values["program"]}
                  onChange={handleChange("program")}
                  options={programOptionValues}
                />
              </Cell>
              <Cell>
                <TextInput
                  type="number"
                  label="Limit"
                  id="limit"
                  value={values["limit"]}
                  onChange={handleChange("limit")}
                />
              </Cell>
            </div>
            <div>
              <Cell>
                <TextInput
                  label="Address"
                  id="address"
                  value={values["address"]}
                  onChange={handleChange("address")}
                />
              </Cell>
              <Cell>
                <TextInput
                  label="City"
                  id="city"
                  value={values["city"]}
                  onChange={handleChange("city")}
                />
              </Cell>
              <Cell>
                <TextInput
                  label="State"
                  id="state"
                  value={values["state"]}
                  onChange={handleChange("state")}
                />
              </Cell>
              <Cell>
                <TextInput
                  label="Providence"
                  id="providence"
                  value={values["providence"]}
                  onChange={handleChange("providence")}
                />
              </Cell>
              <Cell>
                <Select
                  label="Country"
                  id="country"
                  value={values["country"]}
                  onChange={handleChange("country")}
                  options={countryOptionData}
                />
              </Cell>
              <Cell>
                <TextInput
                  label="Postal Code"
                  id="zip"
                  value={values["zip"]}
                  onChange={handleChange("zip")}
                />
              </Cell>
            </div>
          </TwoColumns>
          <Cell>
            <ButtonSet>
              <Button variant="contained" color="primary" type="submit">
                Search
              </Button>
              <Button variant="outlined" color="default" onClick={handleResetClick}>
                Reset
              </Button>
            </ButtonSet>
          </Cell>
          {false && (
            <>
              <Cell>
                <Select
                  disabled={true}
                  label="List"
                  id="list"
                  value={values["list"]}
                  onChange={handleChange("list")}
                  options={listOptionData}
                />
              </Cell>
              <Cell>
                <Slider
                  disabled={true}
                  label="Score"
                  id="score"
                  value={values["score"]}
                  onChange={handleChangeSlider("score")}
                  min={0}
                  max={100}
                  valueLabelDisplay="auto"
                />
              </Cell>
            </>
          )}
        </C.Section>
      </form>
    </Container>
  );
};
