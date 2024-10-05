import { ref } from 'vue'
import client from '../apis'
import type * as Api from '../apis'

export type CommandEnum = 'start' | 'stop' | 'shell' | 'export' | 'set-default'

// コマンド実行
export const useDistributionPut = () => {
  const distributions = ref<Api.ResponseDistributions>([])
  const isFetching = ref(false)
  const error = ref<Api.ResponseError | undefined>(undefined)

  async function distributionPut(
    distribution: string,
    command: CommandEnum,
    path: string | undefined = undefined
  ) {
    error.value = undefined
    isFetching.value = true

    const { data, error: fetchError } = await client.PUT('/distributions/{distribution}', {
      params: {
        path: { distribution }
      },
      body: {
        command,
        path
      }
    })

    if (fetchError) {
      error.value = fetchError
    } else {
      distributions.value = data
    }

    isFetching.value = false
  }

  return {
    distributions,
    isFetching,
    error,
    distributionPut
  }
}
