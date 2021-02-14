<template>
  <div class="q-pa-md q-mx-auto" style="width: 400px">
    <h4>Add New User Data</h4>

    <q-form @submit="onSubmit" @reset="onReset" class="q-gutter-md">
      <q-input
        filled
        v-model="userName"
        label="UserName *"
        hint="Minimum 4 character"
        debounce="600"
        :rules="[
          (val) => (val && val.length >= 4) || 'Please type something',
          isUserNameExist,
        ]"
      />

      <q-input
        filled
        v-model="email"
        label="Email *"
        debounce="600"
        :rules="[
          (val) => (val !== null && val !== '') || 'Please type your Email',
          isEmailValidAndNotUsed,
        ]"
      />

      <q-input
        filled
        type="password"
        v-model="password"
        hint="Minimum 6 character"
        debounce="500"
        label="Password *"
        :rules="[
          (val) => (val && val.length >= 6) || 'Please type your Password',
        ]"
      />

      <q-select
        filled
        v-model="country"
        :value="country"
        label="Country"
        use-input
        hide-selected
        fill-input
        input-debounce="0"
        :options="options"
        @filter="filterFn"
      >
        <template v-slot:no-option>
          <q-item>
            <q-item-section class="text-grey"> No results </q-item-section>
          </q-item>
        </template>
      </q-select>

      <q-input filled v-model="fullName" label="Full Name" />

      <q-input filled type="date" v-model="birthDate" hint="Birth Date" />

      <q-input v-model="bio" filled type="textarea" label="Bio" />

      <q-input filled v-model="facebookLink" label="Facebook Link" />

      <div>
        <q-btn label="Submit" type="submit" color="primary" />
        <q-btn
          label="Reset"
          type="reset"
          color="primary"
          flat
          class="q-ml-sm"
        />
      </div>
    </q-form>
  </div>
</template>

<script>
import axios from "axios";
var today = new Date();
var dd = String(today.getDate()).padStart(2, "0");
var mm = String(today.getMonth() + 1).padStart(2, "0"); //January is 0!
var yyyy = today.getFullYear();

today = mm + "/" + dd + "/" + yyyy;
let stringOptions = [];
export default {
  data() {
    return {
      userName: null,
      email: null,
      password: null,
      country: null,
      fullName: null,
      birthDate: today,
      bio: null,
      facebookLink: null,
      options: stringOptions,
      countries: null,
    };
  },

  methods: {
    onSubmit() {
      var countryIndex = this.options.indexOf(this.country);
      var profile = {
        user_name: this.userName.trim(),
        email: this.email.trim(),
        password: this.password.trim(),
        country_id: this.countries[countryIndex],
        user_full_name: this.fullName.trim(),
        birthdate: this.birthDate,
        bio: this.bio,
        fb_link: this.facebookLink,
      };
      console.log(profile);

      axios
        .post(`${process.env.API}/auth/register`, profile, {
          withCredentials: true,
        })
        .then((res) => {
          if (res.data.status != 200) {
            this.$q.notify({
              type: "negative",
              message: `Error: ${res.data.message}`,
            });
          } else {
            this.$q.notify({
              type: "positive",
              message: `Success`,
            });
            this.userName = null;
            this.email = null;
            this.password = null;
            this.country = null;
            this.fullName = null;
            this.birthDate = today;
            this.bio = null;
            this.facebookLink = null;
          }
        })
        .catch((err) => console.log(err));
    },

    onReset() {
      this.name = null;
      this.age = null;
      this.accept = false;
    },

    isUserNameExist(val) {
      // simulating a delay

      return new Promise((resolve, reject) => {
        var data = {
          user_name: this.userName.trim(),
        };
        axios
          .post(`${process.env.API}/auth/register-username`, data, {
            withCredentials: true,
          })
          .then((res) => {
            if (res.data.status !== 200) {
              resolve(res.data.error);
            } else {
              resolve(true);
            }
            console.log(res.data);
          })
          .catch((err) => console.log(err));
      });
    },
    isEmailValidAndNotUsed(val) {
      return new Promise((resolve, reject) => {
        const emailPattern = /^(?=[a-zA-Z0-9@._%+-]{6,254}$)[a-zA-Z0-9._%+-]{1,64}@(?:[a-zA-Z0-9-]{1,63}\.){1,8}[a-zA-Z]{2,63}$/;
        if (!emailPattern.test(val.trim())) {
          return resolve("Please enter valid email");
        }

        var data = {
          email: val.trim(),
        };
        axios
          .post(`${process.env.API}/auth/register-email`, data, {
            withCredentials: true,
          })
          .then((res) => {
            if (res.data.status !== 200) {
              resolve(res.data.error);
            } else {
              resolve(true);
            }
            console.log(res.data);
          })
          .catch((err) => console.log(err));
      });
    },
    filterFn(val, update, abort) {
      update(
        () => {
          if (val === "") {
            this.options = stringOptions;
          } else {
            const needle = val.toLowerCase();
            this.options = stringOptions.filter(
              (v) => v.toLowerCase().indexOf(needle) > -1
            );
          }
        },

        // next function is available in Quasar v1.7.4+;
        // "ref" is the Vue reference to the QSelect
        (ref) => {
          if (val !== "" && ref.options.length > 0 && ref.optionIndex === -1) {
            ref.moveOptionSelection(1, true); // focus the first selectable option and do not update the input-value
            ref.toggleOption(ref.options[ref.optionIndex], true); // toggle the focused option
          }
        }
      );
    },

    setModel(val) {
      this.country = val;
    },
  },
  mounted() {
    axios.get(`${process.env.API}/public/country-list`).then((res) => {
      stringOptions = res.data.data.map((res) => {
        return res.N;
      });
      this.countries = res.data.data.map((res) => {
        return res.I;
      });
      console.log(stringOptions);
    });
  },
};
</script>
