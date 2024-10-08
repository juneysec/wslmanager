import { ref } from 'vue'
import client from '../apis'
import * as Api from '../apis'

// リスト取得
export const useOnlineDistributionsGet = () => {
  const distributions = ref<Api.ResponseOnlineDistributions>([])
  const isFetching = ref(false)
  const error = ref<Api.ResponseError | undefined>(undefined)

  async function onlineDistributionsGet() {
    error.value = undefined
    isFetching.value = true

    try {
      const fetchResult = await client.GET('/online-distributions')
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
        message: `API呼び出しに失敗しました。\n${e}`
      }
    }
  }

  return {
    distributions,
    isFetching,
    error,
    onlineDistributionsGet
  }
}
