import { createServer, IncomingMessage, ServerResponse } from 'http'
import { summarizeCommits } from './summary'
import { getCommitSummary } from './commit'
import { extractTodos } from './todos'

function handler(req: IncomingMessage, res: ServerResponse){
  const url = req.url || ''
  if(url.startsWith('/summary')){
    res.setHeader('content-type','application/json')
    res.end(JSON.stringify({ ok: true }))
    return
  }
  if(url.startsWith('/commit/')){
    res.setHeader('content-type','application/json')
    res.end(JSON.stringify({ ok: true }))
    return
  }
  if(url.startsWith('/todos')){
    res.setHeader('content-type','application/json')
    res.end(JSON.stringify({ todos: [] }))
    return
  }
  res.statusCode = 404
  res.end('not found')
}

createServer(handler).listen(8080)
console.log('server listening on http://localhost:8080')
