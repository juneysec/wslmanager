import { ref } from 'vue'
import client from '../apis'
import type * as Api from '../apis'

// ディストリビューションのインポート
export const useDistributionsPost = () => {
  const distributions = ref<Api.ResponseDistributions>([])
  const isFetching = ref(false)
  const error = ref<Api.ResponseError | undefined>(undefined)

  async function distributionsPost(name: string, vhdPath: string, sourcePath: string) {
    error.value = undefined
    isFetching.value = true

    try {
      const fetchResult = await client.POST('/distributions', {
        body: {
          name,
          vhdPath,
          sourcePath
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
        code: "TSER",
        message: `API呼び出しに失敗しました。\n${e}`
      }
    }
  }

  return {
    distributions,
    isFetching,
    error,
    distributionsPost
  }
}
