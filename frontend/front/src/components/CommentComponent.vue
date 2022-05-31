<template>
  <v-card class="ma-4">
    <v-card-actions> Comment </v-card-actions>
    <v-card-text>{{comment.Text}} </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "comment-component",
  props: ['comment'],
  data() {},
  methods: {},
};
</script>