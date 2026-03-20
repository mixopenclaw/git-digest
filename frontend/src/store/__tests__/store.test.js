import { addResult, store } from '../store'

test('store addResult', ()=>{
  addResult({x:1})
  expect(store.results.length).toBeGreaterThan(0)
})
