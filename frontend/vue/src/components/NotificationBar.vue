<script setup lang="ts">
import { watch } from 'vue'

// 通知の型
export interface Notification {
  // タイプ
  type: 'info' | 'warn' | 'error'

  // メッセージ
  message: string

  // 自動で閉じる時間(ms)
  closeTimeout: number
}

// v-model 定義
const model = defineModel<Notification[]>({ required: true })

const closeNotifyAt = (index: number) => {
  if (model.value && index < model.value.length) {
    model.value.splice(index, 1)
  }
}

// 親コンポーネントに公開する関数
defineExpose({
  // 通知を行う関数
  notify: (type: 'info' | 'warn' | 'error', message: string, timeout: number) => {
    model.value?.push({ type, message, closeTimeout: timeout })
  }
})

// タイムアウトを実装するための watch
watch(
  () => model.value?.length,
  (_, oldLength) => {
    // 追加があったときだけ処理
    if (model.value && model.value.length > oldLength) {
      const notification = model.value[model.value?.length - 1]

      // 自動クローズが指定されている場合に通知を閉じる
      if (notification && notification.closeTimeout > 0) {
        setTimeout(
          () => (model.value = model.value?.filter((n) => n != notification)),
          notification.closeTimeout
        )
      }
    }
  }
)
</script>

<template>
  <teleport to="body">
    <transition-group name="notifier">
      <template v-for="(item, index) in model" :key="item">
        <v-sheet
          v-if="!!item.message" 
          class="notifier"
          :class="{
            info: item.type == 'info',
            warn: item.type == 'warn',
            error: item.type == 'error'
          }"
          @click="closeNotifyAt(index)"
        >
          <div class="notifier-contents px-3 py-3">
            <pre>{{ item.message }}</pre>
          </div>
        </v-sheet>
      </template>
    </transition-group>
  </teleport>
</template>

<style scoped>
.notifier {
  position: fixed;
  z-index: 10000;
  bottom: 0;
  left: 0;
  width: 100%;
  padding: 0px;
  color: #ffffff;
}

.info {
  background-color: #0079be;
}

.warn {
  background-color: orange;
}

.error {
  background-color: brown;
}

.notifier-contents {
  max-height: 8em;
  overflow-y: auto; /* 垂直スクロールを有効化 */
}

.notifier-enter-active {
  transition: all 0.3s ease-out;
}

.notifier-leave-active {
  transition: all 0.8s ease;
}

.notifier-enter-from,
.notifier-leave-to {
  transform: translateY(100px);
}
</style>
