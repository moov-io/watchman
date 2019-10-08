const apiGet = async path => {
  const response = await fetch(path);
  const payload = await response.json();
  if (!response.ok) throw new Error(payload.error);
  return payload;
};

// grab the base URL/href (IE 11 and older don't support 'new URL(..)')
const { pathname } = new URL(document.baseURI);

export const search = async qs => apiGet(pathname + `/search?${qs}`);

export const getSDNAlts = async sdnId => (await fetch(pathname + `/sdn/${sdnId}/alts`)).json();
export const getSDNAddresses = async sdnId => (await fetch(pathname + `/sdn/${sdnId}/addresses`)).json();
export const getSDNTypes = async qs => (await fetch(pathname + `/ui/values/sdnType`)).json();
export const getPrograms = async qs => (await fetch(pathname + `/ui/values/program`)).json();
