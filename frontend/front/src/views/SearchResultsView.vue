<template>
  <div>
    <header-search-component></header-search-component>

    <v-card class="ma-4">
      <v-card-title class="justify-center">Search Users</v-card-title>
      <v-container>
        <v-row no-gutters no-wrap>
          <v-text-field
            v-model="searchQuery"
            label="Search Users"
          ></v-text-field>

          <v-btn icon @click="search()">
            <v-icon>mdi-magnify</v-icon>
          </v-btn>
        </v-row>
      </v-container>
    </v-card>

<div v-for="user in users" :key="user.ID">
    <user-card-component :user="user"></user-card-component>
</div>

  </div>
</template>

<script>
import HeaderSearchComponent from "@/components/HeaderSearchComponent.vue";
import Axios from "axios";
import UserCardComponent from '@/components/UserCardComponent.vue';

export default {
  name: "SearchResultsView",

  components: { HeaderSearchComponent, UserCardComponent },

  methods: {
    search() {
      Axios.get(
        "http://localhost:8080/api/user/user/search/" + this.searchQuery + " "
      ).then((response) => {
          this.users=response.data;
        console.log(this.users, "response for search data");
      });
    },

    getUsers() {
      Axios.get(
        "http://localhost:8080/api/user/user/search/" + this.query + " "
      ).then((response) => {
          this.users=response.data;

        console.log(response.data, "response data");
      });
    },
  },
  data() {
    return {
      searchQuery: "",
      query: "",
      users: [],
    };
  },
  created() {
    this.query = this.$route.params.query;
    console.log(this.query);
    this.getUsers();
  },
};
</script>

<style scoped>
.v-btn.v-btn--icon.v-btn--round.theme--light.v-size--default {
  top: 15px;
}
</style>
