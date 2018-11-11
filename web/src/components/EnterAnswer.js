import React from 'react'
import { connect } from 'react-redux'
import { actions } from './enterAnswerRedux'
import TextField from '@material-ui/core/TextField'

export const EnterAnswer = ({ answerEntryFieldValue, handleChange }) => (
  <TextField
    id='outlined-name'
    label='Answer'
    value={answerEntryFieldValue}
    onChange={handleChange(answerEntryFieldValue)}
    margin='normal'
    variant='outlined'
  />
)

const mapStateToProps = state => ({
  question: state.question,
  fetching: state.fetching,
  result: state.result,
  error: state.error
})

const mapDispatchToProps = dispatch => ({
  sendAnswer: (id) => dispatch(actions.apiCallRequest(id)),
  answerEntryFieldChange: () => 'foo'
})

export default connect(mapStateToProps, mapDispatchToProps)(EnterAnswer)
