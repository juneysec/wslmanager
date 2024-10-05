<script setup lang="ts">
import { ref, watch } from 'vue'

const inputValue = defineModel()
const props = defineProps({
  validator: {
    type: Function,
    required: false
  }
})

const refShow = ref(false)
const refTitle = ref('')
const refMessage = ref('')
const refLabel = ref('')
const refResolve = ref<(value: unknown) => void>()
const refReject = ref<(reason?: any) => void>()
const refCanSubmit = ref(true)

// ダイアログのオープン
const open = (message: string, label: string, title: string = '入力') => {
  refShow.value = true
  refTitle.value = title
  refLabel.value = label
  refMessage.value = message
  return new Promise((resolve, reject) => {
    refResolve.value = resolve
    refReject.value = reject
  })
}

// サブミット処理
const submit = () => {
  if (refResolve.value) {
    refResolve.value(true)
  }

  refShow.value = false
}

// キャンセル処理
const cancel = () => {
  if (refResolve.value) {
    refResolve.value(false)
  }

  refShow.value = false
}

watch(inputValue, () => {
  if (props.validator) {
    refCanSubmit.value = props.validator(inputValue.value)
  } else {
    refCanSubmit.value = !!inputValue.value
  }
})

// 公開メソッド定義
defineExpose({
  open,
  submit,
  cancel
})
</script>

<template>
  <v-dialog v-model="refShow" class="w-50" @keydown.esc="cancel">
    <v-card>
      <v-toolbar dense flat>
        <v-toolbar-title><v-icon> mdi-help-circle </v-icon> {{ refTitle }}</v-toolbar-title>
      </v-toolbar>
      <v-card-text v-show="!!refMessage">
        <pre class="message">{{ refMessage }}</pre>
      </v-card-text>

      <v-form action="#" @submit.prevent="submit()" class="mx-8">
        <v-text-field
          v-model="inputValue"
          variant="underlined"
          :placeholder="refMessage"
          :label="refLabel"
        />
      </v-form>

      <v-card-actions class="pt-0">
        <v-spacer></v-spacer>
        <v-btn color="error" @click="cancel" class="mx-3">
          <v-icon class="text-button"> mdi-cancel </v-icon>
          キャンセル
        </v-btn>
        <v-btn color="primary" @click="submit" :disabled="!refCanSubmit"
          ><v-icon class="text-button"> mdi-check-bold </v-icon>OK</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<style scoped>
.message {
  white-space: pre-wrap;
}
</style>
