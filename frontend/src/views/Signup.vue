<template>
  <v-container>
    <v-app>
      <v-row justify="center" class="mt-8">
        <v-col cols="12" sm="10" md="8" lg="6">
          <v-card ref="form">
            <v-card-text>
              <v-text-field
                ref="name"
                v-model="signupData.name"
                :rules="[() => !!signupData.name || 'This field is required']"
                :error-messages="errorMessages"
                label="Full Name"
                required
              ></v-text-field>
              <v-text-field
                ref="username"
                v-model="signupData.username"
                :rules="[
                  () => !!signupData.username || 'This field is required',
                ]"
                :error-messages="errorMessages"
                label="Username"
                required
              ></v-text-field>
              <v-text-field
                ref="email"
                v-model="signupData.email"
                :rules="emailRules"
                :error-messages="errorMessages"
                label="E-mail"
                required
              ></v-text-field>
              <v-text-field
                ref="password"
                v-model="signupData.password"
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
              Have an account ?
              <v-btn
                depressed
                color="primary"
                text
                @click="$router.push('/login')"
              >
                Login
              </v-btn>
              <v-spacer></v-spacer>
              <v-btn depressed color="primary" @click="submit"> Signup </v-btn>
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
    signupData: {
      name: null,
      password: null,
      username: null,
      email: null,
    },
    errorMessages: "",
    formHasErrors: false,
    showPassword: true,
    rules: {
      required: (value) => !!value || "Required.",
    },
    emailRules: [
      (v) => !!v || "E-mail is required",
      (v) => /.+@.+\..+/.test(v) || "E-mail must be valid",
    ],
  }),

  computed: {
    form() {
      return {
        name: this.signupData.name,
        username: this.signupData.username,
        password: this.signupData.password,
        email: this.signupData.email,
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
        this.signupData.password &&
        this.signupData.username &&
        !this.signupData.name
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
      this.signup();
    },
    signup() {
      axios.post("/signup", this.signupData, axiosConfig).then(() => {
        this.$router.push({ name: "Login" });
      });
    },
  },
};
</script>
