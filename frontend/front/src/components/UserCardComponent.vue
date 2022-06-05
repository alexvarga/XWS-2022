<template> <div>
  <v-card class="ma-4"  :to="'/user/'+this.user.ID" elevation="3">
    <v-card-title>{{this.user.FirstName}} {{user.LastName}}  </v-card-title>
    <v-card-text>{{this.user.Bio}}</v-card-text>
    

    <v-card-text>
      <div>
          <!-- <h4>Bio:</h4> -->
          {{this.user.bio}}
       </div> 

      </v-card-text>
  </v-card>

</div>

</template>

<script>
export default {
 name: "UserCardComponent",
 props: ['user'],
 data(){
   return {


   }
 },
 methods: {


 },
 mounted() {



 }

}
</script>