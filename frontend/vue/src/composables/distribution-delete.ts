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

    const { data, error: fetchError } = await client.DELETE('/distributions/{distribution}', {
      params: {
        path: { distribution }
      },
      body: {
        name: distribution
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
    distributionDelete
  }
}
