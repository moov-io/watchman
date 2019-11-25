const apiGet = async path => {
  const response = await fetch(path);
  const payload = await response.json();
  if (!response.ok) throw new Error(payload.error);
  return payload;
};

// grab the base URL/href (IE 11 and older don't support 'new URL(..)')
const url = new URL(document.baseURI);
const pathname = url => {
  return url.origin + url.pathname.replace(/\/$/, "");
};

export const search = async qs => apiGet(pathname(url) + `/search?${qs}`);

export const getSDNAlts = async sdnId => (await fetch(pathname(url) + `/ofac/sdn/${sdnId}/alts`)).json();
export const getSDNAddresses = async sdnId => (await fetch(pathname(url) + `/ofac/sdn/${sdnId}/addresses`)).json();
export const getSDNTypes = async qs => (await fetch(pathname(url) + `/ui/values/sdnType`)).json();
export const getOFACPrograms = async qs => (await fetch(pathname(url) + `/ui/values/ofacProgram`)).json();
