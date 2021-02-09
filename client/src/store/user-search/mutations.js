export function updateSearchResult(state, result) {
  state.searchResult = result;
}

export function isNotAuthorized(state) {
  state.isAuthorized = false;
}
