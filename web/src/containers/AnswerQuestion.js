import React from 'react'
import { connect } from 'react-redux'
import { answerQuestion } from '../actions'

let AnswerQuestion = ({ dispatch }) => {
  let input

  return (
    <div>
      <form onSubmit={e => {
        e.preventDefault()
        if (!input.value.trim()) {
          return
        }
        dispatch(answerQuestion(input.value))
        input.value = ''
      }}>
        <input ref={node => {
          input = node
        }} />
        <button type="submit">
          Answer Question
        </button>
      </form>
    </div>
  )
}
AnswerQuestion = connect()(AnswerQuestion)

export default AnswerQuestion
