<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { ref } from 'vue'

const drawer = ref(true)
</script>


<template>
  <v-responsive class="border rounded">
    <v-app>
      <v-app-bar>
        <template v-slot:prepend>
          <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
        </template>

        <v-app-bar-title>WSL Manager</v-app-bar-title>

        <template v-slot:append>
          <v-btn @click="$router.push('/')" icon="mdi-home" />
        </template>
      </v-app-bar>

      <v-navigation-drawer v-model="drawer">
        <v-list>
          <v-list-item
            prepend-icon="mdi-view-dashboard"
            title="Home"
            value="home"
            @click="$router.push('/')"
          ></v-list-item>
          <v-list-item
            prepend-icon="mdi-forum"
            title="About"
            value="about"
            @click="$router.push('/about')"
          ></v-list-item>
        </v-list>
      </v-navigation-drawer>
      
      <v-main>
        <v-container fluid >
          <!-- router-view / -->
    <v-row>
      <v-col cols="12" sm="1" md="1" lg="1" xl="1"> <card >card1</card> </v-col>
      <v-col cols="12" sm="1" md="1" lg="1" xl="1"> <card >card1</card> </v-col>
      <v-col cols="12" sm="1" md="1" lg="1" xl="1"> <card >card1</card> </v-col>
      <v-col cols="12" sm="1" md="1" lg="1" xl="1"> <card >card1</card> </v-col>
      <v-col cols="12" sm="1" md="1" lg="1" xl="1"> <card >card1</card> </v-col>
      <v-col cols="12" sm="1" md="1" lg="1" xl="1"> <card >card1</card> </v-col>
      <v-col cols="12" sm="1" md="1" lg="1" xl="1"> <card >card1</card> </v-col>
      <v-col cols="12" sm="1" md="1" lg="1" xl="1"> <card >card1</card> </v-col>
      <v-col cols="12" sm="1" md="1" lg="1" xl="1"> <card >card1</card> </v-col>
      <v-col cols="12" sm="1" md="1" lg="1" xl="1"> <card >card1</card> </v-col>
    </v-row>
        </v-container>
      </v-main>
    </v-app>
  </v-responsive>
</template>

<style scoped></style>
