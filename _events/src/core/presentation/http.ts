export interface HttpResponse<T = any> {
  code: number
  body: T
  error?: Error
}

export interface HttpRequest<T = any> {
  params: T
}
