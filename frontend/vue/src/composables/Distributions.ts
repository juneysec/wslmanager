import { ref, type Ref } from 'vue'
import client from '../apis'
import type * as Api from '../apis'

export type CommandEnum = 'start' | 'stop' | 'shell' | 'export' | 'set-default'

export type FetchRef = {
  isFetching: Ref<boolean, boolean>
  data: Ref<Record<string, never>[] | undefined, Record<string, never>[] | undefined>
}

export class Distributions {
  constructor(private ref: FetchRef) {}

  // リスト取得
  public get() {
    const self = this
    const error = ref<Api.ResponseError | undefined>(undefined)

    async function execute() {
      error.value = undefined
      self.ref.isFetching.value = true

      const { data, error: fetchError } = await client.GET('/distributions')

      if (fetchError) {
        error.value = fetchError
      } else {
        self.ref.data.value = data
      }

      self.ref.isFetching.value = false
    }

    execute()

    return {
      data: this.ref.data,
      isFetching: this.ref.isFetching,
      error,
      execute
    }
  }

  // コマンド実行
  public put(distribution: string, command: CommandEnum, path: string | undefined = undefined) {
    const self = this
    const error = ref<Api.ResponseError | undefined>(undefined)

    async function execute() {
      error.value = undefined
      self.ref.isFetching.value = true

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
        self.ref.data.value = data
      }

      self.ref.isFetching.value = false
    }

    execute()

    return {
      data: this.ref.data,
      isFetching: this.ref.isFetching,
      error,
      execute
    }
  }
}
