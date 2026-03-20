import { describe, it, expect } from 'vitest'
import { summarizeCommits } from '../summary'

describe('summarizeCommits', ()=>{
  it('returns count', ()=>{
    const out = summarizeCommits([{},{},{}])
    expect(out.count).toBe(3)
  })
})
