import React from 'react'
import { connect } from 'react-redux'
import { actions } from './redux'
import { BusyButton } from './BusyButton'
import { Display } from './Display'
import TextField from '@material-ui/core/TextField'

export const Quiz = ({ busy, question, answer, result, error, newQuestion, sendAnswer, updateAnswer }) =>
  <div>
    <Display busy={busy} result={result} question={question} error={error} />
    <TextField
      id='outlined-name'
      label='Answer'
      value={answer}
      onChange={event => updateAnswer(event.target.value)}
      margin='normal'
      variant='outlined'
    />
    <BusyButton text={'Next'} busy={busy} onClick={newQuestion} />
    <BusyButton text={'Send'} busy={busy} onClick={(question, answer) => sendAnswer(question, answer)} />
  </div>

const mapStateToProps = state => ({
  busy: state.quiz.busy,
  question: state.quiz.question,
  answer: state.quiz.answer,
  result: state.quiz.result,
  error: state.quiz.error
})

const mapDispatchToProps = dispatch => ({
  newQuestion: () => dispatch(actions.newQuestionRequest()),
  sendAnswer: (question, answer) => dispatch(actions.sendAnswerRequest(question, answer)),
  updateAnswer: (answer) => dispatch(actions.updateAnswer(answer))
})

export default connect(mapStateToProps, mapDispatchToProps)(Quiz)
