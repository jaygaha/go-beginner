import webapi from "./gocliRequest"
import * as components from "./explorerComponents"
export * from "./explorerComponents"

/**
 * @description 
 * @param req
 */
export function queryExoplanets(req: components.ExoplanetQueryRequest) {
	return webapi.post<components.ExoplanetQueryResponse>(`/exoplanets/query`, req)
}
