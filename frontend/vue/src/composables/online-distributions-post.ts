import { ref } from 'vue'
import client from '../apis'
import * as Api from '../apis'

// リスト取得
export const useOnlineDistributionsPost = () => {
  const distributions = ref<Api.ResponseOnlineDistributions>([])
  const isFetching = ref(false)
  const error = ref<Api.ResponseError | undefined>(undefined)

  async function onlineDistributionsPost(name: string) {
    error.value = undefined
    isFetching.value = true
    const { data, error: fetchError } = await client.POST('/online-distributions', {
      body: { name }
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
    onlineDistributionsPost
  }
}
