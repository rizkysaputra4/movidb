import axios from "axios";

export async function searchUser({ commit }, component) {
  console.log(component);
  const res = await axios.get(
    `${process.env.API}/admin/user/search?uid=${component.keyword}&limit=${component.limit}&offset=${component.offset}`,
    { withCredentials: true }
  );

  if (res.data.data) {
    commit("updateSearchResult", res.data.data.result);
    const pagination = {
      resultCount: res.data.data.count,
    };
  } else if (res.data.status === 401) {
    commit("isNotAuthorized");
  }

  return res;
}
