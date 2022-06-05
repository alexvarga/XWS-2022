<template>
  <div>
    <v-app-bar color="indigo accent-4" dense dark>
      <v-app-bar-nav-icon></v-app-bar-nav-icon>
<div>

      <v-toolbar-title>
        <v-btn x-large :to="'/'" text>Dislinkt</v-btn>
      </v-toolbar-title>
</div>

      <v-spacer></v-spacer>
      <!-- <div>
        <v-container fluid >
          <v-row  no-gutters justify="center" align-content="center" align="end">
            <v-col  cols="12" sm="8" md="8" align-self="end" >

              <v-text-field v-model="searchQuery"     label="Search Users"></v-text-field>
            </v-col>


            <v-col cols="12" sm="4" md="4">
              <v-btn :to="'/search/results/'+searchQuery" icon>
                <v-icon>mdi-magnify</v-icon>
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
      </div> -->
      <div>


      <v-btn text v-if="isLoggedIn" @click="logOff">
        <!-- <v-icon>Log Off</v-icon> -->
        Log Out
      </v-btn>
      <v-btn text v-if="!isLoggedIn" @click="goToRegister">
        <!-- <v-icon>Log Off</v-icon> -->
        Register
      </v-btn>
      <v-btn text v-if="!isLoggedIn" @click="goToLogin">
        <!-- <v-icon>Log Off</v-icon> -->
        Log In
      </v-btn>

      <v-btn text v-if="isLoggedIn" @click="goToUserHome">
        <!-- <v-icon>Log Off</v-icon> -->
        {{ this.loggedInUser.Email }}
      </v-btn>
      </div>


    </v-app-bar>
  </div>
</template>

<script>
import { getToken, getId, getUsername } from "../token/token.js";

export default {
  data() {
    return {
      isLoggedIn: false,
      loggedInUser: "",
      searchQuery: "",
    };
  },

  methods: {
    logOff() {
      localStorage.clear();
      this.$router.push({ path: "/" });
      window.location.reload();
    },
    checkLoggedInUser() {
      if (getToken() != null) {
        this.isLoggedIn = true;
        axios
          .get("http://localhost:8080/api/user/user/email/" + getUsername())
          .then((response) => {
            this.loggedInUser = response.data;
          });
      }
    },
    goToRegister() {
      this.$router.push({ path: "/register" });
    },
    goToLogin() {
      this.$router.push({ path: "/login" });
    },
    goToUserHome() {
      this.$router.push({ path: "/myHome" });
    },
    // search(){
    //   console.log(this.searchQuery);
    // },
    // getUsers(){

    // },
    
  },
  mounted() {
    this.checkLoggedInUser();
  },
  created() {
    this.checkLoggedInUser();
  },
};
</script>

<style>
.v-input__slot {
	top: 10px;
}
</style>