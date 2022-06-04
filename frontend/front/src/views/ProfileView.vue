<template>
  <div>
          <header-component></header-component>

    <!-- <p>hi {{ this.userId }}</p>
    <p>{{ this.user.firstName }}</p> -->

    <v-card class="ma-4 text-center" elevation="0">
      <v-card-title class="justify-center"
        >{{ this.user.firstName }} {{ this.user.lastName }}</v-card-title
      >
      <v-card-text>Bio: {{ this.user.bio }}</v-card-text>
      <div v-if="!(loggedInUser == this.userId)">
        <v-btn @click="followUser" v-if="!follows && isLoggedIn" class="center">Follow</v-btn>
        <v-btn v-if="follows && isLoggedIn" class="center">Unfollow</v-btn>
      </div>
    </v-card>

    <div class="">
      <div v-for="post in userPosts" :key="post.id">
        <post-card :post="post"></post-card>
      </div>

      <div v-for="experienceItem in user.experience" :key="experienceItem.id">
        <experience-card-component :experienceItem="experienceItem">
        </experience-card-component>
      </div>
    </div>
  </div>
</template>

<script>
import { getToken, getUsername } from "../token/token.js";

import Axios from "axios";
import PostCard from "../components/PostCard.vue";
import HeaderComponent from '../components/HeaderComponent.vue';

import ExperienceCardComponent from "@/components/ExperienceCardComponent.vue";

export default {
  components: { PostCard, ExperienceCardComponent, HeaderComponent },
  name: "ProfileView",
  data() {
    return {
      isLoggedIn: false,
      loggedInUser: "",
      follows: false,
      user: {
        id: "",
        email: "",
        firstName: "",
        lastName: "",
        gender: "",
        bio: "",
        experience: [],
        interests: [],
      },
      userPosts: [],
    };
  },

  methods: {
    checkLoggedInUser() {
      if (getToken() != null) {
        this.isLoggedIn = true;
      }
    },
    isFollower(userProfile) {
      if (this.isLoggedIn) {
        Axios.get(
          "http://localhost:8080/api/user/user/id/" + getUsername()
        ).then((response) => {
          let currentlyLogged = response.data;
          this.loggedInUser = currentlyLogged;
          
          Axios.get(
            "http://localhost:8080/api/follow/followss/" +
              this.userId +
              "/" +
              currentlyLogged
          ).then((response) => {
            this.follows = response.data;
            console.log(this.follows, "follows");
          });
        });
      }
    },
    getUser() {
      Axios.get("http://localhost:8080/api/user/user/" + this.userId).then(
        (response) => {
          this.user.id = response.data.Id;
          this.user.email = response.data.Email;
          this.user.firstName = response.data.FirstName;
          this.user.lastName = response.data.LastName;
          this.user.gender = response.data.Gender;
          this.user.bio = response.data.Bio;
          this.user.experience = response.data.Experience;
          this.user.interests = response.data.Interests;

          this.isFollower();
        }
      );
    },
    getPostsForThisUser() {
      Axios.get("http://localhost:8080/api/post/posts/" + this.userId).then(
        (response) => {
          response.data.forEach((el) => {
            this.userPosts.push(el);
          });
          console.log(this.userPosts);
        }
      );
    },
    followUser() {
      Axios.post(
        "http://localhost:8080/api/follow/follow/" +
          this.userId +
          "/" +
          this.loggedInUser
      ).then((response) => {
        this.follows=true;
        console.log("following response", response.data)
      });
    },
  },

  created() {
    this.userId = this.$route.params.userId;
  },
  mounted() {
    this.checkLoggedInUser();

    this.getUser();
    this.getPostsForThisUser();
  },
};
</script>