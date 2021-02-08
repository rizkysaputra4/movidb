<template>
  <q-no-ssr>
    <div class="q-pa-md">
      <q-table
        grid
        :card-container-class="cardContainerClass"
        title="Admin"
        :data="data"
        :columns="columns"
        row-key="name"
        :filter="filter"
        hide-header
        :pagination.sync="pagination"
        :rows-per-page-options="rowsPerPageOptions"
      >
        <template v-slot:top-right>
          <q-input
            borderless
            dense
            debounce="300"
            v-model="filter"
            placeholder="Search"
          >
            <template v-slot:append>
              <q-icon name="search" />
            </template>
          </q-input>
        </template>

        <template v-slot:item="props">
          <div class="q-pa-xs col-xs-12 col-sm-6 col-md-4">
            <q-card>
              <q-card-section class="text-center">
                <div class="row">
                  <div class="col">{{ props.row.country_id }}</div>
                  <div class="col">{{ props.row.user_name }}</div>
                  <div class="col">{{ props.row.is_online }}</div>
                  <div class="col">.col</div>
                </div>
              </q-card-section>
              <q-separator />
              <q-card-section
                class="flex flex-center"
                :style="{ fontSize: 30 - props.row.role + 'px' }"
              >
                <div>{{ props.row.user_full_name }}</div>
              </q-card-section>
              <q-card-section>
                <div>role: {{ props.row.role }}</div>
              </q-card-section>
            </q-card>
          </div>
        </template>
      </q-table>
    </div>
  </q-no-ssr>
</template>

<script>
const axios = require("axios");
var data = [];

export default {
  data() {
    return {
      filter: "",
      isAuthorized: true,
      pagination: {
        page: 1,
        rowsPerPage: this.getItemsPerPage(),
      },
      columns: [
        { name: "country_id", label: "Country", field: "country_id" },
        { name: "user_name", label: "UserName", field: "user_name" },
        { name: "is_online", label: "Online", field: "is_online" },
        { name: "user_full_name", label: "Name", field: "user_full_name" },
        { name: "role", label: "Role", field: "role" },
      ],
      data,
    };
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

    rowsPerPageOptions() {
      if (this.$q.screen.gt.xs) {
        return this.$q.screen.gt.sm ? [3, 6, 9] : [3, 6];
      }

      return [3];
    },
  },

  watch: {
    "$q.screen.name"() {
      this.pagination.rowsPerPage = this.getItemsPerPage();
    },
  },

  methods: {
    getItemsPerPage() {
      const { screen } = this.$q;
      if (screen.lt.sm) {
        return 3;
      }
      if (screen.lt.md) {
        return 6;
      }
      return 9;
    },
  },

  beforeMount() {
    axios
      .get(`${process.env.API}/admin/admin-list`, { withCredentials: true })
      .then((res) => {
        if (res.data.status === 401) {
          this.isAuthorized = false;
        }
        this.data = res.data.data;
      })
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
