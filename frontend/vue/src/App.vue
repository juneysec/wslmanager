<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { ref } from 'vue'

const drawer = ref(true)
</script>

<template>
  <v-responsive class="border rounded">
    <v-app>
      <v-app-bar style="background-color: darkslateblue">
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
            prepend-icon="mdi-home"
            title="ホーム"
            value="home"
            @click="$router.push('/')"
          />

          <v-list-item
            prepend-icon="mdi-tools"
            title="ツールインストール"
            value="tool"
            @click="$router.push('/toolinstaller')"
          />

          <v-list-item
            prepend-icon="mdi-forum"
            title="このアプリについて"
            value="about"
            @click="$router.push('/about')"
          />
        </v-list>
      </v-navigation-drawer>

      <v-main>
        <router-view />
      </v-main>
    </v-app>
  </v-responsive>
</template>

<style scoped></style>
