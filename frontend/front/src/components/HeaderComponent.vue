<template>
  <div>
    <v-app-bar color="indigo accent-4" dense dark>
      <v-app-bar-nav-icon></v-app-bar-nav-icon>

      <v-toolbar-title >
          
          <v-btn x-large :to="'/'" text  >Dislinkt</v-btn>
          </v-toolbar-title>

      <v-spacer></v-spacer>

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

      <v-btn text v-if="isLoggedIn" @click="goToUserHome" >
        <!-- <v-icon>Log Off</v-icon> -->
        {{ this.loggedInUser.Email }}
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <v-menu left bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-btn icon v-bind="attrs" v-on="on">
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>

        <v-list>
          <v-list-item v-for="n in 5" :key="n" @click="() => {}">
            <v-list-item-title>Option {{ n }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
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
            console.log(response.data, "response data");
            console.log(this.loggedInUser, "logged in user");
          });
      }
    },
    goToRegister() {
      
      this.$router.push({ path: "/register" });
    },
    goToLogin() {
      this.$router.push({ path: "/login" });
    },
    goToUserHome(){

      this.$router.push({path: "/myHome"});
    }
  },
  mounted() {
    this.checkLoggedInUser();
  },
  created() {
    this.checkLoggedInUser();
  },
};
</script>