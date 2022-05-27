<template>
  <div class="home">
  <div v-for="post in allPosts" :key="post.id">

    <post-card :post="post"></post-card>
  </div>



  </div>
</template>

<script>
import {getToken} from "../token/token.js"
import HelloWorld from "../components/HelloWorld";
import HeaderComponent from "@/components/HeaderComponent.vue";
import PostCard from "../components/PostCard.vue"
import axios from 'axios';

export default {
  name: "Home",

  components: {
    HelloWorld,
    HeaderComponent,
    PostCard
  },
  data(){
    return {
      loggedInUser: false,
      allPosts: [],
      
    }
  },
  methods: {
    checkLoggedInUser(){
    if (getToken() != null) {
      this.loggedInUser=true;
      console.log(this.loggedInUser);
    } 
    },
    loadPosts(){
      axios
      .get("http://localhost:8080/api/post/posts").then(response=> {
        this.allPosts=response.data;
        console.log(this.allPosts[0], 'post');
      });
    }

  },

  mounted() {
    this.checkLoggedInUser();
    this.loadPosts();
    
  }
};
</script>
