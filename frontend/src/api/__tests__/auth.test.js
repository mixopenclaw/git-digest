import { login } from '../auth'

test('login returns token', async ()=>{
  const r = await login({})
  expect(r.token).toBe('stub-token')
})
