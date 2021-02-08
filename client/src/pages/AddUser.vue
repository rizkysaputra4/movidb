<template>
  <div class="q-pa-md q-mx-auto" style="width: 400px">
    <h4>Add New User Data</h4>

    <q-form @submit="onSubmit" @reset="onReset" class="q-gutter-md">
      <q-input
        filled
        v-model="userName"
        label="UserName *"
        hint="Minimum 6 character"
        lazy-rules
        :rules="[(val) => (val && val.length > 0) || 'Please type something']"
      />

      <q-input
        filled
        v-model="email"
        label="Email *"
        lazy-rules
        :rules="[
          (val) => (val !== null && val !== '') || 'Please type your Email',
        ]"
      />

      <q-input
        filled
        type="password"
        v-model="password"
        hint="Minimum 6 character"
        label="Password *"
        lazy-rules
        :rules="[
          (val) => (val !== null && val !== '') || 'Please type your Password',
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
        user_name: this.userName,
        email: this.email,
        password: this.password,
        country_id: this.country,
        user_full_name: this.fullName,
        birthdate: this.birthDate,
        bio: this.bio,
        fb_link: this.facebookLink,
      };
      console.log(profile);
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
  },
};
</script>
