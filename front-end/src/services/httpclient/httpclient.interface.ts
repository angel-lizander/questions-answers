export interface IHttpClient {
    url: string,
    parameters?: {} | null,
    headers?: {} | null
}

export interface IHttpResponse {
    success: boolean,
    data: {} | null,
    errorDetails: {} | null
}