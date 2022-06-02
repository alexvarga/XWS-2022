<template>
  <div>
    <p>hi {{ this.userId }}</p>
    <p>{{ this.user.firstName }}</p>

    <div class="">
      <div v-for="post in userPosts" :key="post.id">
        <post-card :post="post"></post-card>
      </div>

      <div v-for="experienceItem in user.experience" :key="experienceItem.id">
        <experience-card-component :experienceItem="experienceItem"> </experience-card-component>
      </div>


    </div>
  </div>
</template>

<script>
import Axios from "axios";
import PostCard from "../components/PostCard.vue";
import ExperienceCardComponent from "@/components/ExperienceCardComponent.vue";

export default {
  components: { PostCard, ExperienceCardComponent },
  name: "ProfileView",
  data() {
    return {
      user: {
        id: "",
        email: "",
        firstName: "",
        lastName: "",
        gender: "",
        experience: [],
        interests: [],
      },
      userPosts: [],
    };
  },

  methods: {
    getUser() {
      Axios.get("http://localhost:8080/api/user/user/" + this.userId).then(
        (response) => {
          this.user.id = response.data.Id;
          this.user.email = response.data.Email;
          this.user.firstName = response.data.FirstName;
          this.user.lastName = response.data.LastName;
          this.user.gender = response.data.Gender;
          this.user.experience = response.data.Experience;
          this.user.interests = response.data.Interests;
          //console.log(response.data.Experience[1], "exp")
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
  },

  created() {
    this.userId = this.$route.params.userId;
  },
  mounted() {
    this.getUser();
    this.getPostsForThisUser();
  },
};
</script>