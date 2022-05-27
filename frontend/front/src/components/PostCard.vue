<template>
  <v-card link elevation="1">
    <v-card-title>Post title </v-card-title>
    <v-card-subtitle> Poster user name</v-card-subtitle>
    
    <v-card-text>
      {{getUserById(post.UserID)}} <h3>{{user.Firstname}} {{user.LastName}} {{user.email}}</h3>    <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto,
      temporibus molestias! Laudantium incidunt omnis fugiat, alias nostrum
      perferendis possimus expedita error vitae id exercitationem ducimus quis
      praesentium in a aperiam?</p>
      <p>{{(post.Content)}}</p> </v-card-text
    >
  </v-card>

</template>

<script>


export default {
 props: ['post'],
 data(){
   return {
     user: {
       email: "",
       firstName: "",
       lastName: ""
     }

   }
 },
 methods: {
   getUserById(id){
      axios
      .get("http://localhost:8080/api/user/user/"+id).then(response=> {
        this.user.email=response.data.Email;
        this.user.firstName=response.data.FirstName;
        this.user.lastName=response.data.LastName;
        console.log("log");
        console.log(this.user.email)
        //return this.user.email;
        
      });
    }
 },

}
</script>