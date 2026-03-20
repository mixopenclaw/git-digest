export function jsonResponse(res:any, obj:any){
  res.setHeader('content-type','application/json')
  res.end(JSON.stringify(obj))
}
