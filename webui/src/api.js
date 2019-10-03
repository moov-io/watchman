const apiGet = async path => {
  const response = await fetch(path);
  const payload = await response.json();
  if (!response.ok) throw new Error(payload.error);
  return payload;
};

export const search = async qs => apiGet(`/search?${qs}`);

export const getSDNAlts = async sdnId => (await fetch(`/sdn/${sdnId}/alts`)).json();
export const getSDNAddresses = async sdnId => (await fetch(`/sdn/${sdnId}/addresses`)).json();
export const getSDNTypes = async qs => (await fetch(`/ui/values/sdnType`)).json();
export const getPrograms = async qs => (await fetch(`/ui/values/program`)).json();
