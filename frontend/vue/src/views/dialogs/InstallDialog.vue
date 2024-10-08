<script setup lang="ts">
import { ref, watch } from 'vue'
import type { ResponseOnlineDistributions } from '@/apis'
import LoadingDialog from '@/components/LoadingDialog.vue'
import NotificationBar, { type Notification } from '@/components/NotificationBar.vue'
import { useOnlineDistributionsGet } from '@/composables/online-distributions-get'
import { useOnlineDistributionsPost } from '@/composables/online-distributions-post'

const props = defineProps({
  onClose: Function
})

const showDialog = ref(false)
const distributions = ref<ResponseOnlineDistributions>([])

// 通知バー関連
const notification = ref()
const notifications = ref<Notification[]>([])
const notify = (type: 'info' | 'warn' | 'error', message: string, timeout: number) => {
  notification.value.notify(type, message, timeout)
}

// OnlineDistributionsGet 関連
const {
  distributions: onlineDistributionsGetData,
  isFetching: isOnlineDistributionsGetFetching,
  error: onlineDistributionsGetError,
  onlineDistributionsGet
} = useOnlineDistributionsGet()
watch(onlineDistributionsGetData, () => {
  distributions.value = onlineDistributionsGetData.value
})
watch(onlineDistributionsGetError, () => {
  if (onlineDistributionsGetError.value) {
    notify('error', onlineDistributionsGetError.value?.message ?? '', 0)
  }
})

// OnlineDistributionsPost 関連
const {
  distributions: onlineDistributionsPostData,
  isFetching: isOnlineDistributionsPostFetching,
  error: onlineDistributionsPostError,
  onlineDistributionsPost
} = useOnlineDistributionsPost()
watch(onlineDistributionsPostData, () => {
  distributions.value = onlineDistributionsPostData.value
})
watch(onlineDistributionsPostError, () => {
  if (onlineDistributionsPostError.value) {
    notify('error', onlineDistributionsPostError.value?.message ?? '', 0)
  }
})
watch(isOnlineDistributionsPostFetching, () => {
  if (!isOnlineDistributionsPostFetching.value) {
    onlineDistributionsGet().then((fetchResult) => {
      if (fetchResult && !fetchResult.error) {
        notify('info', 'ディストリビューションリストの読み込みに成功しました。', 3000)
      }
    })
  }
})

// ダイアログを開く
const open = () => {
  showDialog.value = true
  onlineDistributionsGet().then((fetchResult) => {
    if (fetchResult && !fetchResult.error) {
      notify('info', 'ディストリビューションリストの読み込みに成功しました。', 3000)
    }
  })
}

// ダイアログを閉じる
const close = () => {
  showDialog.value = false
  if (props.onClose) {
    props.onClose()
  }
}

// エクスポート定義
defineExpose({
  open,
  close
})
</script>

<template>
  <LoadingDialog
    :isLoading="isOnlineDistributionsGetFetching || isOnlineDistributionsPostFetching"
  />
  <NotificationBar v-model="notifications" ref="notification" />

  <v-dialog class="w-50" persistent v-model="showDialog">
    <v-card>
      <v-toolbar dense flat>
        <v-toolbar-title>ディストリビューションのインストール </v-toolbar-title>
      </v-toolbar>

      <v-card class="scrollable">
        <v-table>
          <tbody>
            <tr v-for="item in distributions" :key="item.name">
              <td>{{ item.friendlyName }}</td>
              <td>
                <div v-if="item.alreadyInstalled">インストール済</div>
                <div v-else>
                  <v-btn
                    color="primary"
                    @click="
                      onlineDistributionsPost(item.name!).then(
                        () => `${item.name!}をインストールしました。`
                      )
                    "
                    >インストール</v-btn
                  >
                </div>
              </td>
            </tr>
          </tbody>
        </v-table>
      </v-card>

      <v-card-actions class="pt-0">
        <v-spacer></v-spacer>
        <v-btn color="error" @click="close" class="mx-3">
          <v-icon class="text-button"> mdi-close </v-icon>
          閉じる
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<style scoped>
.scrollable {
  overflow-y: auto;
}
</style>
