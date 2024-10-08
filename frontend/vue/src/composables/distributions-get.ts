import { ref } from 'vue'
import client from '../apis'
import * as Api from '../apis'

// リスト取得
export const useDistributionsGet = () => {
  const distributions = ref<Api.ResponseDistributions>([])
  const isFetching = ref(false)
  const error = ref<Api.ResponseError | undefined>(undefined)

  async function distributionsGet() {
    error.value = undefined
    isFetching.value = true

    try {
      const fetchResult = await client.GET('/distributions')
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
        code: "TSER",
        message: `API呼び出しに失敗しました。${e}`
      }
    }
  }

  return {
    distributions,
    isFetching,
    error,
    distributionsGet
  }
}
