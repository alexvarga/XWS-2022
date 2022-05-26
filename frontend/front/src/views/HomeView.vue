<template>
  <div class="home">
  
  <post-card></post-card>



  </div>
</template>

<script>
import {getToken} from "../token/token.js"
import HelloWorld from "../components/HelloWorld";
import HeaderComponent from "@/components/HeaderComponent.vue";
import PostCard from "../components/PostCard.vue"

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
      
    }
  },
  methods: {
    checkLoggedInUser(){
    if (getToken() != null) {
      this.loggedInUser=true;
      console.log(loggedInUser);
    } 
    },


  },

  mounted() {
    this.checkLoggedInUser();
    
  }
};
</script>
