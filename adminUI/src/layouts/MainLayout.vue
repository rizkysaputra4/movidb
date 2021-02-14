<template>
  <q-layout
    view="hHh lpr fFf"
    :class="
      $q.dark.isActive ? 'header_dark bg-grey-9' : 'header_normal bg-grey-3'
    "
  >
    <q-header
      reveal
      elevated
      :class="
        $q.dark.isActive
          ? 'header_dark bg-blue-grey-10'
          : 'header_normal bg-green'
      "
    >
      <q-toolbar>
        <q-btn
          @click="left = !left"
          flat
          round
          dense
          icon="menu"
          class="q-mr-sm"
        />
        <!--          <q-avatar>-->
        <!--            <img src="https://cdn.quasar.dev/logo/svg/quasar-logo.svg">-->
        <!--          </q-avatar>-->

        <q-toolbar-title>Moviwiki Admin</q-toolbar-title>
        <q-btn
          class="q-mr-xs"
          flat
          round
          @click="$q.dark.toggle()"
          :icon="$q.dark.isActive ? 'nights_stay' : 'wb_sunny'"
        >
          <q-tooltip>Day/Night mode</q-tooltip>
        </q-btn>
        <q-btn flat round dense icon="search" class="q-mr-xs"
          ><q-tooltip>Search</q-tooltip></q-btn
        >
        <q-btn
          flat
          round
          dense
          icon="fas fa-sign-out-alt"
          @click="logoutNotify"
          to="/"
        />

        <AccountSetting />
      </q-toolbar>
    </q-header>
    <q-drawer
      class="left-navigation text-white bg-black"
      show-if-above
      v-model="left"
      side="left"
    >
      <div
        class="full-height"
        :class="$q.dark.isActive ? 'drawer_dark bg-grey-10' : 'drawer_normal'"
      >
        <div style="height: calc(100% - 20px)">
          <q-scroll-area style="height: 100%">
            <q-item
              active-class="tab-active"
              to="/dashboard"
              exact
              class="q-ma-sm navigation-item"
              clickable
              v-ripple
            >
              <q-item-section avatar>
                <q-icon name="dashboard" />
              </q-item-section>

              <q-item-section> Dashboard v1 </q-item-section>
            </q-item>

            <q-expansion-item
              expand-separator
              icon="attribution"
              label="Admin"
              style="padding-left: 10px; padding-right: 10px"
            >
              <q-item
                :header-inset-level="1"
                active-class="tab-active"
                to="/admin-list"
                exact
                class="q-ma-sm perm_identity"
                clickable
                v-ripple
              >
                <q-item-section avatar>
                  <q-icon name="person_search" />
                </q-item-section>

                <q-item-section> Admin List </q-item-section>
              </q-item>

              <q-item
                :header-inset-level="1"
                active-class="tab-active"
                to="/admin-promote"
                exact
                class="q-ma-sm perm_identity"
                clickable
                v-ripple
              >
                <q-item-section avatar>
                  <q-icon name="person_search" />
                </q-item-section>

                <q-item-section> Promote </q-item-section>
              </q-item>
            </q-expansion-item>

            <q-expansion-item
              expand-separator
              icon="attribution"
              label="User"
              style="padding-left: 10px; padding-right: 10px"
            >
              <q-item
                :header-inset-level="1"
                active-class="tab-active"
                to="/new-user"
                exact
                class="q-ma-sm perm_identity"
                clickable
                v-ripple
              >
                <q-item-section avatar>
                  <q-icon name="person_search" />
                </q-item-section>

                <q-item-section> Add User </q-item-section>
              </q-item>
            </q-expansion-item>

            <q-item
              active-class="tab-active"
              to="/profile"
              class="q-ma-sm navigation-item"
              clickable
              v-ripple
            >
              <q-item-section avatar>
                <q-icon name="contact_page" />
              </q-item-section>

              <q-item-section> My Profile </q-item-section>
            </q-item>
          </q-scroll-area>
        </div>
      </div>
    </q-drawer>

    <q-page-container>
      <q-page class="row no-wrap">
        <div class="col">
          <div class="full-height">
            <q-scroll-area class="col q-pr-sm full-height" visible>
              <router-view />
            </q-scroll-area>
          </div>
        </div>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
import AccountSetting from "../components/AccountSetting.vue";

export default {
  name: "Layout",
  components: {
    AccountSetting,
  },
  data() {
    return {
      left: false,
    };
  },
  methods: {
    logoutNotify() {
      this.$q.notify({
        message: "Logged out",
      });
    },
  },
};
</script>

<style>
.q-drawer {
  /*background-image: url(https://demos.creative-tim.com/vue-material-dashboard/img/sidebar-2.32103624.jpg) !important;*/
  background-size: cover !important;
}

.drawer_normal {
  background-color: rgba(1, 1, 1, 0.75);
}

.drawer_dark {
  background-color: #010101f2;
}

.navigation-item {
  border-radius: 5px;
}

.tab-active {
  background-color: green;
}

body {
  background: #f1f1f1 !important;
}

.header_normal {
  background: linear-gradient(
    145deg,
    rgb(32, 106, 80) 15%,
    rgb(21, 57, 102) 70%
  );
}

.header_dark {
  background: linear-gradient(145deg, rgb(61, 14, 42) 15%, rgb(14, 43, 78) 70%);
}
</style>
