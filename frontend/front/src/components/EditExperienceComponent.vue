<template>
  <div>
        <v-snackbar v-model="snackbar" centered timeout="2500">
      <span>{{ snackbarText }}</span>
    </v-snackbar>
    <v-container>
      <v-card elevation="3">
        <v-card-title class="justify-center">Edit experience</v-card-title>
        <v-row justify="center">
          <v-col cols="12" sm="8" md="8"> </v-col>
        </v-row>
        <v-row justify="center">
          <v-col cols="12" sm="8" md="8">
            <v-text-field v-model="postTitle" label="Title"></v-text-field>
          </v-col>
        </v-row>

        <v-row justify="center">
          <v-col cols="12" sm="4" md="4">
            <v-menu
              ref="menu"
              v-model="menu"
              :close-on-content-click="false"
              :return-value.sync="date"
              transition="scale-transition"
              offset-y
              min-width="auto"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-text-field
                  v-model="date"
                  label="Date started"
                  prepend-icon="mdi-calendar"
                  readonly
                  v-bind="attrs"
                  v-on="on"
                ></v-text-field>
              </template>
              <v-date-picker v-model="date" no-title scrollable>
                <v-btn text color="primary" @click="menu = false">
                  Cancel
                </v-btn>
                <v-btn text color="primary" @click="$refs.menu.save(date)">
                  OK
                </v-btn>
              </v-date-picker>
            </v-menu>
          </v-col>

          <v-col cols="12" sm="4" md="4">
            <v-menu
              ref="menu2"
              v-model="menu2"
              :close-on-content-click="false"
              :return-value.sync="date2"
              transition="scale-transition"
              offset-y
              min-width="auto"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-text-field
                  v-model="date2"
                  label="Date ended"
                  prepend-icon="mdi-calendar"
                  readonly
                  v-bind="attrs"
                  v-on="on"
                ></v-text-field>
              </template>
              <v-date-picker v-model="date2" no-title scrollable>
                <v-spacer></v-spacer>
                <v-btn text color="primary" @click="menu2 = false">
                  Cancel
                </v-btn>
                <v-btn text color="primary" @click="$refs.menu2.save(date2)">
                  OK
                </v-btn>
              </v-date-picker>
            </v-menu>
          </v-col>
        </v-row>
        <v-row justify="center">
          <v-col cols="12" sm="8" md="8">
            <v-textarea
              v-model="info"
              label="Info/description"
              width="100px"
            ></v-textarea>
          </v-col>
        </v-row>
        <v-row justify="center">
          <v-col cols="4" sm="2" md="2">
            <v-btn @click="update">Update</v-btn>
          </v-col>
          <v-col cols="4" sm="2" md="2">
            <v-btn @click="cancelUpdate">Cancel</v-btn>
          </v-col>
        </v-row>
      </v-card>
    </v-container>
  </div>
</template>


<script>
import Axios from "axios";
export default {
  name: "EditExperienceComponent",
  props: ["experienceItem", "userEmail"],
  data() {
    return {
      sendData: {
        email: "",
        experience: [
          {
            dateStarted: "",
            dateEnded: "",
            title: "",
            info: "",
          },
        ],
      },

      postTitle: this.experienceItem.Title,
      info: this.experienceItem.Info,

      date: this.experienceItem.DateStarted,
      date2: this.experienceItem.DateEnded,

      date2: new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
        .toISOString()
        .substr(0, 10),
      menu: false,
      menu2: false,
      snackbar: false,
      snackbarText: "",
    };
  },
  methods: {
    update() {
      //this needs to change

      this.sendData.email = this.userEmail;
      this.sendData.experience[0].dateStarted = this.date;
      this.sendData.experience[0].dateEnded = this.date2;
      this.sendData.experience[0].title = this.postTitle;
      this.sendData.experience[0].info = this.info;

      Axios.put(
        "http://localhost:8080/api/user/user/experience/" +
          this.experienceItem.ID,
        this.sendData
      ).then((response) => {
        this.snackbar = true;
        this.snackbarText = "Sucessfully Edited";
      });
    },
    cancelUpdate() {
      this.$emit("cancelEdit", "test");
    },
  },
};
</script>
