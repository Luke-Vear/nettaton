import React from 'react'
import { connect } from 'react-redux'
import { actions } from './questionRedux'
import Button from '@material-ui/core/Button'

export const Question = ({ fetching, fetchQuestion }) => (
  <div>
    {fetching ? (
      <Button variant='contained' color='primary' disabled>Fetching...</Button>
    ) : (
      <Button variant='contained' color='primary' onClick={fetchQuestion}>New Question</Button>
    )}
  </div>
)

const mapStateToProps = state => ({
  fetching: state.fetching,
  question: state.question,
  error: state.error
})

const mapDispatchToProps = dispatch => ({
  fetchQuestion: () => dispatch(actions.apiCallRequest())
})

export default connect(mapStateToProps, mapDispatchToProps)(Question)
