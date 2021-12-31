<template>
  <v-container>
    <v-app>
      <v-row justify="center" class="mt-8">
        <v-col cols="12" sm="10" md="8" lg="6">
          <v-card ref="form">
            <v-card-text>
              <v-text-field
                ref="username"
                v-model="loginData.username"
                :rules="[
                  () => !!loginData.username || 'This field is required',
                ]"
                :error-messages="errorMessages"
                label="Username"
                required
              ></v-text-field>
              <v-text-field
                ref="password"
                v-model="loginData.password"
                :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                :rules="[rules.required]"
                :type="showPassword ? 'text' : 'password'"
                counter
                @click:append="showPassword = !showPassword"
                label="Password"
                required
              ></v-text-field>
            </v-card-text>
            <v-divider class="mt-12"></v-divider>
            <v-card-actions>
              Don't have an account ?
              <v-btn
                depressed
                color="primary"
                text
                @click="$router.push('/signup')"
              >
                Signup
              </v-btn>
              <v-spacer></v-spacer>
              <v-btn depressed color="primary" @click="submit"> Login </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-app>
  </v-container>
</template>
<script>
import axios from "axios";
import axiosConfig from "../axios.config";
export default {
  data: () => ({
    loginData: {
      password: null,
      username: null,
    },
    errorMessages: "",
    formHasErrors: false,
    showPassword: true,
    rules: {
      required: (value) => !!value || "Required.",
    },
  }),

  computed: {
    form() {
      return {
        username: this.loginData.username,
        password: this.loginData.password,
      };
    },
  },

  watch: {
    name() {
      this.errorMessages = "";
    },
  },

  methods: {
    addressCheck() {
      this.errorMessages =
        this.loginData.password &&
        this.loginData.username &&
        !this.loginData.name
          ? `Hey! I'm required`
          : "";

      return true;
    },
    resetForm() {
      this.errorMessages = [];
      this.formHasErrors = false;

      Object.keys(this.form).forEach((f) => {
        this.$refs[f].reset();
      });
    },
    submit() {
      this.formHasErrors = false;

      Object.keys(this.form).forEach((f) => {
        if (!this.form[f]) this.formHasErrors = true;

        this.$refs[f].validate(true);
      });
      this.login();
    },
    login() {
      axios.post("/login", this.loginData, axiosConfig).then((response) => {
        localStorage.setItem("token", response.data.data.token);
        this.$router.push({ name: "Home" });
      });
    },
  },
};
</script>
