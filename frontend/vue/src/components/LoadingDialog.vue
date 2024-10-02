<script setup>
const props = defineProps({
  isLoading: Boolean
})
</script>

<template>
    <v-dialog v-model="props.isLoading" persistent width="300">
      <v-card height="300" class="text-center">
        <v-progress-circular
          indeterminate
          :size="200"
          color="primary"
          class="mx-auto my-auto"
          >Loading...
        </v-progress-circular>
      </v-card>
    </v-dialog>
</template>