import React from 'react'
import { render } from '@testing-library/react'
import ScanForm from '../ScanForm'

test('scan form renders', () => {
  const { getByPlaceholderText } = render(<ScanForm />)
  expect(getByPlaceholderText('Path')).toBeTruthy()
})
