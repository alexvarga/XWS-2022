<template> <div>
  <v-card class="ma-4"  :to="'/user/'+userId" elevation="3">
    <v-card-title>{{this.user.firstName}} {{user.lastName}}  </v-card-title>
    

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
 name: "FollowingCardComponent",
 props: ['userId'],
 data(){
   return {
     user: {
       firstName: "",
       lastName: "",
       bio: ""
     },
     decodedContent: "",

   }
 },
 methods: {
   getUserById(){
      axios
      .get("http://localhost:8080/api/user/user/"+this.userId).then(response=> {
        console.log( response.data)


        this.user.firstName=response.data.FirstName;
        this.user.lastName=response.data.LastName;
        this.user.bio = response.data.Bio;
        console.log("log");

        //return this.user.email;
        
      });
    },

 },
 mounted() {

   this.getUserById();

 }

}
</script>