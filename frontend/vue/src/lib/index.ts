import type { paths } from '../generated/schema'
import createClient from 'openapi-fetch'

const client = createClient<paths>({ baseUrl: 'http://localhost:8080/api/v1' })
export default client
