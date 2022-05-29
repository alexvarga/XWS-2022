<template>
  <v-card class="ma-4" link elevation="3">
    <v-card-title>Post title </v-card-title>
    <v-card-subtitle>Autor: {{this.user.firstName}} {{user.lastName}} {{user.email}} {{post.UserID}}</v-card-subtitle>
    
    <v-card-text>
      <div ></div>
      <!-- <h3>{{user.firstName}} {{user.lastName}} {{user.email}}</h3>   -->
      <!-- <p>{{(post.Content)}}</p>  -->
      <div v-html="decodedContent"></div>
      
      </v-card-text
    >
  </v-card>

</template>

<script>


export default {
 name: "post-card",
 props: ['post'],
 data(){
   return {
     user: {
       email: "",
       firstName: "",
       lastName: ""
     },
     decodedContent: "",

   }
 },
 methods: {
   getUserById(id){
      axios
      .get("http://localhost:8080/api/user/user/"+id).then(response=> {
        console.log(id, response.data)

        this.user.email=response.data.Email;
        this.user.firstName=response.data.FirstName;
        this.user.lastName=response.data.LastName;
        console.log("log");
        console.log(this.user.email)
        //return this.user.email;
        
      });
    },
    decodeContent(){
      var decodedContent = atob(this.post.Content)
     // console.log(decodedContent, "content");
      this.decodedContent=decodedContent;
      
    }
 },
 mounted() {
   this.decodeContent();
   this.getUserById(this.post.UserID)

 }

}
</script>