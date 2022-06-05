<template>
  <div>
    <new-experience-component
      :userEmail="this.user.email"
    ></new-experience-component>

    <div
      v-for="experienceItem in this.user.experience"
      :key="experienceItem.id"
    >
      <v-card class="ma-2 justify-center">
        <experience-card-component :experienceItem="experienceItem">
        </experience-card-component>
        <v-card-text class="text-right">
          <v-btn>Edit</v-btn>
        </v-card-text>
      </v-card>
    </div>

    <v-content>
      <v-menu
        offset-y
        :close-on-click="false"
        :close-on-content-click="false"
        v-model="menuOpen"
      >
        <template v-slot:activator="{ on }">
          <v-btn v-on="on" class="ma-4">BUTTON</v-btn>
        </template>
        <v-card>
          <v-card-text>
            <v-text-field label="test" v-model="firstname"></v-text-field>
            <v-text-field label="test" v-model="lastname"></v-text-field>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="grey darken-2" @click="menuOpen = false" dark
              >CANCEL</v-btn
            >
            <v-btn color="primary" @click="save">OK</v-btn>
          </v-card-actions>
        </v-card>
      </v-menu>
    </v-content>
  </div>
</template>

<script>
import NewExperienceComponent from "../components/NewExperienceComponent.vue";
import ExperienceCardComponent from "../components/ExperienceCardComponent.vue";

export default {
  components: { NewExperienceComponent, ExperienceCardComponent },
  name: "ExperienceComponent",
  props: ["user"],
  data() {
    return {
      menuOpen: false,
      firstname: "",
      lastname: "",
    };
  },

  methods: {
        save() {
    alert(`Your name is ${this.firstname} ${this.lastname}!`);
    this.menuOpen = false;
  },
  },
  created() {
    console.log(
      this.user,
      "this user object from experience componenet AAAAAAAAAAAAAAAA"
    );
  },

};
</script>