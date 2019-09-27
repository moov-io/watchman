export const search = async qs => {
  const r = await fetch(`/api/search?${qs}`);
  return await r.json();
};

export const getSDNAlts = async sdnId => {
  const r = await fetch(`/api/sdn/${sdnId}/alts`);
  return await r.json();
};

export const getSDNAddresses = async sdnId => {
  const r = await fetch(`/api/sdn/${sdnId}/addresses`);
  return await r.json();
};
