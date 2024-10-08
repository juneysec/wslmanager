<script setup lang="ts">
import { ref, watch } from 'vue'
import * as Apis from '../apis'
import { useDistributionsGet } from '@/composables/distributions-get'
import { useDistributionPut } from '@/composables/distribution-put'
import { useDistributionDelete } from '@/composables/distribution-delete'
import LoadingDialog from '@/components/LoadingDialog.vue'
import NotificationBar, { type Notification } from '@/components/NotificationBar.vue'
import ConfirmDialog from '@/components/ConfirmDialog.vue'
import InputDialog from '@/components/InputDialog.vue'
import ImportDialog from '@/views/dialogs/ImportDialog.vue'
import InstallDialog from '@/views/dialogs/InstallDialog.vue'
import MoveDialog from '@/views/dialogs/MoveDialog.vue'
import { useDistributionsPost } from '@/composables/distributions-post'

const distributions = ref<Apis.ResponseDistributions>([])

// distributionsGet 関連
const {
  distributions: distributionsGetData,
  isFetching: isDistributionsGetFetching,
  error: distributionsGetError,
  distributionsGet
} = useDistributionsGet()
watch(distributionsGetData, () => {
  distributions.value = distributionsGetData.value
})
watch(distributionsGetError, () => {
  if (distributionsGetError.value) {
    err(
      `ディストリビューションリストの取得に失敗しました。\n${distributionsGetError.value?.message}`
    )
  }
})

// ディストリビューションリストの読み込み
const loadDistributions = () => {
  distributionsGet().then((fetchResult) => {
    if (fetchResult) {
      info('ディストリビューションリストを読み込みました。')
    }
  })
}

loadDistributions()

// インポートダイアログ 関連
const importDialog = ref()
const onImportSubmit = () => {
  loadDistributions()
}

// distributionPut 関連
const {
  distributions: distributionPutData,
  isFetching: isDistributionPutFetching,
  error: distributionPutError,
  distributionPut
} = useDistributionPut()
watch(distributionPutData, () => {
  distributions.value = distributionPutData.value
})
watch(distributionPutError, () => {
  if (distributionPutError.value) {
    err(distributionPutError.value?.message ?? '')
  }
})

// distributionDelete 関連
const {
  distributions: distributionDeleteData,
  isFetching: isDistributionDeleteFetching,
  error: distributionDeleteError,
  distributionDelete
} = useDistributionDelete()
watch(distributionDeleteData, () => {
  distributions.value = distributionDeleteData.value
})
watch(distributionDeleteError, () => {
  if (distributionDeleteError.value) {
    err(distributionDeleteError.value?.message ?? '')
  }
})

// 通知バー関連
const notification = ref()
const notifications = ref<Notification[]>([])
const notify = (type: 'info' | 'warn' | 'error', message: string, timeout: number) => {
  notification.value.notify(type, message, timeout)
}
const info = (message: string) => notify('info', message, 3000)
const err = (message: string) => notify('error', message, 0)

// 確認ダイアログ関連
const confirmDialog = ref()
const stop = async (distribution: string) => {
  if (await confirmDialog.value.open(`${distribution}を停止します。\nよろしいですか？`)) {
    distributionPut(distribution, 'stop').then((fetchResult) => {
      info(`${distribution}を停止しました。`)
    })
  }
}

// インストールダイアログ
const installDialog = ref()

// VHD移動ダイアログ
const moveDialog = ref()

// 入力ダイアログ関連
const pathInputDialog = ref()
const pathValue = ref('')
const deleteInputDialog = ref()
const deleteValue = ref('')

// ディストリビューションのエクスポート
const exportDistribution = async (distribution: string) => {
  if (
    await pathInputDialog.value.open(
      `${distribution}のエクスポート先を入力してください。`,
      `${distribution}のエクスポート先`
    )
  ) {
    if (pathValue.value) {
      distributionPut(distribution, 'export', pathValue.value).then(() =>
        info(
          `${distribution}のエクスポート処理を開始しました。\n進捗についてはターミナルを確認してください。`
        )
      )
    }
  }
}

const targetName = ref('')
const deleteDistribution = async (distribution: string) => {
  deleteValue.value = ''
  targetName.value = distribution

  if (
    await deleteInputDialog.value.open(
      `${distribution}を削除します。\nこの操作が間違いでない場合、下の入力ボックスに ${distribution} と入力してOKを押してください。\n※完全に削除するには、プログラムの追加と削除からディストリビューションをアンインストールしてください。`,
      `${distribution}`
    )
  ) {
    if (deleteValue.value === distribution) {
      distributionDelete(deleteValue.value).then(() => info(`${distribution}を削除しました。`))
    }
  }
}

const deleteValidator = (val: string) => {
  return val == targetName.value
}
</script>

<template>
  <NotificationBar v-model="notifications" ref="notification" />
  <LoadingDialog
    :isLoading="
      isDistributionPutFetching || isDistributionsGetFetching || isDistributionDeleteFetching
    "
  />
  <ConfirmDialog ref="confirmDialog" />
  <InputDialog ref="pathInputDialog" v-model="pathValue" />
  <InputDialog ref="deleteInputDialog" v-model="deleteValue" :validator="deleteValidator" />
  <ImportDialog ref="importDialog" :onSubmit="loadDistributions" />
  <InstallDialog ref="installDialog" :onClose="loadDistributions" />
  <MoveDialog ref="moveDialog" :onSubmit="loadDistributions" />

  <v-layout-item model-value position="bottom" class="text-right" size="88">
    <div class="ma-4">
      <v-btn
        icon="mdi-plus"
        size="large"
        color="primary"
        elevation="8"
        @click="installDialog.open()"
      />
    </div>
  </v-layout-item>

  <v-sheet
    :loading="
      isDistributionPutFetching || isDistributionsGetFetching || isDistributionDeleteFetching
    "
    variant="tonal"
    class="mx-auto w-auto"
  >
    <v-table>
      <thead style="background-color: darkslategrey">
        <tr>
          <th></th>
          <th nowrap width="100%">ディストリビューション</th>
          <th nowrap>バージョン</th>
          <th nowrap width="0" class="text-center">状態</th>
          <th width="0" colspan="3" nowrap>コマンド</th>
          <th width="0" colspan="3" class="text-right" nowrap>
            <v-tooltip text="インポート...">
              <template v-slot:activator="{ props }">
                <v-btn
                  icon="mdi-import"
                  class="smbtn"
                  variant="text"
                  v-bind="props"
                  @click="importDialog.open()"
                />
              </template>
            </v-tooltip>

            <v-tooltip text="再読込">
              <template v-slot:activator="{ props }">
                <v-btn
                  icon="mdi-refresh"
                  class="smbtn ml-4"
                  variant="text"
                  @click="loadDistributions"
                  v-bind="props"
                />
              </template>
            </v-tooltip>
          </th>
        </tr>
      </thead>

      <tbody>
        <template v-for="item in distributions" :key="item.name">
          <tr class="upper-row">
            <td>
              <div v-if="item.isDefault" class="text-right">*</div>
            </td>
            <td>{{ item.name }}</td>
            <td>{{ item.version }}</td>
            <td>
              <div :class="[item.state == 'Running' ? 'play' : 'stop']">{{ item.state }}</div>
            </td>
            <td>
              <v-tooltip :text="item.name + 'を起動する'">
                <template v-slot:activator="{ props }">
                  <v-btn
                    icon="mdi-play"
                    class="play smbtn"
                    :disabled="item.state != 'Stopped'"
                    v-bind="props"
                    variant="text"
                    @click="
                      distributionPut(item.name!, 'start').then(() =>
                        info(`${item.name!}を起動しました。`)
                      )
                    "
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
                    :disabled="item.state !== 'Running'"
                    v-bind="props"
                    variant="text"
                    @click="stop(item.name!)"
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
                    @click="
                      distributionPut(item.name!, 'shell').then(() =>
                        info(`${item.name!}のシェルを起動しました。`)
                      )
                    "
                  />
                </template>
              </v-tooltip>
            </td>

            <!-- デフォルトに設定 -->
            <td>
              <v-tooltip :text="item.name + 'をデフォルトに設定する'">
                <template v-slot:activator="{ props }">
                  <div v-if="!item.isDefault" class="text-left">
                    <v-btn
                      icon="mdi-asterisk"
                      class="smbtn"
                      variant="text"
                      v-bind="props"
                      @click="
                        distributionPut(item.name!, 'set-default').then(() =>
                          info(`${item.name!}をデフォルトに設定しました。`)
                        )
                      "
                    ></v-btn>
                  </div>
                </template>
              </v-tooltip>
            </td>

            <td>
              <!-- エクスポート -->
              <v-tooltip :text="item.name + 'をエクスポートする...'">
                <template v-slot:activator="{ props }">
                  <v-btn
                    icon="mdi-export"
                    variant="text"
                    class="smbtn"
                    v-bind="props"
                    @click="exportDistribution(item.name!)"
                  />
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
                    @click="deleteDistribution(item.name!)"
                  />
                </template>
              </v-tooltip>
            </td>
          </tr>
          <tr class="down-row">
            <td></td>
            <td colspan="8">
              <div class="mx-4">
                VHDパス:
                <a
                  href="#"
                  @click="
                    distributionPut(item.name!, 'open-vhd').then(() =>
                      info('VHDのパスを開きました。')
                    )
                  "
                  >{{ item.vhdPath }}</a
                >
                ({{ (item.vhdSize! / 1024 / 1024 / 1024).toFixed(2) }} GiB)
              </div>
            </td>
            <td>
              <v-tooltip :text="item.name + 'のVHDパスを移動する...'" v-if="!!item.vhdPath">
                <template v-slot:activator="{ props }">
                  <v-btn
                    icon="mdi-folder-move"
                    class="smbtn"
                    :disabled="item.state !== 'Stopped'"
                    variant="text"
                    v-bind="props"
                    @click="moveDialog.open(item.name!, item.vhdPath)"
                  />
                </template>
              </v-tooltip>
            </td>
          </tr>
        </template>
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

.upper-row > td {
  border-bottom: 0 !important;
}

.down-row > td {
  border-top: 0 !important;
  height: 0 !important;
  padding-bottom: 16px !important;
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
