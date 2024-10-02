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
const test = ref(true)
const save = () =>
{
  console.log('test')
}
</script>

<template>
  <loading-dialog :isLoading="isFetching" />

  <v-sheet :loading="isFetching" variant="tonal" class="mx-auto w-auto">
    <v-table>
      <thead>
        <tr>
          <th></th>
          <th nowrap>ディストリビューション</th>
          <th nowrap>状態</th>
          <th nowrap>バージョン</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in state" :key="item.distributions">
          <td>{{ item.isDefault ? '*' : '' }}</td>
          <td>{{ item.name }}</td>
          <td>{{ item.state }}</td>
          <td>{{ item.version }}</td>
        </tr>
      </tbody>
    </v-table>
  </v-sheet>
</template>
