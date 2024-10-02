<script setup lang="ts">
const tools: { tool: string; command: string }[] = [
  { tool: 'XLaunch', command: 'winget install marha.VcXsrv' },
  { tool: 'Ubuntu-22.04', command: 'winget install Ubuntu-22.04' },
  { tool: 'Docker', command: 'winget install Docker.DockerDesktop' }
]
</script>

<template>
  <v-sheet>
    <v-container v-for="tool in tools" :key="tool.tool">
      <h1>{{ tool.tool }}</h1>
      <v-sheet>{{ tool.command }}</v-sheet>
    </v-container>
  </v-sheet>
</template>
