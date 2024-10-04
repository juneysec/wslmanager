<script setup lang="ts">
import { ref } from 'vue'
import LoadingDialog from '../components/LoadingDialog.vue'
import * as Apis from '../apis'
import { Distributions } from '../composables/Distributions'
import NotificationBar from '../components/NotificationBar.vue'
import { type Notification } from '../components/NotificationBar.vue'

const isFetching = ref(false)
const data = ref<Apis.ResponseDistributions>()
const distributions = new Distributions({ isFetching, data })
const { execute: refresh } = distributions.get()

const notification = ref()
const notifications = ref<Notification[]>([])
const notify = (type: 'info' | 'warn' | 'error', message: string, timeout: number) => {
  notification.value.notify(type, message, timeout)
}
</script>

<template>
  <notification-bar v-model="notifications" ref="notification" />
  <loading-dialog :isLoading="isFetching" />

  <!-- 通知テスト TODO:削除 -->
  <v-btn @click="notify('warn', '<pre>test\nahoge</pre>', 3000)" />

  <v-sheet :loading="isFetching" variant="tonal" class="mx-auto w-auto">
    <v-table>
      <thead style="background-color: darkslategrey">
        <tr>
          <th></th>
          <th nowrap width="100%">ディストリビューション</th>
          <th nowrap>バージョン</th>
          <th nowrap width="0" class="text-center">状態</th>
          <th width="0" colspan="3">コマンド</th>
          <th width="0" colspan="3" class="text-right">
            <v-tooltip text="インポート...">
              <template v-slot:activator="{ props }">
                <v-btn icon="mdi-import" class="smbtn" variant="text" v-bind="props" />
              </template>
            </v-tooltip>

            <v-tooltip text="シャットダウン">
              <template v-slot:activator="{ props }">
                <v-btn icon="mdi-power" class="smbtn" variant="text" v-bind="props" />
              </template>
            </v-tooltip>

            <v-tooltip text="再読込">
              <template v-slot:activator="{ props }">
                <v-btn
                  icon="mdi-refresh"
                  class="smbtn ml-4"
                  variant="text"
                  @click="refresh()"
                  v-bind="props"
                />
              </template>
            </v-tooltip>
          </th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="item in data" :key="item.distributions">
          <td>
            <v-container v-if="item.isDefault" class="text-right">*</v-container>
          </td>
          <td>{{ item.name }}</td>
          <td>{{ item.version }}</td>
          <td>
            <v-container :class="[item.state == 'Running' ? 'play' : 'stop']">{{
              item.state
            }}</v-container>
          </td>
          <td>
            <v-tooltip :text="item.name + 'を起動する'">
              <template v-slot:activator="{ props }">
                <v-btn
                  icon="mdi-play"
                  class="play smbtn"
                  :disabled="item.state == 'Running'"
                  v-bind="props"
                  variant="text"
                  @click="distributions.put(item.name!, 'start')"
                />
              </template>
            </v-tooltip>
          </td>
          <td>
            <v-tooltip :text="item.name + 'を停止する'">
              <template v-slot:activator="{ props }">
                <v-btn
                  icon="mdi-stop"
                  class="stop smbtn"
                  :disabled="item.state == 'Stopped'"
                  v-bind="props"
                  variant="text"
                  @click="distributions.put(item.name!, 'stop')"
                />
              </template>
            </v-tooltip>
          </td>
          <td>
            <!-- bash 起動 -->
            <v-tooltip :text="item.name + 'の bash を起動する'">
              <template v-slot:activator="{ props }">
                <v-btn
                  icon="mdi-bash"
                  variant="text"
                  class="smbtn"
                  v-bind="props"
                  @click="distributions.put(item.name!, 'shell')"
                />
              </template>
            </v-tooltip>
          </td>

          <!-- デフォルトに設定 -->
          <td>
            <v-tooltip :text="item.name + 'をデフォルトに設定する'">
              <template v-slot:activator="{ props }">
                <v-container v-if="!item.isDefault" class="text-left"
                  ><v-btn
                    icon="mdi-asterisk"
                    class="smbtn"
                    variant="text"
                    v-bind="props"
                    @click="distributions.put(item.name!, 'set-default')"
                  ></v-btn>
                </v-container>
              </template>
            </v-tooltip>
          </td>

          <td>
            <!-- エクスポート -->
            <v-tooltip :text="item.name + 'をエクスポートする...'">
              <template v-slot:activator="{ props }">
                <v-btn icon="mdi-export" variant="text" class="smbtn" v-bind="props" />
              </template>
            </v-tooltip>
          </td>
          <td>
            <!-- 削除 -->
            <v-tooltip :text="item.name + 'を削除する...'">
              <template v-slot:activator="{ props }">
                <v-btn
                  icon="mdi-delete"
                  class="del smbtn"
                  :disabled="item.state == 'Running'"
                  variant="text"
                  v-bind="props"
                />
              </template>
            </v-tooltip>
          </td>
        </tr>
      </tbody>
    </v-table>
  </v-sheet>
</template>

<style scoped>
.play {
  color: lightgreen;
}

.stop {
  color: lightsalmon;
}

.del {
  color: lightcoral;
}

.smbtn {
  width: 2em !important;
  height: 2em !important;
}

.notifier {
  position: fixed;
  z-index: 10000;
  bottom: 0;
  left: 0;
  width: 100%;
  padding: 0px;
  background-color: #0079be;
  color: #ffffff;
}

.notifier-contents {
  max-height: 5em;
  overflow-y: auto; /* 垂直スクロールを有効化 */
}

.notifier-message {
  flex-grow: 1; /* 残りのスペースを占有 */
  margin-right: 10px; /* ボタンとの間に余白をつける */
}

.notifier-close-button {
  flex-shrink: 0; /* ボタンのサイズを固定 */
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
