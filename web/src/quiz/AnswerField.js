import React from 'react'
import TextField from '@material-ui/core/TextField'

export const AnswerField = ({ question, onChange }) =>
  <div>
    <TextField
      disabled={question === null}
      label='Answer'
      onChange={event => onChange(event.target.value)}
      margin='normal'
      variant='outlined'
    />
  </div>
