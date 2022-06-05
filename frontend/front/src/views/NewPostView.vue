<template>
  <div>
          <header-component></header-component>

    <div v-if="isLoggedIn==true">

    <v-card class="ma-4 text-center">
      <v-text-field
        v-model="postTitle"
        label="Post TItle"
        solo
      ></v-text-field>

      <vue-editor v-model="data.content" />

        <v-btn @click="submit">Submit</v-btn>

    </v-card>
    </div>
    <div v-else>
        Not logged in
    </div>
  </div>
</template>

<script>
import HeaderComponent from '../components/HeaderComponent.vue';
import Axios from "axios";
import { VueEditor } from "vue2-editor";
import { getToken, getId, getUsername } from "../token/token.js";


export default {
  components: { VueEditor, HeaderComponent },
  methods: {
    submit() {

      this.data.userId=this.loggedInUserId;
      this.data.title=this.postTitle;
      this.data.content = btoa(this.data.content);
      Axios.post("http://localhost:8080/api/post/post", this.data).then(
        (response) => {
          this.$router.push({ path: "/myHome" });
        }
      );
    },
    checkLoggedInUser() {
      if (getToken() != null) {
        this.isLoggedIn = true;
        Axios.get("http://localhost:8080/api/user/user/id/"+getUsername()).then((response)=>
        {
          this.loggedInUserId=response.data;
        })
      }
    },
  },

  data() {
    return {
      isLoggedIn: false,
      loggedInUserId: "",
      postTitle: "Post Title",
      data: {
        title: "Post Title",
        userId: "",
        content: "initial content",
      },
    };
  },
  created() {
    this.checkLoggedInUser();
  }
};
</script>