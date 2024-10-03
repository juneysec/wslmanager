import type { paths as schemas } from '../generated/schema'
import type { components } from '../generated/schema'
import createClient from 'openapi-fetch'

export type ResponseError = components['schemas']['response_error']
export type ResponseDistributions = schemas['/distributions']['get']['responses']['200']['content']['application/json']
export type ResponseDistribution = schemas['/distributions/{distribution}']['get']['responses']['200']['content']['application/json']

const client = createClient<schemas>({ baseUrl: 'http://localhost:8080/api/v1' })
export default client;

