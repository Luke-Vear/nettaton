import React from 'react'
import Button from '@material-ui/core/Button'

export const BusyButton = ({ text, busy, onClick }) =>
  <div>
    {busy ? (
      <Button variant='contained' color='primary' disabled>Busy</Button>
    ) : (
      <Button variant='contained' color='primary' onClick={onClick}>{text}</Button>
    )}
  </div>
