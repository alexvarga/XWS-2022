<template>
  <v-card class="ma-4">
    <v-card-subtitle>
      <v-btn :to="'/user/'+this.comment.UserID" :ripple="false" class="pa-0 ma-0" plain text
        >Commenter: {{ this.userFirstName }} {{ this.userLastName }}
      </v-btn></v-card-subtitle>

    <v-card-text>{{ comment.Text }} </v-card-text>
    <!-- <v-card-text>{{ comment.UserID }}</v-card-text> -->
    <v-card-text></v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "comment-component",
  props: ["comment", "usersCommented"],
  data() {
    return {
      userFirstName: "",
      userLastName: "",
    };
  },
  methods: {
    getUserName() {

      axios
        .get("http://localhost:8080/api/user/user/" + this.comment.UserID)
        .then((response) => {
          this.userFirstName = response.data.FirstName;
          this.userLastName = response.data.LastName;

        });
    },
  },
  mounted() {
    this.getUserName();
  },
};
</script>