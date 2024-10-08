<script setup lang="ts">
import { ref, watch } from 'vue'
import { useDistributionsPost } from '@/composables/distributions-post'
import LoadingDialog from '@/components/LoadingDialog.vue'
import NotificationBar, { type Notification } from '@/components/NotificationBar.vue'

const props = defineProps({
  onSubmit: Function
})

const showDialog = ref(false)
const distributionName = ref('')
const importPath = ref('')
const sourcePath = ref('')
const canSubmit = ref(false)
const commandPreview = ref('')

// distributionsPost 関連
const importDialog = ref()
const {
  isFetching: isDistributionsPostFetching,
  error: distributionsPostError,
  distributionsPost
} = useDistributionsPost()
watch(distributionsPostError, () => {
  if (distributionsPostError.value) {
    notify('error', distributionsPostError.value?.message ?? '', 0)
  }
})

// 通知バー関連
const notification = ref()
const notifications = ref<Notification[]>([])
const notify = (type: 'info' | 'warn' | 'error', message: string, timeout: number) => {
  notification.value.notify(type, message, timeout)
}

// サブミット可能かどうかを更新する
const updateCanSubmit = () => {
  canSubmit.value =
    distributionName.value.length > 0 && importPath.value.length > 0 && sourcePath.value.length > 0

  if (canSubmit.value) {
    commandPreview.value = `wsl.exe --import "${distributionName.value}" "${importPath.value}" "${sourcePath.value}"`
  }
}

// ダイアログを開く
const open = () => {
  showDialog.value = true
}

// ダイアログを閉じる
const close = () => {
  showDialog.value = false
}

// 入力値監視
watch(distributionName, updateCanSubmit)
watch(importPath, updateCanSubmit)
watch(sourcePath, updateCanSubmit)

const submit = () => {
  distributionsPost(
    importDialog.value.distributionName,
    importDialog.value.importPath,
    importDialog.value.sourcePath
  ).then((fetchResult) => {
    if (fetchResult) {
      const { error } = fetchResult

      if (error) {
        notify('error', error.message ?? '', 0)
      }
    }
    if (props.onSubmit) {
      props.onSubmit()
    }

    close()
  })
}

// エクスポート定義
defineExpose({
  open,
  close
})
</script>

<template>
  <LoadingDialog :isLoading="isDistributionsPostFetching" />
  <NotificationBar v-model="notifications" ref="notification" />

  <v-dialog v-model="showDialog" class="w-50" @keydown.esc="close">
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
