<template>
  <div class="home">
    <div v-for="post in allPosts" :key="post.id">
      <post-card :post="post"></post-card>
    </div>
  </div>
</template>

<script>
import { getToken, getId, getUsername } from "../token/token.js";
import HelloWorld from "../components/HelloWorld";
import HeaderComponent from "@/components/HeaderComponent.vue";
import PostCard from "../components/PostCard.vue";
import axios from "axios";

export default {
  name: "Home",

  components: {
    HelloWorld,
    HeaderComponent,
    PostCard,
  },
  data() {
    return {
      loggedInUser: false,
      loggedInUserId: "",
      allPosts: [],
      allFollows: [],
    };
  },
  methods: {
    checkLoggedInUser() {
      if (getToken() != null) {
        this.loggedInUser = true;
      }
    },
    loadPostsForLoggedInUser() {
      if (this.loggedInUser) {
        axios
          .get("http://localhost:8080/api/user/user/id/" + getUsername())
          .then((response) => {
            this.loggedInUserId = response.data;
            console.log(this.loggedInUserId, "logged in user id");
            axios
              .get(
                "http://localhost:8080/api/follow/following/" +
                  this.loggedInUserId
              )
              .then((response2) => {
                console.log(this.loggedInUserId);
                this.allFollows = response2.data;

                this.allFollows.forEach((element) => {
                  axios
                    .get(
                      "http://localhost:8080/api/post/posts/" +
                        element.FolloweeID
                    )
                    .then((response3) => {
                      response3.data.forEach((el) => {
                        this.allPosts.push(el);
                      });
                    });
                });
              });
          });
      }
    },

    loadPosts() {
      if (!this.loggedInUser) {
        axios.get("http://localhost:8080/api/post/posts").then((response) => {
          this.allPosts = response.data;
        });
      } else {
        console.log("hi");
      }
    },
  },
  created() {
    this.checkLoggedInUser();

    this.loadPostsForLoggedInUser();
  },

  mounted() {
    this.loadPosts();
  },
};
</script>
