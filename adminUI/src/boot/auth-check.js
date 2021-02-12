import { Cookies } from "quasar";
var jwt = require("jsonwebtoken");
require("dotenv").config();

// "async" is optional;
// more info on params: https://quasar.dev/quasar-cli/boot-files
export default ({ router, ssrContext, redirect }) => {
  const cookies = process.env.SERVER ? Cookies.parseSSR(ssrContext) : Cookies;
  const token = cookies.get("Auth-Token");
  const salt = process.env.SALT;

  router.beforeEach((to, from, next) => {
    if (to.matched.some((record) => record.meta.requiresAuth)) {
      jwt.verify(token, salt, (err, decoded) => {
        if (err || decoded.role > 11) {
          redirect("/login");
        } else {
          next();
        }
      });
    } else {
      next();
    }
  });
};
