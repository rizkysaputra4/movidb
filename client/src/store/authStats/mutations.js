function capitalizeName(name) {
  return name.replace(/\b(\w)/g, (s) => s.toUpperCase());
}

export function updateDataAuth(state, data) {
  data.fullName = capitalizeName(data.fullName);
  state.data = data;
}
