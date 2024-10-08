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

    try {
      const fetchResult = await client.PUT('/distributions/{distribution}', {
        params: {
          path: { distribution }
        },
        body: {
          command,
          path
        }
      })

      const { data, error: fetchError } = fetchResult
      if (fetchError) {
        error.value = fetchError
      } else {
        distributions.value = data
      }

      isFetching.value = false
      return fetchResult
    } catch(e) {
      isFetching.value = false
      error.value = {
        code: "TSER",
        message: `API呼び出しに失敗しました。\n${e}`
      }
    }
  }

  return {
    distributions,
    isFetching,
    error,
    distributionPut
  }
}
