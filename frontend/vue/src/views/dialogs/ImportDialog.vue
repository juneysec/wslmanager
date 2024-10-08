<script setup lang="ts">
import { ref, watch } from 'vue'

const showDialog = defineModel<boolean>()
const props = defineProps({
  onSubmit: Function
})

const distributionName = ref('')
const importPath = ref('')
const sourcePath = ref('')
const canSubmit = ref(false)
const command = ref('')

const updateCanSubmit = () => {
  canSubmit.value =
    distributionName.value.length > 0 && importPath.value.length > 0 && sourcePath.value.length > 0

  if (canSubmit.value) {
    command.value = `wsl.exe --import "${distributionName.value}" "${importPath.value}" "${sourcePath.value}"`
  }
}

defineExpose({
  distributionName,
  importPath,
  sourcePath
})

watch(distributionName, updateCanSubmit)
watch(importPath, updateCanSubmit)
watch(sourcePath, updateCanSubmit)

const submit = async () => {
  if (props.onSubmit) {
    await props.onSubmit()
  }

  showDialog.value = false
}

const cancel = () => {
  showDialog.value = false
}
</script>

<template>
  <v-dialog v-model="showDialog" class="w-50" @keydown.esc="cancel">
    <v-card>
      <v-toolbar dense flat>
        <v-toolbar-title
          ><v-icon> mdi-help-circle </v-icon> ディストリビューションのインポート</v-toolbar-title
        >
      </v-toolbar>
      <v-card-text>
        <pre class="message">インポートに必要な下記の項目を入力してください。</pre>
      </v-card-text>

      <v-form action="#" @submit.prevent="submit()" class="mx-8">
        <v-text-field
          v-model="distributionName"
          variant="underlined"
          placeholder="ディストリビューション名"
          label="ディストリビューション名"
        />

        <v-text-field
          v-model="importPath"
          variant="underlined"
          placeholder="VHDファイルの作成先(フォルダを指定)"
          label="VHDファイルの作成先(フォルダを指定)"
        />

        <v-text-field
          v-model="sourcePath"
          variant="underlined"
          placeholder="インポートファイル"
          label="インポートファイル"
        />
      </v-form>

      <v-card-text>コマンドプレビュー：{{ command }}</v-card-text>

      <v-card-actions class="pt-0">
        <v-spacer></v-spacer>
        <v-btn color="error" @click="cancel" class="mx-3">
          <v-icon class="text-button"> mdi-cancel </v-icon>
          キャンセル
        </v-btn>
        <v-btn color="primary" @click="submit" :disabled="!canSubmit"
          ><v-icon class="text-button"> mdi-check-bold </v-icon>OK</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
