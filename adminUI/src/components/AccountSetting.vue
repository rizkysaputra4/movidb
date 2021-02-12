<template>
  <q-btn-dropdown
    round
    flat
    fab-mini
    dense
    dropdown-icon="img:https://cdn.quasar.dev/img/boy-avatar.png"
  >
    <div class="row no-wrap q-pa-md">
      <div class="column">
        <div class="text-h6 q-mb-md">Settings</div>
        <q-toggle v-model="mobileData" label="Use Mobile Data" />
        <q-toggle v-model="bluetooth" label="Bluetooth" />
      </div>

      <q-separator vertical inset class="q-mx-lg" />

      <div class="column items-center">
        <q-avatar size="72px">
          <img src="https://cdn.quasar.dev/img/boy-avatar.png" />
        </q-avatar>

        <div class="text-subtitle1 q-mt-md q-mb-xs">
          {{ getAuthState.fullName }}
        </div>

        <q-btn
          color="primary"
          label="Logout"
          push
          size="sm"
          v-close-popup
          @click="logOut"
        />
      </div>
    </div>
    <q-tooltip>Account</q-tooltip>
  </q-btn-dropdown>
</template>

<script>
import axios from "axios";
import { mapGetters } from "vuex";

export default {
  data() {
    return {
      mobileData: true,
      bluetooth: false,
    };
  },
  methods: {
    logOut() {
      axios
        .get(`${process.env.API}/auth/logout`, { withCredentials: true })
        .then((res) => {
          this.$router.push("/login");
        });
    },
  },
  computed: {
    ...mapGetters("authStats", ["getAuthState"]),
  },
};
</script>
