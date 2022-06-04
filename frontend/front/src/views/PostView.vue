<template>
  <div class="post">
          <header-component></header-component>

    <v-snackbar
      transition="fade-transition"
      v-model="snackbar"
      centered
      timeout="1500"
    >
      <div>{{ snackbarText }}</div>
    </v-snackbar>
    <div>
      <v-card outlined class="ma-4">
        <v-card-title>{{this.post.title}}</v-card-title>
        <v-card-subtitle>
          Posted: {{ this.post.published }}
          <br />
          <v-btn
            :to="'/user/' + post.userId"
            :ripple="false"
            class="pa-0 ma-0"
            plain
            text
            >Author: {{ this.post.userFirstName }}
            {{ this.post.userLastName }}</v-btn
          >
        </v-card-subtitle>
        <v-card-text>
          <div v-html="this.post.decodedContent"></div>
        </v-card-text>
        <v-card-text>
          <v-btn icon color="blue">
            {{ this.post.likes }}
            <v-icon> mdi-thumb-up </v-icon>
          </v-btn>
          <v-btn icon out color="pink">
            {{ this.post.dislikes }}

            <v-icon>mdi-thumb-down</v-icon>
          </v-btn>
        </v-card-text>
      </v-card>
      <v-card class="ma-4">
        <v-card-title>Comments</v-card-title>
        <div>
          <v-card-text class="ma-4">
            <div v-if="loggedInUser">
              <!-- <p style="white-space: pre-line">{{ comment }}</p> -->
              <v-textarea
                id="ta"
                class="mr-8"
                solo
                auto-grow
                v-model="commentHere"
                label="Your comment"
              ></v-textarea>

              <v-btn @click="postAComment">Submit</v-btn>
            </div>
          </v-card-text>
        </div>
        <div v-for="comment in this.post.comments" :key="comment.Id">
          <comment-component :comment="comment"></comment-component>
        </div>
      </v-card>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import moment from "moment";
import CommentComponent from "@/components/CommentComponent.vue";
import HeaderComponent from '../components/HeaderComponent.vue';
import { getToken, getUsername } from "../token/token.js";

export default {
  components: { CommentComponent, HeaderComponent },
  name: "PostView",
  data() {
    return {
      snackbarText: "",
      snackbar: false,
      loggedInUser: false,
      loggedInUserId: "",
      commentHere: "",
      post: {
        id: "",
        title: "",
        decodedContent: "",
        userId: "",
        userFirstName: "",
        userLastName: "",
        published: "",
        likes: 0,
        dislikes: 0,
        comments: [],
      },
      // usersWhoCommented: [],
      // commentWithNames: {
      //   id: "",
      //   text: "",
      //   userFirstName: "",
      //   userLastName: "",
      //   userId: "",
      //   timePublished: "",
      // },
      // commentsWithNames: {
      //   commentWithNames: [],
      // },
    };
  },
  methods: {
    async getPost() {
      axios
        .get("http://localhost:8080/api/post/post/" + this.postId)
        .then((response) => {
          var decodedContent = atob(response.data.Content);
          this.post.decodedContent = decodedContent;
          this.post.userId = response.data.UserID;
          this.post.title=response.data.Title;
          var time = moment(response.data.Published).format(
            "dddd DD-MMMM-YYYY HH:mm"
          );

          this.post.published = time;
          this.post.likes = response.data.Likes;
          this.post.dislikes = response.data.Dislikes;
          this.post.comments = response.data.Comments;


          axios
            .get("http://localhost:8080/api/user/user/" + this.post.userId)
            .then((response3) => {
              console.log(response3.data, "response data duh");
              this.post.userFirstName = response3.data.FirstName;
              this.post.userLastName = response3.data.LastName;
            });
        });
    },
    checkLoggedInUser() {
      if (getToken() != null) {
        this.loggedInUser = true;
        axios
          .get("http://localhost:8080/api/user/user/id/" + getUsername())
          .then((response) => {
            this.loggedInUserId = response.data;
            console.log(this.loggedInUserId, "logged in user id");
          });
      }
    },
    postAComment() {
      if (this.loggedInUser == true) {
        axios.post(
          "http://localhost:8080/api/post/post/comment/" + this.postId,
          {
            userId: this.loggedInUserId,
            text: this.commentHere,
          }
        );

        this.commentHere = "";
        this.snackbarText = "Your comment has been submitted.";
        this.snackbar = true;
      }
    },
    // async getCommentNames() {
    //   // try {
    //   //const responses = [];
    //   for (let i = 0; i < this.post.comments.length; i++) {
    //     await axios
    //       .get(
    //         "http://localhost:8080/api/user/user/" +
    //           this.post.comments[i].UserID
    //       )
    //       .then((response) => {
    //         this.usersWhoCommented.push(response.data);
    //       });
    //   }

    //   // } catch {
    //   //   console.log("AAAAAAAAAAAAAAAAAAAAAAAAAAA");
    //   // }
    // },
  },
  created() {
    this.postId = this.$route.params.postId;
  },
  mounted() {
    this.getPost();
    this.checkLoggedInUser();
  },
};
</script>

<style scoped>
::v-deep .v-snack__wrapper {
  min-width: 0px;
}
</style>