<script setup lang="ts">
import client from '../lib'
import LoadingDialog from '../components/LoadingDialog.vue'
import type { paths } from '../generated/schema'
import { ref } from 'vue'
import type { ParamsOption, RequestBodyOption } from 'openapi-fetch'

interface AppError {
  code: number
  message: string
}

type distributionsQueryOptions<T> = ParamsOption<T> & RequestBodyOption<T>
type DistributionsResponse =
  paths['/distributions']['get']['responses']['200']['content']['application/json']

const get = (fetchOptions: distributionsQueryOptions<paths['/distributions']['get']>) => {
  const state = ref<DistributionsResponse>()
  const isReady = ref(false)
  const isFetching = ref(false)
  const error = ref<AppError | undefined>(undefined)

  async function execute() {
    error.value = undefined
    isReady.value = false
    isFetching.value = true

    const { data, error: fetchError } = await client.GET('/distributions', fetchOptions)

    if (fetchError) {
      error.value = fetchError
    } else {
      state.value = data
      isReady.value = true
    }

    isFetching.value = false
  }

  execute()

  return {
    state,
    isReady,
    isFetching,
    error,
    execute
  }
}

const { state, isReady, isFetching, error, execute } = get({})
</script>

<template>
  <loading-dialog :isLoading="isFetching" />

  <v-sheet :loading="isFetching" variant="tonal" class="mx-auto w-auto">
    <v-table>
      <thead style="background-color: darkslategrey">
        <tr>
          <th></th>
          <th nowrap width="100%">ディストリビューション</th>
          <th nowrap>バージョン</th>
          <th nowrap width="0" class="text-center">状態</th>
          <th width="0" colspan="5">コマンド</th>
          <th width="0" class="text-right">
            <v-btn icon="mdi-import" class="mx-4 smbtn" variant="text" />
            <v-btn icon="mdi-power" class="smbtn" variant="text" />
          </th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="item in state" :key="item.distributions">
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
            <v-btn
              icon="mdi-play"
              class="play smbtn"
              :disabled="item.state == 'Running'"
              variant="text"
            >
            </v-btn>
          </td>
          <td>
            <v-btn
              icon="mdi-stop"
              class="stop smbtn"
              :disabled="item.state == 'Stopped'"
              variant="text"
            >
            </v-btn>
          </td>
          <td>
            <!-- bash 起動 -->
            <v-btn
              icon="mdi-bash"
              :disabled="item.state == 'Stopped'"
              variant="text"
              class="smbtn"
            />
          </td>
          <td>
            <!-- エクスポート -->
            <v-btn icon="mdi-export" variant="text" class="smbtn" />
          </td>
          <td>
            <!-- 削除 -->
            <v-btn
              icon="mdi-delete"
              class="del smbtn"
              :disabled="item.state == 'Running'"
              variant="text"
            />
          </td>
          <td>
            <v-container v-if="!item.isDefault" class="text-left"
              ><v-btn color="primary">デフォルトに設定</v-btn>
            </v-container>
          </td>
        </tr>
      </tbody>
    </v-table>
  </v-sheet>
</template>

<style>
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
</style>
