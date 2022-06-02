<template>
  <v-card class="ma-4">
    <v-card-title>{{ experienceItem.Title }} </v-card-title>
    <v-card-subtitle>Started: {{experienceItem.DateStarted}} Ended: {{experienceItem.DateEnded}}</v-card-subtitle>

    <v-card-text>{{ experienceItem.Info }} </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "experience-card-component",
  props: ["experienceItem"],
  data() {
    return {};
  },
  methods: {},
};
</script>