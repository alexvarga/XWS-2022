<template>
  <div>
    <new-experience-component
      :userEmail="this.user.email"
    ></new-experience-component>

    <div
      v-for="experienceItem in this.user.experience"
      :key="experienceItem.id"
    >
      <v-card
        @click="shownId = experienceItem.ID"
        :ripple="false"
        class="ma-2 justify-center"
      >
        <experience-card-component :experienceItem="experienceItem">
        </experience-card-component>

        <div>
          <div v-if="shownId == experienceItem.ID">
            <edit-experience-component
              @cancelEdit="clickCancelEdit"
              :experienceItem="experienceItem"
              :userEmail="uEmail"
            ></edit-experience-component>
            <v-card-text class="text-right"> </v-card-text>
          </div>
        </div>
      </v-card>
    </div>
  </div>
</template>

<script>
import NewExperienceComponent from "../components/NewExperienceComponent.vue";
import ExperienceCardComponent from "../components/ExperienceCardComponent.vue";
import EditExperienceComponent from "../components/EditExperienceComponent.vue";

export default {
  components: {
    NewExperienceComponent,
    ExperienceCardComponent,
    EditExperienceComponent,
  },
  name: "ExperienceComponent",
  props: ["user"],
  data() {
    return {
      uEmail: this.user.email,
      shownId: "",
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
    changeId() {},
    clickCancelEdit(value) {

      this.shownId = "";

    },
  },
};
</script>