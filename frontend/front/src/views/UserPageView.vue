<template>
  <div>
    <v-tabs fixed-tabs>
      <v-tab href="#myPosts">My Posts</v-tab>
      <v-tab-item value="myPosts">

        <v-card class="ma-4 text-center"  elevation="0"> 
        <v-btn class="center"  > New Post</v-btn>
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
      <v-tab-item value="following"> </v-tab-item>

      <v-tab href="#followers">Followers</v-tab>
      <v-tab-item value="followers"> </v-tab-item>

      <v-tab href="#exp">Experience</v-tab>
      <v-tab-item value="exp"> </v-tab-item>

      <v-tab href="#interests">Interests</v-tab>
      <v-tab-item value="interests"> </v-tab-item>

      <v-tab href="#profileData">Update Profile</v-tab>
      <v-tab-item value="profileData">
        <update-profile-componenet :user="this.user" />
      </v-tab-item>
    </v-tabs>
  </div>
</template>

<script>
import { getToken, getId, getUsername } from "../token/token.js";
import axios from "axios";
import PostCard from "../components/PostCard.vue";
import ExperienceCardComponent from "@/components/ExperienceCardComponent.vue";
import UpdateProfileComponenet from "../components/UpdateProfileComponent.vue";

export default {
  components: { PostCard, ExperienceCardComponent, UpdateProfileComponenet },
  name: "UserPageView",
  data() {
    return {
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
    getUser() {
      axios
        .get("http://localhost:8080/api/user/user/email/" + getUsername())
        .then((response) => {
          console.log(response.data, "response data");
          this.user.id = response.data.ID;
          this.user.email = response.data.Email;
          this.user.firstName = response.data.FirstName;
          this.user.lastName = response.data.LastName;
          this.user.gender = response.data.Gender;
          this.user.age = response.data.Age;
          this.user.phoneNumber=response.data.PhoneNumber;
          this.user.bio = response.data.Bio;
          this.user.privateProfile=response.data.PrivateProfile;

          this.user.experience = response.data.Experience;
          this.user.interests = response.data.Interests;
          //console.log(response.data.Experience[1], "exp")
          // console.log(this.user, "user");
          this.getPostsForThisUser();
        });
    },
    getPostsForThisUser() {
      axios
        .get("http://localhost:8080/api/post/posts/" + this.user.id)
        .then((response) => {
          response.data.forEach((el) => {
            this.userPosts.push(el);
          });
          console.log(this.userPosts);
        });
    },
  },
  created() {
    console.log("test created");
  },

  mounted() {
    this.getUser();

    //this.getPostsForThisUser();
  },
};
</script>