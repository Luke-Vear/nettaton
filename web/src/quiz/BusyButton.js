import React from 'react'
import Button from '@material-ui/core/Button'

export const BusyButton = ({ text, busy, color, onClick }) =>
  <div>
    {busy ? (
      <Button variant='contained' color={color} disabled>{text}</Button>
    ) : (
      <Button variant='contained' color={color} onClick={onClick}>{text}</Button>
    )}
  </div>
