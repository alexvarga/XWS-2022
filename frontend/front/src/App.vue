<template>
  <v-app>
  
  <!-- <header-component></header-component> -->

    <v-main>
      <router-view/>
    </v-main>

  </v-app>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue';
import TestComponent from './components/TestComponent.vue';
import HeaderComponent from './components/HeaderComponent.vue';

export default {
  components: { HelloWorld, TestComponent, HeaderComponent },
  name: 'App',

  data: () => ({
    //
  }),
};
</script>
