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
    const { data, error: fetchError } = await client.GET('/distributions')

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
    distributionsGet
  }
}
