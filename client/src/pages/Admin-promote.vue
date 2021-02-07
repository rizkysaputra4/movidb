<template>
  <q-no-ssr>
    <div class="q-pa-md">
      <q-table
        grid
        :card-container-class="cardContainerClass"
        title="Search User"
        :data="results"
        :columns="columns"
        row-key="name"
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
            <q-card>
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
    </div>
  </q-no-ssr>
</template>
<script>
import axios from "axios";
var results = [];
export default {
  data() {
    return {
      userName: "admin",
      offset: 0,
      results,
      loading: false,
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
    onSubmit(e) {
      e.preventDefault;
      this.loading = true;
      if (this.userName) {
        axios
          .get(
            `${process.env.API}/admin/user/search?uid=${this.userName}&limit=${this.pagination.rowsPerPage}&offset=0`,
            { withCredentials: true }
          )
          .then((res) => {
            console.log(res.data.data);
            if (res.data.data) {
              this.results = res.data.data.result;
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
      console.log(page);
      this.loading = true;
      if (this.userName) {
        axios
          .get(
            `${process.env.API}/admin/user/search?uid=${
              this.userName
            }&limit=${rowsPerPage}&offset=${(page - 1) * rowsPerPage}`,
            { withCredentials: true }
          )
          .then((res) => {
            if (res.data.data) {
              this.results = res.data.data.result;
              this.pagination.rowsNumber = res.data.data.count;
            } else {
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
  },

  mounted() {
    // get initial data from server (1st page)
    this.onSubmit({
      pagination: this.pagination,
      filter: undefined,
    });
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
