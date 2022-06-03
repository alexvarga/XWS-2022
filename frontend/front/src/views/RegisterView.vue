<template>
  <v-app class="grey lighten-4">
    <v-snackbar v-model="snackbar" centered timeout="3500">
      <span>{{ snackbarText }}</span>
    </v-snackbar>
    <v-container center>
      <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">
          <v-toolbar flat height="45" width="800px">
            <v-app-bar app height="45" color="grey lighten-3">
              <v-app-bar-nav-icon @click="$router.push('/')">
                <v-icon>mdi-arrow-left</v-icon>
              </v-app-bar-nav-icon>
              <v-row>
                <v-col>
                  <v-toolbar-title> </v-toolbar-title>
                </v-col>
              </v-row>
            </v-app-bar>
          </v-toolbar>
          <v-container>
            <v-form ref="loginForm">
              <v-row>
                <v-col>
                  <v-text-field
                    v-model="form.email"
                    label="Username"
                    width="100px"
                    :rules="rules.usernameRules"
                  ></v-text-field>
                  <v-text-field
                    v-model="form.password"
                    type="password"
                    label="Password"
                    width="100px"
                    :rules="rules.passwordRules"
                  ></v-text-field>
                  <v-text-field
                    v-model="form.firstName"
                    label="First Name"
                    width="100px"
                  ></v-text-field>
                  <v-text-field
                    v-model="form.lastName"
                    label="Last Name"
                    width="100px"
                  ></v-text-field>
                  <v-select
                    v-on:change="setGender"
                    label="Gender"
                    :items="items"
                  ></v-select>
                  <v-text-field
                    v-model="form.age"
                    label="Age"
                    width="100px"
                    :rules="rules.ageRules"
                  ></v-text-field>

                  <v-btn class="success" @click="submit">Register</v-btn>
                </v-col>
              </v-row>
            </v-form>
          </v-container>
        </v-col>
        <v-col width="300px"></v-col>
      </v-row>
    </v-container>
  </v-app>
</template>

<script>
import axios from "axios";
import { setToken } from "../token/token.js";

export default {
  name: "RegisterView",
  data() {
    return {
      items: ["Male", "Female", "Other"],
      form: {
        email: "",
        password: "",
        firstName: "",
        lastName: "",
        age: "",
        gender: 0,
      },
      rules: {
        passwordRules: [(password) => !!password || "Password is required!"],
        usernameRules: [(email) => !!email || "Username is required"],
        ageRules: [
          (age) => /^[0-9\s]+$/.test(age) || "Age must contain numbers only",
        ],
      },
      snackbar: false,
      snackbarText: "",
    };
  },
  methods: {
    submit() {
      let registerCredentials = {
        email: this.form.email,
        password: this.form.password,
      };
      let userCredentials = {
        email: this.form.email,
        password: this.form.password,
        firstName: this.form.firstName,
        lastName: this.form.lastName,
        gender: this.form.gender,
        age: this.form.age,
      };
      axios
        .post("http://localhost:8080/api/auth/register", registerCredentials)
        .then((response) => {
          console.log(response);
        })
        .catch((error) => {
          if (error.response.status === 400) {
            this.snackbarText = "errror 1";
            this.snackbar = true;
          } else if (error.response.status === 500) {
            this.snackbarText = "server error!";
            this.snackbar = true;
          } else if (error.response.status === 405) {
            this.snackbarText = "error !";
            this.snackbar = true;
          }
        });


      axios
        .post("http://localhost:8080/api/user/user", userCredentials)
        .then((response2) => {
          this.$router.go(-1);

        })
        .catch((error2) => {
          if (error2.response2.status === 400) {
            this.snackbarText = "errror 1";
            this.snackbar = true;
          } else if (error2.response2.status === 500) {
            this.snackbarText = "server error!";
            this.snackbar = true;
          } else if (error2.response2.status === 405) {
            this.snackbarText = "error !";
            this.snackbar = true;
          }
        });
    },
    setGender(a) {
      if (a == "Male") {
        this.form.gender = 0;
      } else if (a == "Female") {
        this.form.gender = 1;
      } else {
        this.form.gender = 2;
      }
    },
  },
};
</script>

<style>
</style>