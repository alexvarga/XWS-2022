<template>
<div>
  <div v-if="this.type=='ers'">
    
    <div v-for="follower in this.followers" :key="follower.ID">
 
<following-card-component :userId="follower.FollowerID" ></following-card-component>
  
</div>

</div>

  <div v-if="this.type=='ing'">
    
    <div v-for="following in this.followings" :key="following.ID">

    <!-- {{followings}} -->
 
<following-card-component :userId="following.FolloweeID" ></following-card-component>
  
</div>

</div>


</div>

</template>

<script>
import FollowingCardComponent from '../components/FollowingCardComponent.vue';

export default {
  name: "FollowingComponent",
  props: ["followers", "followings", "type"],
  components: {
    FollowingCardComponent,
  },
    

  data() {
    return {
      followerId: "",
    };
  },
  methods: {

  },
  mounted() {

  }
};
</script>