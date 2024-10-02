<script setup lang="ts">
import client from '../lib'
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
  <v-main>
    <v-dialog v-model="isFetching" persistent width="300">
      <v-card height="300" class="text-center">
        <v-progress-circular
          :indeterminate="isFetching"
          :size="200"
          color="primary"
          class="mx-auto my-auto"
          >Loading...
        </v-progress-circular>
      </v-card>
    </v-dialog>

    <v-card :loading="isFetching" class="mx-auto">
      <v-table>
        <thead>
          <tr>
            <th></th>
            <th>ディストリビューション</th>
            <th>状態</th>
            <th>バージョン</th>
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
    </v-card>
  </v-main>
</template>
