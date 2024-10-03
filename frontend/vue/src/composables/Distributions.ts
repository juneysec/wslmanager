import { ref } from 'vue'
import client from '../apis'
import type * as Api from '../apis'

export class Distributions {
    // リスト取得
    static get() {
        const refData = ref<Api.ResponseDistributions>()
        const isFetching = ref<boolean>(false)
        const error = ref<Api.ResponseError | undefined>(undefined)

        async function execute() {
            error.value = undefined
            isFetching.value = true

            const { data, error: fetchError } = await client.GET('/distributions')

            if (fetchError) {
                error.value = fetchError
            } else {
                refData.value = data
            }

            isFetching.value = false
        }

        execute()

        return {
            data: refData,
            isFetching,
            error,
            execute
        }
    }


    // コマンド実行
    static put(distribution: string, command: "start" | "stop" | "shell" | "export" | "set-default", path: string | undefined) {
        const refData = ref<Api.ResponseDistributions>()
        const isFetching = ref<boolean>(false)
        const error = ref<Api.ResponseError | undefined>(undefined)

        async function execute() {
            error.value = undefined
            isFetching.value = true

            const { data, error: fetchError } = await client.PUT('/distributions/{distribution}', {
                params: {
                    path: { distribution },
                    body: { command, path }
                }
            })

            if (fetchError) {
                error.value = fetchError
            } else {
                refData.value = data
            }

            isFetching.value = false
        }

        execute()

        return {
            data: refData,
            isFetching,
            error,
            execute
        }
    }
}
