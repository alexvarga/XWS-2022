<template>
  <div>
    <v-snackbar v-model="snackbar" centered timeout="3500">
      <span>{{ snackbarText }}</span>
    </v-snackbar>
    <v-container center>
      <v-row>
        <v-col width="300px"></v-col>
        <v-col width="600px">
          <v-toolbar flat height="45" width="800px"> </v-toolbar>
          <v-container>
            <v-form ref="loginForm">
              <v-row>
                <v-col>
                  <!-- <v-text-field
                    v-model="user.password"
                    type="password"
                    label="Password"
                    width="100px"
                    :rules="rules.passwordRules"
                  ></v-text-field> -->
                  <v-text-field
                    @change="dataChanged('firstName')"
                    v-model="user.firstName"
                    label="First Name"
                    width="100px"
                  ></v-text-field>
                  <v-text-field
                    @change="dataChanged('lastName')"
                    v-model="user.lastName"
                    label="Last Name"
                    width="100px"
                  ></v-text-field>
                  <v-select
                    @change="dataChanged('gender')"
                    v-on:change="setGender"
                    label="Gender"
                    :items="items"
                  ></v-select>
                  <v-text-field
                    @change="dataChanged('age')"
                    v-model="user.age"
                    label="Age"
                    width="100px"
                    :rules="rules.ageRules"
                  ></v-text-field>
                  <v-text-field
                    @change="dataChanged('phoneNumber')"
                    v-model="user.phoneNumber"
                    label="Phone Number"
                    width="100px"
                    :rules="rules.ageRules"
                  ></v-text-field>
                  <v-textarea
                    @change="dataChanged('bio')"
                    v-model="user.bio"
                    label="Bio"
                    width="100px"
                  ></v-textarea>
                  <v-checkbox
                  @change="dataChanged('privateProfile')"
                    v-model="user.privateProfile"
                    :label="`Private profile`"
                  ></v-checkbox>

                  <v-btn class="success" @click="submit">Update</v-btn>
                </v-col>
              </v-row>
            </v-form>
          </v-container>
        </v-col>
        <v-col width="300px"></v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import Axios from "axios";
export default {
  name: "UpdateProfileComponent",
  props: ["user"],

  data() {
    return {
      items: ["Male", "Female", "Other"],
      data2: {},
      changes: [],
      changed: {},
      rules: {
        passwordRules: [(password) => !!password || "Password is required!"],
        usernameRules: [(email) => !!email || "Username is required"],
        ageRules: [
          (age) => /^[0-9\s]+$/.test(age) || "Must contain numbers only",
        ],
      },
      snackbar: false,
      snackbarText: "",
    };
  },

  methods: {
    submit() {
      console.log(this.user, " user object");
      console.log(this.data2, "data2 object");

      Axios.put("http://localhost:8080/api/user/user", this.data2).then(
        (response) => {
          console.log(response);
          this.data2 = {};

          this.snackbar = true;
          this.snackbarText = "Sucessfully updated";
        }
      );
    },

    displayGender() {},

    makeAnObject() {
      this.data2 = { email: this.user.email };

      for (let i = 0; i < this.changes.length; i++) {
        let helper = "";
        if (this.changes[i] == "firstName") {
          helper = this.user.firstName;
        } else if (this.changes[i] == "lastName") {
          helper = this.user.lastName;
        } else if (this.changes[i] == "gender") {
          helper = this.user.gender;
        } else if (this.changes[i] == "age") {
          helper = this.user.age;
        } else if (this.changes[i] == "bio") {
          helper = this.user.bio;
        } else if (this.changes[i] == "phoneNumber") {
          helper = this.user.phoneNumber;
        }else if (this.changes[i] == "privateProfile") {
          helper = this.user.privateProfile;
        }

        this.data2[this.changes[i]] = helper;
      }
      console.log(this.data2, "created data");
    },

    dataChanged(a) {
      if (this.changes.indexOf(a) === -1) {
        this.changes.push(a);
      }
      this.makeAnObject();
      console.log(this.changes);
    },

    setGender(a) {
      if (a == "Male") {
        this.user.gender = 0;
      } else if (a == "Female") {
        this.user.gender = 1;
      } else {
        this.user.gender = 2;
      }
    },
  },
};
</script>