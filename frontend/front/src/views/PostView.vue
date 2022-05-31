<template>
  <div class="post">
    <div>
      <v-card outlined class="ma-4">
        <v-card-title>Post title</v-card-title>
        <v-card-subtitle>
          Posted: {{ this.post.published }}
          <br />
          <v-btn :ripple="false" class="pa-0 ma-0" plain text
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
        <v-card-text class="ma-4">
          <div>

            <!-- <p style="white-space: pre-line">{{ comment }}</p> -->
            <v-textarea class="mr-8" solo auto-grow
              v-model="comment"
                label="Your comment"
            ></v-textarea>
          
          <v-btn>Submit</v-btn>
          </div>
        </v-card-text>
        <div v-for="comment in this.post.comments" :key="comment.Id">
        <comment-component :comment="comment" ></comment-component>
        </div>
      </v-card>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import moment from "moment";
import CommentComponent from "@/components/CommentComponent.vue";

export default {
  components: { CommentComponent },
  name: "PostView",
  data() {
    return {
        comment: "",
      post: {
        decodedContent: "",
        userId: "",
        userFirstName: "",
        userLastName: "",
        published: "",
        likes: 0,
        dislikes: 0,
        comments: [],
      },
    };
  },
  methods: {
    getPost() {
      axios
        .get("http://localhost:8080/api/post/post/" + this.postId)
        .then((response) => {
          var decodedContent = atob(response.data.Content);
          this.post.decodedContent = decodedContent;
          this.post.userId = response.data.UserID;
          var time = moment(response.data.published).format(
            "dddd DD-MMMM-YYYY HH:mm"
          );
          this.post.published = time;
          console.log(time);
          this.post.likes = response.data.Likes;
          this.post.dislikes = response.data.Dislikes;
          this.post.comments = response.data.Comments;

          axios
            .get("http://localhost:8080/api/user/user/" + this.post.userId)
            .then((response) => {
              console.log(response.data, "response data duh");
              this.post.userFirstName = response.data.FirstName;
              this.post.userLastName = response.data.LastName;
            });
        });
    },
  },
  created() {
    this.postId = this.$route.params.postId;
  },
  mounted() {
    this.getPost();
  },
};
</script>