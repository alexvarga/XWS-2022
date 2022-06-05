<template>
  <div>
    <header-component></header-component>
    <div v-if="isLogged">
      <v-tabs fixed-tabs>
        <v-tab href="#myPosts">My Posts</v-tab>
        <v-tab-item value="myPosts">
          <v-card class="ma-4 text-center" elevation="0">
            <v-btn :test="test" :to="'/new/post'" class="center">
              New Post</v-btn
            >
          </v-card>

          <div class="">
            <div v-for="post in userPosts" :key="post.id">
              <post-card :post="post"></post-card>
            </div>

            <div
              v-for="experienceItem in user.experience"
              :key="experienceItem.id"
            >
              <experience-card-component :experienceItem="experienceItem">
              </experience-card-component>
            </div>
          </div>
        </v-tab-item>

        <v-tab href="#following">Following</v-tab>
        <v-tab-item value="following">
          <div>
            <following-component
              :type="'ing'"
              :followings="this.followings"
            ></following-component>
          </div>
        </v-tab-item>

        <v-tab href="#followers">Followers</v-tab>
        <v-tab-item value="followers">
          <div>
            <following-component
              :type="'ers'"
              :followers="this.followers"
            ></following-component></div
        ></v-tab-item>

        <v-tab href="#exp">Experience</v-tab>
        <v-tab-item value="exp">
          <experience-component
            :user="this.user" 
          ></experience-component>
        </v-tab-item>

        <v-tab href="#interests">Interests</v-tab>
        <v-tab-item value="interests"> </v-tab-item>

        <v-tab href="#profileData">Update Profile</v-tab>
        <v-tab-item value="profileData">
          <update-profile-componenet :user="this.user" />
        </v-tab-item>
      </v-tabs>
    </div>
  </div>
</template>

<script>
import { getToken, getId, getUsername } from "../token/token.js";
import axios from "axios";
import PostCard from "../components/PostCard.vue";
import ExperienceCardComponent from "@/components/ExperienceCardComponent.vue";
import UpdateProfileComponenet from "../components/UpdateProfileComponent.vue";
import HeaderComponent from "../components/HeaderComponent.vue";
import FollowingComponent from "../components/FollowingComponent.vue";
import NewExperienceComponent from "../components/NewExperienceComponent.vue";
import ExperienceComponent from "../components/ExperienceComponent.vue";

export default {
  components: {
    PostCard,
    ExperienceCardComponent,
    UpdateProfileComponenet,
    FollowingComponent,
    HeaderComponent,
    ExperienceComponent,
  },
  name: "UserPageView",
  data() {
    return {
      isLogged: false,
      followers: [],
      followings: [],
      test: "testtesttest",
      user: {
        id: "",
        email: "",
        firstName: "",
        lastName: "",
        age: "",
        bio: "",
        phoneNumber: "",
        gender: "",
        privateProfile: "",
        experience: [],
        interests: [],
      },
      userPosts: [],
    };
  },

  methods: {
    isLoggedIn() {
      if (getToken() != null) {
        this.isLogged = true;
      }
    },
    getUser() {
      axios
        .get("http://localhost:8080/api/user/user/email/" + getUsername())
        .then((response) => {
          this.user.id = response.data.ID;
          this.user.email = response.data.Email;
          this.user.firstName = response.data.FirstName;
          this.user.lastName = response.data.LastName;
          this.user.gender = response.data.Gender;
          this.user.age = response.data.Age;
          this.user.phoneNumber = response.data.PhoneNumber;
          this.user.bio = response.data.Bio;
          this.user.privateProfile = response.data.PrivateProfile;

          this.user.experience = response.data.Experience;
          this.user.interests = response.data.Interests;
          //console.log(response.data.Experience[1], "exp")
          // console.log(this.user, "user");

          this.getPostsForThisUser();
          this.getFollowersForThisUser();
          this.getFollowingsForThisUser();
        });
    },
    getPostsForThisUser() {
      axios
        .get("http://localhost:8080/api/post/posts/" + this.user.id)
        .then((response) => {
          response.data.forEach((el) => {
            this.userPosts.push(el);
          });
        });
    },
    getFollowersForThisUser() {
      axios
        .get("http://localhost:8080/api/follow/followers/" + this.user.id)
        .then((response) => {
          this.followers = response.data;
        });
    },
    getFollowingsForThisUser() {
      axios
        .get("http://localhost:8080/api/follow/following/" + this.user.id)
        .then((response) => {
          this.followings = response.data;
        });
    },
  },

  created() {
    this.isLoggedIn();
    this.getUser();

  },
};
</script>