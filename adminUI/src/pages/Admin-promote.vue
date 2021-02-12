<template>
  <q-no-ssr>
    <div class="q-pa-md">
      <q-table
        grid
        :card-container-class="cardContainerClass"
        title="Search User"
        :data="getResult"
        :columns="columns"
        row-key="user_name"
        hide-header
        :pagination.sync="pagination"
        :loading="loading"
        @request="onRequest"
      >
        <template v-slot:top-right>
          <q-input
            borderless
            dense
            debounce="300"
            v-model="userName"
            placeholder="UserName"
            @keyup.enter="onSubmit"
          >
            <template v-slot:append>
              <q-icon name="search" @click.prevent="onSubmit" />
            </template>
          </q-input>
        </template>

        <template v-slot:item="props">
          <div class="q-pa-xs col-xs-12 col-sm-6 col-md-4">
            <q-card
              @click="onRowClick(props.row)"
              v-ripple
              class="q-hoverable cursor-pointer"
              :key="props.row.user_id"
              ><span class="q-focus-helper"></span>
              <q-card-section class="text-center">
                <div class="row">
                  <div class="col">User ID: {{ props.row.user_id }}</div>
                  <div class="col">UserName: {{ props.row.user_name }}</div>
                  <div class="col">Role: {{ props.row.role }}</div>
                </div>
              </q-card-section>
              <q-separator />
            </q-card>
          </div>
        </template>
      </q-table>

      <q-dialog v-model="prompt" persistent>
        <q-card style="min-width: 350px">
          <q-card-section>
            <div class="text-h6">User</div>
          </q-card-section>

          <q-card-section class="q-pt-none">
            <div class="col">User ID: {{ currentUser.user_id }}</div>
            <div class="col">UserName: {{ currentUser.user_name }}</div>
            <div class="col">Role: {{ currentUser.role }}</div>
            <div class="col">New Role: {{ newUserRole }}</div>
          </q-card-section>

          <q-card-section>
            <q-badge color="secondary"> Model: {{ newUserRole }} </q-badge>

            <q-slider
              v-model="newUserRole"
              :min="myRole + 1"
              :max="51"
              :step="1"
              label
              label-always
              color="light-green"
            />
          </q-card-section>

          <q-card-actions align="right" class="text-primary">
            <q-btn flat label="Cancel" v-close-popup />
            <q-btn
              flat
              label="Submit"
              v-close-popup
              @click.prevent="onSubmitNewRole()"
            />
          </q-card-actions>
        </q-card>
      </q-dialog>
    </div>
  </q-no-ssr>
</template>
<script>
import axios from "axios";
import { mapGetters, mapActions } from "vuex";

var results = [];
export default {
  data() {
    return {
      userName: "admin",
      offset: 0,
      results,
      loading: false,
      prompt: false,
      hover: false,
      currentUser: {},
      myRole: 99,
      newUserRole: 21,
      pagination: {
        sortBy: "desc",
        descending: false,
        page: 1,
        rowsPerPage: 30,
        rowsNumber: 10,
      },
      columns: [
        { name: "user_name", label: "UserName", field: "user_name" },
        { name: "user_id", label: "User ID", field: "user_id" },
        { name: "role", label: "Role", field: "role" },
      ],
    };
  },

  methods: {
    ...mapActions("userSearch", ["searchUser"]),

    onSubmit(e) {
      this.loading = true;
      if (this.userName) {
        var searchComp = {
          keyword: this.userName,
          limit: this.pagination.rowsPerPage,
          offset: (this.pagination.page - 1) * this.pagination.rowsPerPage,
        };
        this.searchUser(searchComp)
          .then((res) => {
            if (res.data.data) {
              this.pagination.rowsNumber = res.data.data.count;
            } else {
              this.results = [];
              this.$q.notify({
                type: "warning",
                multiLine: true,
                icon: "warning",
                message: "Username not found",
                position: "center",
                actions: [
                  {
                    label: "Dismiss",
                    color: "white",
                  },
                ],
              });
            }
            this.loading = false;
          })
          .catch((err) => console.log(err));
      } else {
        this.$q.notify({
          icon: "warning",
          type: "warning",
          message: "Keyword empty",
          position: "center",
          multiLine: true,
          actions: [
            {
              label: "Dismiss",
              color: "white",
            },
          ],
        });
        this.loading = false;
      }
    },

    getItemsPerPage() {
      const { screen } = this.$q;
      if (screen.lt.sm) {
        return 20;
      }
      if (screen.lt.md) {
        return 40;
      }
      return 60;
    },
    onRequest(props) {
      const { page, rowsPerPage, sortBy, descending } = props.pagination;
      console.log(props);
      console.log(page);
      this.loading = true;
      if (this.userName) {
        var searchComp = {
          keyword: this.userName,
          limit: rowsPerPage,
          offset: (page - 1) * rowsPerPage,
        };
        this.searchUser(searchComp)
          .then((res) => {
            if (res.data.data) {
              this.pagination.rowsNumber = res.data.data.count;
            }
            this.loading = false;
          })
          .catch((err) => console.log(err));
      } else {
        console.log("not executed", this.userName);
      }
      console.log("page", page);
      this.pagination.page = page;
      this.pagination.rowsPerPage = rowsPerPage;
      this.pagination.sortBy = sortBy;
      this.pagination.descending = descending;
    },
    onSubmitNewRole() {
      console.log(this.getResult);
      var data = {
        user_id: this.currentUser.user_id,
        role: this.newUserRole,
      };
      axios
        .put(`${process.env.API}/admin/admin-level`, data, {
          withCredentials: true,
        })
        .then((res) => {
          if (res.data.status === 403) {
            this.$q.notify({
              type: "warning",
              multiLine: true,
              icon: "warning",
              message: res.data.message,
              position: "center",
              actions: [
                {
                  label: "Dismiss",
                  color: "white",
                },
              ],
            });
          } else if (res.data.status === 401) {
            this.isNotAuthorized();
          } else if (res.data.status === 200) {
            this.$q.notify({
              message: "Success",
              color: "green",
            });
          }
          this.onSubmit({
            pagination: this.pagination,
            filter: undefined,
          });
        })
        .catch((err) => {
          console.log(err);
        });
    },
    onRowClick(row) {
      this.currentUser = row;
      this.prompt = true;
      console.log(this.myRole);
    },
  },
  computed: {
    cardContainerClass() {
      if (this.$q.screen.gt.xs) {
        return (
          "grid-masonry grid-masonry--" + (this.$q.screen.gt.sm ? "3" : "2")
        );
      }

      return void 0;
    },
    ...mapGetters("userSearch", ["getResult", "isAuthorized"]),
    // searchResultState: {
    //   get() {
    //     return this.$store.state.userSearch.searchResult;
    //   },
    // },
  },

  mounted() {
    // get initial data from server (1st page)
    // this.onSubmit({
    //   pagination: this.pagination,
    //   filter: undefined,
    // });
  },
  beforeMount() {
    axios
      .get(`${process.env.API}/public/my-role`, { withCredentials: true })
      .then((res) => (this.myRole = res.data.data.role))
      .catch((err) => console.log(err));
  },
  beforeUpdate() {
    if (!this.isAuthorized) {
      this.$router.push("/login");
    }
  },
};
</script>
<style lang="sass">


.grid-masonry
  flex-direction: column
  height: 700px

  &--2
    > div
      &:nth-child(2n + 1)
        order: 1
      &:nth-child(2n)
        order: 2

    &:before
      content: ''
      flex: 1 0 100% !important
      width: 0 !important
      order: 1
  &--3
    > div
      &:nth-child(3n + 1)
        order: 1
      &:nth-child(3n + 2)
        order: 2
      &:nth-child(3n)
        order: 3

    &:before,
    &:after
      content: ''
      flex: 1 0 100% !important
      width: 0 !important
      order: 2
</style>
