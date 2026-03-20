export function extractTodos(text:string){
  // naive: return lines containing TODO
  const lines = (text||'').split('\n')
  return lines.filter(l => l.includes('TODO'))
}
