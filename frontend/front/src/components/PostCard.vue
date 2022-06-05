<template> <div>
  <v-card class="ma-4"  :to="'/post/'+post.ID" elevation="3">
    <v-card-title>{{post.Title}} </v-card-title>
    <v-card-subtitle>Autor: {{this.user.firstName}} {{user.lastName}} {{user.email}} {{post.UserID}}</v-card-subtitle>
    
    <v-card-text>
      <div >
      <div v-if="decodedContent.length<400" v-html="decodedContent"></div>
      <div v-else v-html="decodedContent.substring(0, 400)"></div>
       </div> 

      </v-card-text>
  </v-card>

</div>

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
        this.user.email=response.data.Email;
        this.user.firstName=response.data.FirstName;
        this.user.lastName=response.data.LastName;
        
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