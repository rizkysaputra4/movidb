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

      <q-input filled v-model="country" label="Country" />

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
    };
  },

  methods: {
    onSubmit() {
      var profile = {
        user_name: this.userName.trim(),
        email: this.email,
        password: this.password,
        country_id: this.country,
        user_full_name: this.fullName,
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
          }
        })
        .catch((err) => console.log(err));
      //   if (this.accept !== true) {
      //     this.$q.notify({
      //       color: "red-5",
      //       textColor: "white",
      //       icon: "warning",
      //       message: "You need to accept the license and terms first",
      //     });
      //   } else {
      //     this.$q.notify({
      //       color: "green-4",
      //       textColor: "white",
      //       icon: "cloud_done",
      //       message: "Submitted",
      //     });
      //   }
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
  },
};
</script>
