import { Cookies } from "quasar";
var jwt = require("jsonwebtoken");
require("dotenv").config();

// "async" is optional;
// more info on params: https://quasar.dev/quasar-cli/boot-files
export default ({ router, ssrContext, redirect }) => {
  const cookies = process.env.SERVER ? Cookies.parseSSR(ssrContext) : Cookies;
  const token = cookies.get("Auth-Token");
  const salt = process.env.SALT;
  let auth;

  router.beforeEach((to, next) => {
    console.log(token);
    jwt.verify(token, salt, (err, decoded) => {
      auth = err ? false : true;
    });

    if (to.matched.some((record) => record.meta.requiresAuth)) {
      console.log(auth);
      if (!auth) {
        redirect("/login");
      } else {
        next();
      }
    } else {
      next();
    }
  });
};
