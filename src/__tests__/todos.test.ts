import { describe, it, expect } from 'vitest'
import { extractTodos } from '../todos'

describe('extractTodos', ()=>{
  it('finds TODO lines', ()=>{
    const out = extractTodos('line1\n// TODO: fix\nend')
    expect(out.length).toBe(1)
  })
})
