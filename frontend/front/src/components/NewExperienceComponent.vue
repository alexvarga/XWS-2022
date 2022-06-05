<template>
  <div>
    <v-container>
      <v-card elevation="3">
        <v-card-title class="justify-center">Add experience</v-card-title>
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
          <v-col cols="1" sm="1" md="1">
            <v-btn @click="submit">Add</v-btn>
          </v-col>
        </v-row>
      </v-card>
    </v-container>
  </div>
</template>


<script>
import Axios from "axios";
export default {
  name: "NewExperienceComponent",
  props: ["userEmail"],
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

      postTitle: "Title",
      info: "Info placeholder",

      date: new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
        .toISOString()
        .substr(0, 10),
      date2: new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
        .toISOString()

        .substr(0, 10),
      menu: false,
      menu2: false,
    };
  },
  methods: {
    submit() {
      this.sendData.email = this.userEmail;
      this.sendData.experience[0].dateStarted = this.date;
      this.sendData.experience[0].dateEnded = this.date2;
      this.sendData.experience[0].title = this.postTitle;
      this.sendData.experience[0].info = this.info;

      console.log(this.sendData);
        Axios.post(
          "http://localhost:8080/api/user/user/experience",
          this.sendData
        ).then((response) => {
      console.log(this.sendData, "send data from inside");

          console.log(response.data);
        });
    },
  },
};
</script>
