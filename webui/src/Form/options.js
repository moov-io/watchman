import { useEffect, useReducer } from "react";
import * as R from "ramda";
import { getSDNTypes, getPrograms } from "api";
import { makeOptionData } from "utils";

const allOption = { val: "", name: "All" };

const buildUseOptions = api => () => {
  const initialState = { loading: false, values: [] };
  const reducer = (state, action) => {
    switch (action.type) {
      case "INIT":
        return R.assoc("loading", true)(state);
      case "SUCCESS":
        return R.assoc("values", action.payload)(state);
      default:
        return state;
    }
  };
  const [state, dispatch] = useReducer(reducer, initialState);
  useEffect(() => {
    dispatch({ type: "INIT" });
    api().then(payload => {
      dispatch({ type: "SUCCESS", payload: [allOption, ...makeOptionData(payload)] });
    });
  }, []);
  return state;
};

export const useTypeOptions = buildUseOptions(getSDNTypes);
export const useProgramOptions = buildUseOptions(getPrograms);
