import React from 'react'
import { connect } from 'react-redux'
import { actions } from './redux'
import { BusyButton } from './BusyButton'
import { ResultDisplay } from './ResultDisplay'
import { AnswerField } from './AnswerField'
import { QuestionDisplay } from './QuestionDisplay'

export const Quiz = ({ busy, question, answer, correct, error, newQuestion, sendAnswer, updateAnswer }) =>
  <div>
    <QuestionDisplay busy={busy} correct={correct} question={question} error={error} />
    <ResultDisplay correct={correct} />
    <AnswerField question={question} onChange={updateAnswer} />
    <BusyButton text={'Next'} busy={busy} color='primary' onClick={newQuestion} />
    <BusyButton text={'Send'} busy={busy || question === null} color='secondary' onClick={() => sendAnswer(question, answer)} />
  </div>

const mapStateToProps = state => ({
  busy: state.quiz.busy,
  question: state.quiz.question,
  answer: state.quiz.answer,
  correct: state.quiz.correct,
  error: state.quiz.error
})

const mapDispatchToProps = dispatch => ({
  newQuestion: () => dispatch(actions.newQuestionRequest()),
  sendAnswer: (question, answer) => dispatch(actions.sendAnswerRequest(question, answer)),
  updateAnswer: (answer) => dispatch(actions.answerUpdate(answer))
})

export default connect(mapStateToProps, mapDispatchToProps)(Quiz)
