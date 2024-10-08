import { ref } from 'vue'
import client from '../apis'
import type * as Api from '../apis'

// コマンド実行
export const useDistributionDelete = () => {
  const distributions = ref<Api.ResponseDistributions>([])
  const isFetching = ref(false)
  const error = ref<Api.ResponseError | undefined>(undefined)

  async function distributionDelete(distribution: string) {
    error.value = undefined
    isFetching.value = true

    try {
      const fetchResult = await client.DELETE('/distributions/{distribution}', {
        params: {
          path: { distribution }
        },
        body: {
          name: distribution
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
    } catch (e) {
      isFetching.value = false
      error.value = {
        code: 'TSER',
        message: `API呼び出しに失敗しました。\n${e}`
      }
    }
  }

  return {
    distributions,
    isFetching,
    error,
    distributionDelete
  }
}
