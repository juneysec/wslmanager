<script setup lang="ts">
import { ref, watch } from 'vue'
import { useDistributionPut } from '@/composables/distribution-put'
import LoadingDialog from '@/components/LoadingDialog.vue'
import NotificationBar, { type Notification } from '@/components/NotificationBar.vue'
import type * as Api from '@/apis'


const props = defineProps({
  onSubmit: Function
})

const showDialog = ref(false)
const distribution = ref('')
const destinationPath = ref('')
const canSubmit = ref(false)
const commandPreview = ref('')

// 通知バー関連
const notification = ref()
const notifications = ref<Notification[]>([])
const notify = (type: 'info' | 'warn' | 'error', message: string, timeout: number) => {
  notification.value.notify(type, message, timeout)
}

// useDistributionPut 関連
const importDialog = ref()
const {
  distributions,
  isFetching,
  error,
  distributionPut
} = useDistributionPut()
watch(error, () => {
  if (error.value) {
    notify('error', error.value?.message ?? '', 0)
  }
})

// サブミット可能かどうかを更新する
const updateCanSubmit = () => {
  canSubmit.value = destinationPath.value.length > 0

  if (canSubmit.value) {
    commandPreview.value = `wsl.exe --manage "${distribution.value}" --move "${destinationPath.value}"`
  }
}

// ダイアログを開く
const open = (distributionName: string, path: string) => {
  if (!!distributionName) {
      showDialog.value = true
      distribution.value = distributionName
      destinationPath.value = path
  }
}

// ダイアログを閉じる
const close = () => {
  showDialog.value = false
}

// 入力値監視
watch(distribution, updateCanSubmit)
watch(destinationPath, updateCanSubmit)

const submit = () => {
  distributionPut(distribution.value, "move-vhd", destinationPath.value).then((fetchResult) => {
    const { error } = fetchResult

    if (error instanceof Api.ResponseError) {
      notify('error', `コマンド実行でエラーが発生しました。\n${error.message}`, 0)
    } else {
      if (props.onSubmit) {
        props.onSubmit()
      }

      close()
    }
  })
}

// エクスポート定義
defineExpose({
  open,
  close,
})
</script>

<template>
  <LoadingDialog :isLoading="isFetching" />
  <NotificationBar v-model="notifications" ref="notification" />

  <v-dialog v-model="showDialog" class="w-50" @keydown.esc="close" persistent>
    <v-card>
      <v-toolbar dense flat>
        <v-toolbar-title
          ><v-icon> mdi-help-circle </v-icon> ディストリビューションの移動</v-toolbar-title
        >
      </v-toolbar>
      <v-card-text>
        <pre class="message">{{ distribution }}の移動先を入力してください。</pre>
      </v-card-text>

      <v-form action="#" @submit.prevent="submit()" class="mx-8">
        <v-text-field
          v-model="destinationPath"
          variant="underlined"
          placeholder="移動先(フォルダを指定)"
          label="移動先(フォルダを指定)"
        />

      </v-form>

      <v-card-text>コマンドプレビュー：{{ commandPreview }}</v-card-text>

      <v-card-actions class="pt-0">
        <v-spacer></v-spacer>
        <v-btn color="error" @click="close" class="mx-3">
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
