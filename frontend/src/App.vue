<template>
  <v-app>
    <v-app-bar app dark>
      <div class="d-flex align-center">
        <router-link to="/">
          <v-icon>mdi-home</v-icon>
        </router-link>
      </div>

      <v-spacer></v-spacer>

      <v-btn color="error" @click="signout" text v-if="isUserAuth">
        <span class="mr-2">Signout</span>
      </v-btn>
      <v-btn @click="$router.push({ name: 'Login' })" text v-else>
        <span class="mr-2">Login / Signup</span>
      </v-btn>
    </v-app-bar>

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import axios from "axios";
import axiosConfig from "./axios.config";
export default {
  name: "App",

  data: () => ({
    isUserAuth: localStorage.getItem("token"),
  }),
  beforeCreate() {
    // const isUserAuth = localStorage.getItem("token");
    // if (isUserAuth) {
    //   this.$router.push({ name: "Home" });
    // } else {
    //   this.$router.push({ name: "Login" });
    // }
  },
  methods: {
    signout() {
      axios.delete("/signout", axiosConfig).then(() => {
        localStorage.removeItem("token");
        this.$router.push({ name: "Login" });
      });
    },
  },
};
</script>

<style>
#app,
html,
body,
.v-application--wrap {
  background: #f4f4f4 !important;
}
a {
  text-decoration: none;
}
</style>
