import React from 'react'
import { render } from '@testing-library/react'
import Nav from '../Nav'

test('renders nav links',()=>{
  const { getByText } = render(<Nav />)
  expect(getByText('Home')).toBeTruthy()
})
