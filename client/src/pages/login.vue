<template>
  <q-layout id="particles-js">
    <q-page-container>
      <q-page class="flex flex-center">
        <div
          id="particles-js"
          :class="$q.dark.isActive ? 'dark_gradient' : 'normal_gradient'"
        >
          <Particles id="tsparticles" :options="particleOptions" />
        </div>

        <q-btn
          color="white"
          class="absolute-top-right"
          flat
          round
          @click="$q.dark.toggle()"
          :icon="$q.dark.isActive ? 'nights_stay' : 'wb_sunny'"
        />

        <q-card
          class="login-form"
          v-bind:style="
            $q.platform.is.mobile ? { width: '80%' } : { width: '30%' }
          "
        >
          <q-card-section>
            <q-avatar
              class="absolute"
              style="top: 0; right: 25px; transform: translateY(-50%)"
            >
              <img src="https://cdn.quasar.dev/img/boy-avatar.png" />
            </q-avatar>

            <div class="row no-wrap items-center">
              <div class="col text-h6 ellipsis">MoviWiki Admin Section</div>
            </div>
          </q-card-section>
          <q-card-section>
            <q-form class="q-gutter-md" @submit="onSubmit">
              <q-input
                filled
                v-model="username"
                label="Username"
                lazy-rules
                :rules="[
                  (val) => (val && val.length > 0) || 'UserName Required',
                ]"
              />

              <q-input
                type="password"
                filled
                v-model="password"
                label="Password"
                lazy-rules
                :rules="[
                  (val) => (val && val.length > 0) || 'Password Required',
                ]"
              />

              <div>
                <q-btn
                  label="Login"
                  :loading="loading"
                  type="submit"
                  color="primary"
                  @click="loginNotify"
                >
                  <template v-slot:loading>
                    <q-spinner-facebook />
                  </template>
                </q-btn>
              </div>
            </q-form>
          </q-card-section>
        </q-card>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script type="text/javascript"></script>
<script>
import Vue from "vue";
import Particles from "particles.vue";
import axios from "axios";
import { mapMutations } from "vuex";

Vue.use(Particles);

export default {
  name: "Login",
  data() {
    return {
      username: null,
      password: null,
      error: null,
      loading: false,
      particleOptions: {
        background: {
          color: {
            value: "black",
          },
        },
        fpsLimit: 60,
        interactivity: {
          detectsOn: "canvas",
          events: {
            onClick: {
              enable: true,
              mode: "push",
            },
            onHover: {
              enable: true,
              mode: "grab",
            },
            resize: true,
          },
          modes: {
            bubble: {
              distance: 400,
              duration: 2,
              opacity: 0.8,
              size: 40,
            },
            grab: {
              distance: 150,
              links: {
                blink: true,
                consent: false,
                opacity: 0.5,
              },
            },
            push: {
              quantity: 4,
            },
            repulse: {
              distance: 200,
              duration: 0.4,
            },
          },
        },
        particles: {
          color: {
            value: "#ffffff",
          },
          links: {
            color: "#ffffff",
            distance: 150,
            enable: true,
            opacity: 0.5,
            width: 1,
          },
          collisions: {
            enable: false,
          },
          move: {
            direction: "none",
            enable: true,
            outMode: "bounce",
            random: true,
            speed: 3,
            straight: false,
          },
          number: {
            density: {
              enable: true,
              value_area: 1500,
            },
            value: 80,
          },
          opacity: {
            value: 0.5,
          },
          shape: {
            type: "circle",
          },
          size: {
            random: true,
            value: 5,
          },
        },
        detectRetina: true,
      },
    };
  },
  methods: {
    onSubmit(e) {},
    loginNotify() {
      this.loading = true;
      if (!this.username) {
        this.$q.notify({
          color: "red-5",
          textColor: "white",
          icon: "warning",
          message: "UserName Required",
        });
        this.loading = false;
        return;
      }
      if (!this.password) {
        this.$q.notify({
          color: "red-5",
          textColor: "white",
          icon: "warning",
          message: "Password Required",
        });
        this.loading = false;
        return;
      }
      const data = {
        user_name: this.username,
        password: this.password,
      };

      axios
        .post(`${process.env.API}/auth/login-password`, data, {
          withCredentials: true,
        })
        .then((res) => {
          if (res.data.status !== 200) {
            this.$q.notify({
              color: "red-5",
              textColor: "white",
              icon: "warning",
              message: res.data.message,
            });
            this.loading = false;
            return;
          } else if (res.data.data.role <= 11) {
            this.$q.notify({
              color: "green",
              textColor: "white",
              message: "Logged In",
            });
            this.loading = false;
            this.$router.push("/dashboard");
          } else if (res.data.data.role > 11) {
            this.$q.notify({
              color: "red-5",
              textColor: "white",
              icon: "warning",
              message: "For Admin Only",
            });
            this.loading = false;
          }

          let data = {
            userID: res.data.data.userID,
            userName: res.data.data.user_name,
            fullName: res.data.data.user_full_name,
            role: res.data.data.role,
            email: res.data.data.email,
          };

          this.updateDataAuth(data);
          console.log(data);
        })
        .catch((err) => {
          console.log(err);
          this.loading = false;
        });
    },
    ...mapMutations("authStats", ["updateDataAuth"]),
  },
  meta() {
    return {
      title: "MoviWiki Admin",
    };
  },
  mounted() {},
};
</script>

<style>
#particles-js,
#tsparticles {
  position: absolute;
  width: 100%;
  height: 100%;
  background-repeat: no-repeat;
  background-size: cover;
  background-position: 50% 50%;
}
.normal_gradient {
  background: linear-gradient(145deg, rgb(74, 94, 137) 15%, #b61924 70%);
}
.dark_gradient {
  background: linear-gradient(145deg, rgb(11, 26, 61) 15%, #4c1014 70%);
}
.login-form {
  position: absolute;
}
</style>
